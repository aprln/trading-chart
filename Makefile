LOCAL_BIN:=$(CURDIR)/bin
export PATH:=$(LOCAL_BIN):$(PATH)

GOLANGCI_VERSION:=1.60.3

DB_DSN="postgres://postgres:postgres@localhost:5433/trading_chart?sslmode=disable"

install-tools: install-golang-migrate
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v$(GOLANGCI_VERSION)

install-golang-migrate:
	GOBIN=$(LOCAL_BIN) go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

protoc:
	protoc --go_out=./grpc-server/pkg/ --go-grpc_out=./grpc-server/pkg/ ./grpc-server/protobuf/trading_chart.proto

db-up: ## Run the DB in a docker container
	docker-compose up postgres -d

db-down: ## Stop the DB in a docker container
	docker-compose stop postgres

migrate-up: install-golang-migrate ## Migration up
	@$(LOCAL_BIN)/migrate -path=./grpc-server/db/migrations -database $(DB_DSN) up

migrate-down: ## Migration down
	@$(LOCAL_BIN)/migrate -path=./grpc-server/db/migrations -database $(DB_DSN) down 1

run-server: db-up migrate-up ## Run server
	go run ./grpc-server

run-ws-listener: ## Run aggregator
	go run ./websocket-listener

# -------
# Docker
# -------

run-server-docker: db-up migrate-up
	docker-compose up grpc-server --build

run-ws-listener-docker:
	docker-compose up websocket-listener --build