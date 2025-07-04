FROM golang:1.24

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY ./auth-service/go.mod ./auth-service/go.sum ./auth-service/

WORKDIR /app/auth-service

RUN go mod tidy

COPY ./auth-service/. .

COPY ./auth-service/air.toml ./

EXPOSE 8080

CMD ["air"]
