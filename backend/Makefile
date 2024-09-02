SERVER_TCP_PORT=8000
GO_BIN_PATH=/usr/local/go/bin

.PHONY: run-server migrate-up migrate-up1 migrate-down migrate-down1 sqlc test
run-server:
	PATH=$(GO_BIN_PATH):$(PATH) go run ./cmd/httpserver/

migrate-up:
	PATH=$(GO_BIN_PATH):$(PATH) sql-migrate up --env="local"

migrate-up1:
	PATH=$(GO_BIN_PATH):$(PATH) sql-migrate up --env="local" -limit=1

migrate-down:
	PATH=$(GO_BIN_PATH):$(PATH) sql-migrate down --env="local"

migrate-down1:
	PATH=$(GO_BIN_PATH):$(PATH) sql-migrate down --env="local" -limit=1

sqlc:
	sqlc generate

test:
	 PATH=$(GO_BIN_PATH):$(PATH) find . -name "*_test.go" -exec dirname {} \; | sort -u | xargs go test -v -cover

.DEFAULT_GOAL := run-server

#go install github.com/rubenv/sql-migrate/...@latest

