server:
	@echo "Starting go server ..."
	@go run cmd/web/main.go

tests:
	@echo "Testing..."
	@go test ./test/...
