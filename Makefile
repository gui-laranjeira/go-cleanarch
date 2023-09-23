build:
	@go build -o bin/main cmd/server/main.go

run:
	@go run cmd/server/main.go

compile:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin cmd/server/main.go

create_container: 
	@docker run --name library-api -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -d postgres:12-alpine

create_db:
	@docker exec -it library-api createdb --username=postgres --owner=postgres library-create_db
