FROM golang:1.21 as builder


ENV PORT 9000

WORKDIR /app 

# Dependencies
COPY go.mod go.sum ./
RUN go mod download


COPY . ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /library-api main.go

EXPOSE 9000

# Distroless image for small size (900MB to 4.25MB)
FROM  gcr.io/distroless/static-debian11 as final

COPY --from=builder /library-api /library-api

ENTRYPOINT [ "/library-api" ]