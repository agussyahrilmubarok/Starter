api:
	go run cmd/api/main.go

swagger/install:
	go install github.com/swaggo/swag/cmd/swag@latest
	go get github.com/swaggo/swag@latest
	go mod tidy

swagger/gin/generate:
	swag init -g internal/app/gin.go -o api/gin/docs

# export PATH=$PATH:$(go env GOPATH)/bin