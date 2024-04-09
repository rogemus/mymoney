server:
	@echo "Starting go server ..."
	@go run cmd/web/main.go

build:
	@echo "Starting build ..."
	@go build -o ./tmp/main ./cmd/web/main.go

dev:
	@echo "Starting go server with live reload"

tests:
	@echo "Testing..."
	@go test ./test/...
