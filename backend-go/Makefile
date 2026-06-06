# export PATH=$PATH:$(go env GOPATH)/bin

backend/run:
	go run cmd/backend/main.go

backend/build:
	go build -o bin/backend cmd/backend/main.go

backend/test:
	go test ./...

backend/test/coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out

backend/fmt:
	go fmt ./...

backend/vet:
	go vet ./...

backend/tidy:
	go mod tidy

backend/download:
	go mod download

backend/clean:
	go clean
	rm -rf bin coverage.out

backend/check: backend/fmt backend/vet backend/test

swagger/setup:
	go install github.com/swaggo/swag/cmd/swag@latest
	go get github.com/swaggo/swag@latest
	go mod tidy

swagger/generate:
	swag init -g cmd/backend/main.go -o api/docs

mockery/setup:
	go install github.com/vektra/mockery/v2@latest

# make mockery/domain/generate name=UserRepository
mockery/domain/generate:
	mockery --name=$(name) --dir=internal/domain --output=internal/domain/mocks --outpkg=mocks

# make mockery/usecase/generate name=UserUseCase
mockery/usecase/generate:
	mockery --name=$(name) --dir=internal/application/usecase --output=internal/application/usecase/mocks --outpkg=mocks