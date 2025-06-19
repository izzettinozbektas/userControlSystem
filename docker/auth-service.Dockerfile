FROM golang:1.22

WORKDIR /app

COPY ./auth-service/go.mod ./go.mod
COPY ./auth-service/go.sum ./go.sum

RUN go mod download

COPY ./auth-service/. .

RUN go build -o auth-service ./cmd/main.go

EXPOSE 8080

CMD ["./auth-service"]
