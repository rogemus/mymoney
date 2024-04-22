ifneq (,$(wildcard ./.env))
    include .env
    export
endif

server:
	@echo "Starting go server ..."
	@go run cmd/web/main.go

buildProd:
	@echo "Starting production build..."
	GOOS=linux GOARCH=amd64 go build -o ./tmp/main-linux ./cmd/web/main.go

buildDev:
	@echo "Starting build ..."
	@go build -v -o ./tmp/main ./cmd/web/main.go

dev:
	@echo "Starting go server with live reload"
	@air

tests:
	@echo "Testing..."
	@go test ./test/...

testRepos:
	@echo "Testing repos ..."
	@go test ./test/pkg/repository/...

testHandlers:
	@echo "Testing handlers..."
	@go test ./test/pkg/handler/...


migration:
	GOOSE_DRIVER=postgres \
	GOOSE_DBSTRING="host=$(DB_HOST) port=$(DB_PORT) user=$(DB_USER) password=$(DB_PASS) dbname=$(DB_NAME) sslmode=disable" \
	goose -dir db/migrations $(action)
