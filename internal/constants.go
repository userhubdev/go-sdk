// Code generated. DO NOT EDIT.

package internal

import (
	"time"
)

const (
	ApiBaseUrl = "https://api.userhub.com"
	UserAgent  = "UserHub-Go/0.6.2"
	Version    = "0.6.2"

	AuthHeader     = "Authorization"
	ApiKeyHeader   = "UserHub-Api-Key"
	AdminKeyPrefix = "sk_"
	UserKeyPrefix  = "pk_"

	WebhookActionHeader        = "UserHub-Action"
	WebhookAgentHeader         = "Webhook-Agent"
	WebhookMaxRequestSizeBytes = 5242880
	WebhookMaxTimestampDiff    = 300000 * time.Millisecond
	WebhookSignatureHeader     = "UserHub-Signature"
	WebhookTimestampHeader     = "UserHub-Timestamp"
	WebhookServerErrorJson     = `{"message":"Webhook server error","code":"INTERNAL"}`

	SummarizeBodyLength = 20

	ConnectTimeout   = 30000 * time.Millisecond
	IdleConnTimeout  = 30000 * time.Millisecond
	MaxBodySizeBytes = 5242880
	MaxConns         = 100
	MaxIdleConns     = 5
	RetryMaxAttempts = 5
	ReadTimeout      = 30000 * time.Millisecond
	RequestTimeout   = 60000 * time.Millisecond
	RetryMaxSleep    = 20000 * time.Millisecond
	RetryMultiplier  = 100 * time.Millisecond
	RetryOverhead    = 100 * time.Millisecond
	TlsTimeout       = 10000 * time.Millisecond
	WriteTimeout     = 30000 * time.Millisecond
)
