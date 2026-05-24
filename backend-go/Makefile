# Run API development server
dev/api:
	go run cmd/api/main.go

# Clean and verify Go modules
tidy:
	go mod tidy

# Run all tests
test:
	go test ./...

# Update Go version in go.mod
# Usage:
# make go/edit VERSION=1.25.0
go/edit:
	go mod edit -go=$(VERSION)
	go mod tidy

# Update all project dependencies to latest version
go/update:
	go get -u ./...
	go mod tidy

# Install Swagger CLI and dependencies
swagger/install:
	go install github.com/swaggo/swag/cmd/swag@latest
	go get github.com/swaggo/swag@latest
	go mod tidy

# Generate Swagger documentation for Gin application
swagger/gin/generate:
	swag init -g internal/app/gin.go -o api/gin/docs

# comments
## Add Go binary path to PATH environment variable
## export PATH=$PATH:$(go env GOPATH)/bin
#