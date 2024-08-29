FROM golang:1.22-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /app/bin/api ./cmd/api/main.go
RUN go build -o /app/bin/cron ./cmd/cron/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/bin/api /app/bin/api
COPY --from=builder /app/bin/cron /app/bin/cron

ARG DB_HOST
ARG DB_USER
ARG DB_PASSWORD
ARG DB_NAME
ARG DB_PORT

ENV DB_HOST=${DB_HOST:-db}
ENV DB_USER=${DB_USER:-postgres}
ENV DB_PASSWORD=${DB_PASSWORD:-postgres}
ENV DB_NAME=${DB_NAME:-postgres}
ENV DB_PORT=${DB_PORT:-5432}

CMD ["/app/bin/api"]
