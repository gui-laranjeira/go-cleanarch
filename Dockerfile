FROM golang:1.21 as builder


ENV PORT 9000

WORKDIR /app

# Dependencies
COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o /library-api/ cmd/server/main.go

EXPOSE 8080

# Distroless image for small size (900MB to 4.25MB)
FROM  gcr.io/distroless/static-debian11 as final

WORKDIR /app
COPY --from=builder --chmod=777 /library-api/ ./
COPY --from=builder --chmod=777 /app/config.toml ./

ENTRYPOINT [ "/app/main" ]