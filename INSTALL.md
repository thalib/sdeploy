# SDeploy Installation

## Requirements

- Go 1.24+ (or Docker)

## Build & Test

### Standard Build

```sh
go mod tidy
go build -o sdeploy ./cmd/sdeploy
go test ./cmd/sdeploy/... -v
```

### Build with Docker

```sh
docker run --rm -v "$(pwd):/app" -w /app golang:latest \
  sh -c "go mod tidy"

docker run --rm -v "$(pwd):/app" -w /app golang:latest \
  sh -c "go build -buildvcs=false -o sdeploy ./cmd/sdeploy"

## Run Test
docker run --rm -v "$(pwd):/app" -w /app golang:latest \
  sh -c "go test ./cmd/sdeploy/... -v"
```

## Run

```sh
## Console mode:
./sdeploy -c config.json

## Daemon mode:

./sdeploy -c config.json -d
```

## Install as systemd Service

1. Copy binary:

```sh
sudo cp sdeploy /usr/local/bin/
```

2. Create config:

```sh
sudo mkdir -p /etc/sdeploy
sudo cp samples/config.json /etc/sdeploy/config.json
# Edit config as needed

sudo cp samples/sdeploy.service /etc/systemd/system/sdeploy.service
```

3. Enable and start:

```sh
sudo systemctl daemon-reload
sudo systemctl enable sdeploy
sudo systemctl start sdeploy
```



### Example Scenario

If your config specifies:
```json
{
  "listen_port": 8080,
  "log_filepath": "/var/log/sdeploy/daemon.log",
  "projects": [
    {
      "name": "SDeploy Test",
      "webhook_path": "/shooks/sdeploy-test",
      "webhook_secret": "your_webhook_secret_here",
      "git_repo": "https://github.com/devnodesin/sdeploy-test.git",
      "git_branch": "main",
      "git_update": true,
      "local_path": "/opt/sdeploy/sdeploy-test",
      "execute_command": "sh build.sh"
    }
  ]
}
```

## Verify

```sh
# Check status
sudo systemctl status sdeploy

# Test webhook
curl -X POST "http://localhost:8080/shooks/sdeploy-test?secret=your_webhook_secret_here" \
  -d '{"ref":"refs/heads/main"}'
```
