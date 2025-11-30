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
./sdeploy -c sdeploy.conf

## Daemon mode:

./sdeploy -c sdeploy.conf -d
```

## Install as systemd Service

1. Copy binary:

```sh
sudo cp sdeploy /usr/local/bin/
```

2. Create config:

```sh
sudo cp samples/sdeploy.conf /etc/sdeploy.conf
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
```yaml
listen_port: 8080
log_filepath: /var/log/sdeploy/daemon.log

projects:
  - name: SDeploy Test
    webhook_path: /hooks/sdeploy-test
    webhook_secret: your_webhook_secret_here
    git_repo: https://github.com/devnodesin/sdeploy-test.git
    git_branch: main
    git_update: true
    local_path: /opt/sdeploy/sdeploy-test
    execute_command: sh build.sh
```

## Verify

```sh
# Check status
sudo systemctl status sdeploy

# Test webhook
curl -X POST "http://localhost:8080/hooks/sdeploy-test?secret=your_webhook_secret_here" \
  -d '{"ref":"refs/heads/main"}'
```

## Migration from JSON to YAML

If you're upgrading from an older version that used JSON configuration:

1. **Config file location changed:**
   - Old: `/etc/sdeploy/config.json` or `./config.json`
   - New: `/etc/sdeploy.conf` or `./sdeploy.conf`

2. **Convert your JSON config to YAML format:**
   - Remove curly braces `{}` and square brackets `[]`
   - Replace `:` with `: ` (colon followed by space)
   - Use indentation for nested objects
   - Arrays use `- ` prefix for each item

3. **Example conversion:**

   **Before (JSON):**
   ```json
   {
     "listen_port": 8080,
     "projects": [
       {
         "name": "My Project",
         "webhook_path": "/hooks/myproject",
         "webhook_secret": "secret123"
       }
     ]
   }
   ```

   **After (YAML):**
   ```yaml
   listen_port: 8080
   projects:
     - name: My Project
       webhook_path: /hooks/myproject
       webhook_secret: secret123
   ```

4. **Update systemd service** (if applicable):
   - The service file no longer needs a config directory
   - Config is now at `/etc/sdeploy.conf`

5. **Backup and remove old files:**
   ```sh
   sudo mv /etc/sdeploy/config.json /etc/sdeploy/config.json.bak
   sudo rmdir /etc/sdeploy  # if empty
   ```
