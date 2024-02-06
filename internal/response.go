package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	req  *Request
	res  *http.Response
	body []byte
}

func (r *Response) DecodeBody(v any) error {
	if r != nil && len(r.body) > 0 {
		err := json.Unmarshal(r.body, v)
		if err != nil {
			return Errorf(r.req, r.res, "Failed to decode response%s", summarizeBody(r.body))
		}
	}
	return nil
}

func summarizeBody(b []byte) string {
	if len(b) == 0 {
		return ""
	}

	s := removeWhitespace.ReplaceAllString(string(b[:SummarizeBodyLength*2]), " ")
	if len(s) == 0 {
		return ""
	}

	return fmt.Sprintf(": %s...", s[:SummarizeBodyLength])
}
