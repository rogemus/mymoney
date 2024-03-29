server:
	@echo "Starting go server ..."
	@go run cmd/web/main.go

test:
	@echo "Testing..."
	@go test ./...
