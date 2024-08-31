
SERVER_TCP_PORT=8000

PHONY: run-server
run-server:
	go run ./cmd/httpserver/

migrate-up:
	sql-migrate up --env="local"

migrate-up1:
	sql-migrate up --env="local" -limit=1

migrate-down:
	sql-migrate down --env="local"go install github.com/rubenv/sql-migrate/...@latest

migrate-down1:
	sql-migrate down --env="local" -limit=1

sqlc:
	sqlc generate

test:
	 find . -name "*_test.go" -exec dirname {} \; | sort -u | xargs go test -v -cover

.DEFAULT_GOAL := run-server

