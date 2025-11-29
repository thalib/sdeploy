# SDeploy

A lightweight, Go-based daemon that automates deployments via webhooks.

## Features

- **Webhook Listener** — HTTP endpoint for GitHub, GitLab, or CI/CD triggers
- **HMAC & Secret Auth** — Secure requests via signature or query parameter
- **Branch Filtering** — Only deploy matching branches
- **Single Execution** — One deployment at a time per project (lock-based)
- **Git Integration** — Optional `git pull` before running deploy commands
- **Email Notifications** — Send deployment summaries on completion
- **Daemon Mode** — Run as a background service with logging

## Quick Start

```sh
# Build
go build -o sdeploy ./cmd/sdeploy

# Run (console mode)
./sdeploy -c config.json

# Run (daemon mode)
./sdeploy -c config.json -d
```

See [INSTALL.md](INSTALL.md) for detailed build, test, and deployment instructions.

## Usage

```
sdeploy [options]

Options:
  -c <path>  Path to config file
  -d         Run as daemon (background service)
  -h         Show help
```

Config file search order:
1. Path from `-c` flag
2. `/etc/sdeploy/config.json`
3. `./config.json`

## Configuration

See [samples/config.json](samples/config.json) for a complete example.

| Key             | Description                              |
|-----------------|------------------------------------------|
| `listen_port`   | HTTP port (default: 8080)                |
| `log_filepath`  | Log file path (daemon mode)              |
| `email_config`  | SMTP settings for notifications          |
| `projects`      | Array of project configurations          |

### Project Config

| Key               | Description                                  |
|-------------------|----------------------------------------------|
| `name`            | Project identifier                           |
| `webhook_path`    | Unique URI path (e.g., `/hooks/api`)         |
| `webhook_secret`  | Secret for authentication                    |
| `git_branch`      | Branch required to trigger deployment        |
| `execute_command` | Shell command to run                         |
| `execute_path`    | Working directory for command                |
| `git_update`      | Run `git pull` before deployment             |
| `email_recipients`| Notification email addresses                 |

## Triggering Deployments

**Via webhook (HMAC signature):**
```sh
curl -X POST http://localhost:8080/hooks/myproject \
  -H "X-Hub-Signature: sha1=..." \
  -d '{"ref":"refs/heads/main"}'
```

**Via secret query parameter (internal/cron):**
```sh
curl -X POST "http://localhost:8080/hooks/myproject?secret=your_secret" \
  -d '{"ref":"refs/heads/main"}'
```

## Documentation

- [INSTALL.md](INSTALL.md) — Build, test, and deployment instructions
- [SPEC.md](SPEC.md) — Full specification and architecture details

## License

See [LICENSE](LICENSE).
