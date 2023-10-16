build:
	@go build -o bin/main cmd/server/main.go

run:
	@go run cmd/server/main.go

test: 
	@go test ./...

test_coverage:
	@go test ./... -coverprofile=coverage.out

up:
	@docker compose up

down:
	@docker compose down

up_db:
	@docker compose up postgresdb