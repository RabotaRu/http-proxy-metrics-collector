# http-proxy-metrics-collector

## Development

1. Run target server `cd hello-server && go run main.go`
2. Run proxy server `TARGET_URL="http://localhost:1323" go run main.go`
3. Check proxy: `curl http://localhost:3000`  You will get `HelloWorld`

