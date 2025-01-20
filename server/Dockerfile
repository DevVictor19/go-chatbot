FROM golang:1.23.4-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY ./src ./src
COPY ./.env ./.env

WORKDIR /app/src

RUN go build -o /app/main .

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main .
COPY --from=builder /app/.env .

EXPOSE 8080

CMD ["./main"]
