.PHONY: test server

test:
	docker compose down psql_test
	docker compose up -d psql_test
	sleep 1
	cd server && go test ./... -cover
	docker compose down psql_test

server:
	docker compose up -d psql
	sleep 2
	cd server && go run main.go