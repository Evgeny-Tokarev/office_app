
SERVER_TCP_PORT=8000

PHONY: run-server
run-server:
	go run ./cmd/httpserver/

migrate-up:
	sql-migrate up --env="local"

migrate-down:
	sql-migrate down --env="local"

sqlc:
	sqlc generate

test:
	go test -v -cover

.DEFAULT_GOAL := run-server

