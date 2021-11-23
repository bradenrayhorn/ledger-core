SHELL = bash
PG_PORT?=55001
PG_DATABASE?=ledger_core

migrate:
	migrate -database "postgres://postgres:password@127.0.0.1:${PG_PORT}/${PG_DATABASE}?sslmode=disable" -path sql/migrations up

migrate-down:
	migrate -database 'postgres://postgres:password@127.0.0.1:${PG_PORT}/${PG_DATABASE}?sslmode=disable' -path sql/migrations down

build:
	go build -o ./ledger-core ./cmd/core

run:
	./ledger-core

start: build run

test:
	docker-compose -f docker-compose.test.yml up --abort-on-container-exit --build
	docker-compose -f docker-compose.test.yml down --volumes

report:
	go tool cover -html=./reports/coverage.txt
