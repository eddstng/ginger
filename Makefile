.PHONY: test server db

test-int:
	docker compose down psql_test
	docker compose up -d psql_test
	go clean -testcache
	sleep 5
	cd server && TEST_MODE=true go test ./tests/integration/... -cover


server:
	docker compose up -d psql
	sleep 5
	cd server && go run main.go

db:
	docker compose up -d psql
