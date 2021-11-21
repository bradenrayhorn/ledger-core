migrate:
	 migrate -database 'postgres://postgres:password@127.0.0.1:${POSTGRES_PORT}/ledger_core?sslmode=disable' -path sql/migrations up

migrate-down:
	 migrate -database 'postgres://postgres:password@127.0.0.1:${POSTGRES_PORT}/ledger_core?sslmode=disable' -path sql/migrations down

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
