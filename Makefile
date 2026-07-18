test:
	go test -count=0 -race ./...

lint:
	golangci-lint run ./...

coverage:
	go test -coverprofile=coverage.out ./...; go tool cover -html=coverage.out; rm coverage.out

publish:
	GOPROXY=proxy.golang.org go list -m github.com/getaudited/audited@v0.0.1