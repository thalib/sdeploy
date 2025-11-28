## Overview
- Implement HTTP webhook listener and request handling
- Route requests by URI path to correct project
- Authenticate via HMAC signature or secret query parameter

## Requirements
- HTTP server on configurable port (default 8080)
- Route POST requests by `webhook_path` to matching project
- HMAC-SHA256 signature validation via `X-Hub-Signature-256` header
- Fallback to `?secret=` query parameter for internal triggers
- Parse webhook payload to extract branch information
- Return 202 Accepted for valid requests, appropriate errors otherwise
- Classify trigger source: WEBHOOK (HMAC) or INTERNAL (secret param)

## Acceptance
- TDD: Write `webhook_test.go` tests BEFORE implementation
- Tests cover: routing, HMAC validation, secret fallback, branch extraction, error responses
- `go test ./cmd/sdeploy/... -run TestWebhook` passes
- Code location: `cmd/sdeploy/webhook.go`, `cmd/sdeploy/webhook_test.go`
