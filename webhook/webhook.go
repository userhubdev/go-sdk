package webhook

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
	"io"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/userhubdev/go-sdk/code"
	"github.com/userhubdev/go-sdk/internal"
	"github.com/userhubdev/go-sdk/internal/options"
	"github.com/userhubdev/go-sdk/option"
	"github.com/userhubdev/go-sdk/types"
)

var (
	errInternal = types.NewError("Webhook unhandled error").
		SetApiCode(code.Internal)
)

// Webhook is a parsing and dispatch helper for UserHub webhooks.
type Webhook struct {
	mu            sync.RWMutex
	signingSecret []byte
	handlers      map[string]Handler
	onError       func(ctx context.Context, err error)
}

// New returns a new instance of Webhook.
func New(signingSecret string, opts ...option.WebhooksOption) *Webhook {
	o := &options.WebhooksOptions{
		OnError: defaultOnError,
	}
	for _, opt := range opts {
		if opt != nil {
			opt.Apply(o)
		}
	}

	return &Webhook{
		signingSecret: []byte(signingSecret),
		handlers:      map[string]Handler{},
		onError:       o.OnError,
	}
}

// HandleAction executes a handler based on specified Request.
func (w *Webhook) HandleAction(ctx context.Context, r *Request) *Response {
	err := w.Verify(r)
	if err != nil {
		return w.CreateResponse(ctx, err)
	}

	action, err := r.GetAction()
	if err != nil {
		return w.CreateResponse(ctx, err)
	}

	w.mu.RLock()
	handler := w.handlers[action]
	if handler == nil {
		if action == "challenge" {
			handler = challengeHandler{}
		} else {
			handler = w.handlers[""]
		}
	}
	w.mu.RUnlock()

	if handler == nil {
		handler = unimplementedHandler{}
	}

	value, err := handler.HandleAction(ctx, r)
	if err != nil {
		return w.CreateResponse(ctx, err)
	}

	return w.CreateResponse(ctx, value)
}

// OnAction registers an action handler.
func (w *Webhook) OnAction(name string, handler Handler) *Webhook {
	w.mu.Lock()
	defer w.mu.Unlock()

	if handler != nil {
		w.handlers[name] = handler
	} else {
		delete(w.handlers, name)
	}

	return w
}

// OnDefault registers a fallback action handler.
func (w *Webhook) OnDefault(fn HandlerFunc) *Webhook {
	return w.OnAction("", fn)
}

// Verify ensures the body matches the specified signature/timestamp and returns
// an error if verification fails.
func (w *Webhook) Verify(req *Request) error {
	if w == nil || len(w.signingSecret) == 0 {
		return types.NewError("Signing secret is required")
	}
	if req == nil {
		return types.NewError("Request is required")
	}
	if len(req.Headers) == 0 {
		return types.NewError("Headers are required")
	}
	if len(req.Body) == 0 {
		return types.NewError("Body is required")
	}

	_, err := req.GetAction()
	if err != nil {
		return err
	}

	timestamp, err := req.GetTimestamp()
	if err != nil {
		return err
	}

	signatures, err := req.GetSignatures()
	if err != nil {
		return err
	}

	unixTime, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		return types.NewError("Timestamp is invalid: %s", timestamp)
	}

	diff := time.Until(time.Unix(unixTime, 0))
	if diff > internal.WebhookMaxTimestampDiff {
		return types.NewError("Timestamp is too far in the future: %s", timestamp)
	} else if diff < -internal.WebhookMaxTimestampDiff {
		return types.NewError("Timestamp is too far in the past: %s", timestamp)
	}

	mac := hmac.New(sha256.New, w.signingSecret)
	mac.Write([]byte(timestamp))
	mac.Write([]byte("."))
	mac.Write(req.Body)
	digest := mac.Sum(nil)

	var matched bool

	for _, signature := range signatures {
		signatureBytes, err := hex.DecodeString(signature)
		if err != nil {
			continue
		}

		if subtle.ConstantTimeCompare(digest, signatureBytes) == 1 {
			matched = true
			break
		}
	}

	if !matched {
		return types.NewError("Signatures do not match")
	}

	return nil
}

// CreateResponse creates a response from an object that can be encoded
// using json.Marshal or error.
func (w *Webhook) CreateResponse(ctx context.Context, value any) *Response {
	errValue, ok := value.(error)
	if ok {
		w.tryOnError(ctx, errValue)
	}

	// createResponse always returns a Response even if it also returns an
	// error.
	res, err := createResponse(value)
	if err != nil {
		w.tryOnError(ctx, err)
	}
	return res
}

// ServeHTTP handles incoming webhook requests.
func (w *Webhook) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	r := http.MaxBytesReader(res, req.Body, internal.WebhookMaxRequestSizeBytes)
	defer r.Close()

	if req.Method != "POST" {
		w.writeResponse(res, req, w.CreateResponse(ctx,
			types.NewError("Request should be a POST: %s", req.Method)))
		return
	}

	body, err := io.ReadAll(r)
	if err != nil {
		w.writeResponse(res, req, w.CreateResponse(ctx,
			types.NewError("Failed to read request body: %w", err)))
		return
	}

	w.writeResponse(res, req, w.HandleAction(ctx, &Request{
		Headers: req.Header.Clone(),
		Body:    body,
	}))
}

func (w *Webhook) writeResponse(res http.ResponseWriter, req *http.Request, r *Response) {
	headers := res.Header()

	for name, value := range r.Headers {
		headers[name] = value
	}

	res.WriteHeader(r.StatusCode)

	_, err := res.Write(r.Body)
	if err != nil {
		w.tryOnError(req.Context(), err)
	}
}

func (w *Webhook) tryOnError(ctx context.Context, err error) {
	if w.onError != nil && err != nil {
		w.onError(ctx, err)
	}
}
