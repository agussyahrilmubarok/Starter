# export PATH=$PATH:$(go env GOPATH)/bin

web/run:
	go run cmd/web/main.go

web/build:
	go build -o bin/web cmd/web/main.go

web/test:
	go test ./...

web/test/coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out

web/fmt:
	go fmt ./...

web/vet:
	go vet ./...

web/tidy:
	go mod tidy

web/download:
	go mod download

web/clean:
	go clean
	rm -rf bin coverage.out

web/check: web/fmt web/vet web/test


mockery/setup:
	go install github.com/vektra/mockery/v2@latest

# mockery --name=UserRepository --dir=internal/domain --output=internal/domain/mocks --outpkg=mocks
mockery/domain/generate:
	mockery --name=$(name) --dir=internal/domain --output=internal/domain/mocks --outpkg=mocks