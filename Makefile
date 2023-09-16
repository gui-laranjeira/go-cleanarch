build:
	go build -o bin/main cmd/server/main.go

run:
	go run cmd/server/main.go

compile:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin cmd/server/main.go
