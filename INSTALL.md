
# SDeploy Installation & Usage

## Build

```sh
go mod tidy
go build -o sdeploy ./cmd/sdeploy
```

## Test

```sh
go test ./cmd/sdeploy/...
```

## Run

### Console mode
```sh
./sdeploy -c /path/to/config.json
```

### Daemon mode
```sh
./sdeploy -d
```

## Trigger Deployment via Webhook

```sh
curl -X POST "http://localhost:8080/hooks/frontend?secret=token" \
	-d '{"ref":"refs/heads/main"}'
```
