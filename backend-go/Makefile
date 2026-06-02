# export PATH=$PATH:$(go env GOPATH)/bin

backend/run:
	go run cmd/backend/main.go

swagger/setup:
	go install github.com/swaggo/swag/cmd/swag@latest
	go get github.com/swaggo/swag@latest
	go mod tidy

swagger/generate:
	swag init -g cmd/backend/main.go -o api/docs

mockery/setup:
	go install github.com/vektra/mockery/v2@latest

# mockery --name=UserRepository --dir=internal/domain --output=internal/domain/mocks --outpkg=mocks
mockery/domain/generate:
	mockery --name=$(name) --dir=internal/domain --output=internal/domain/mocks --outpkg=mocks

# mockery --name=UserUseCase --dir=internal/application/usecase --output=internal/application/usecase/mocks --outpkg=mocks
mockery/usecase/generate:
	mockery --name=$(name) --dir=internal/application/usecase --output=internal/application/usecase/mocks --outpkg=mocks