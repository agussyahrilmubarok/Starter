# export PATH=$PATH:$(go env GOPATH)/bin

web/run:
	go run cmd/web/main.go

mockery/setup:
	go install github.com/vektra/mockery/v2@latest

# mockery --name=UserRepository --dir=internal/domain --output=internal/domain/mocks --outpkg=mocks
mockery/domain/generate:
	mockery --name=$(name) --dir=internal/domain --output=internal/domain/mocks --outpkg=mocks