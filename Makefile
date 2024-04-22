ifneq (,$(wildcard ./.env))
    include .env
    export
endif

server:
	@echo "Starting go server ..."
	@go run cmd/web/main.go

build:
	@make buildGoProd
	@echo "Backend Build DONE !!!"

	@echo "---"

	@make buildUiProd
	@echo "UI Build DONE !!!"

buildGoProd:
	@echo "Starting go production build..."
	GOOS=linux GOARCH=amd64 go build -o ./tmp/main-linux ./cmd/web/main.go

buildGoDev:
	@echo "Starting go build ..."
	@go build -v -o ./tmp/main ./cmd/web/main.go

buildUiProd:
	@echo "Starting UI build ..."
	@npm run build --prefix ui

devGo:
	@echo "Starting go server with live reload..."
	@air

devUI:
	@echo "Starting ui server with live reload..."
	@npm run start --prefix ui

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
