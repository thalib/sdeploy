# SDeploy Installation

## Requirements

- Go 1.24+ (or Docker)

## Build

```sh
go mod tidy
go build -o sdeploy ./cmd/sdeploy
```

**Using Docker:**
```sh
docker run --rm -v "$(pwd):/app" -w /app golang:latest \
  sh -c "go mod tidy && go build -buildvcs=false -o sdeploy ./cmd/sdeploy"
```

## Test

```sh
go test ./cmd/sdeploy/...
```

**Using Docker:**
```sh
docker run --rm -v "$(pwd):/app" -w /app golang:latest \
  sh -c "go test ./cmd/sdeploy/..."
```

## Run

**Console mode:**
```sh
./sdeploy -c config.json
```

**Daemon mode:**
```sh
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
   ```

3. Create service file `/etc/systemd/system/sdeploy.service`:
   ```ini
   [Unit]
   Description=SDeploy Webhook Daemon
   After=network.target

   [Service]
   Type=simple
   ExecStart=/usr/local/bin/sdeploy -d
   Restart=always

   [Install]
   WantedBy=multi-user.target
   ```

4. Enable and start:
   ```sh
   sudo systemctl daemon-reload
   sudo systemctl enable sdeploy
   sudo systemctl start sdeploy
   ```

## Verify

```sh
# Check status
sudo systemctl status sdeploy

# Test webhook
curl -X POST "http://localhost:8080/hooks/myproject?secret=your_secret" \
  -d '{"ref":"refs/heads/main"}'
```
