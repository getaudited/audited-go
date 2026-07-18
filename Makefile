test:
	go test -count=0 -race ./...

lint:
	golangci-lint run ./...

coverage:
	go test -coverprofile=coverage.out ./...; go tool cover -html=coverage.out; rm coverage.out