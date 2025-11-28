## Overview
- Implement email notification for deployment results
- Send summary emails on deployment completion (success or failure)
- Use SMTP with TLS support

## Requirements
- Send email via SMTP using global `email_config` settings
- Support TLS/STARTTLS on port 587
- Email contains: project name, trigger source, status, start/end time, output summary
- Send to all addresses in project's `email_recipients` array
- Handle SMTP errors gracefully (log, don't crash)
- Skip email if `email_recipients` is empty or `email_config` not set

## Acceptance
- TDD: Write `email_test.go` tests BEFORE implementation
- Tests cover: email composition, SMTP connection, error handling, skip when unconfigured
- `go test ./cmd/sdeploy/... -run TestEmail` passes
- Code location: `cmd/sdeploy/email.go`, `cmd/sdeploy/email_test.go`
