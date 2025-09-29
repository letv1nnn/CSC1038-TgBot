FROM golang:latest AS builder

WORKDIR /tgbot

COPY go.mod go.sum ./
RUN go mod download

COPY main.go .
COPY botui ./botui
COPY sqlite ./sqlite

RUN go build -o csc1038bot main.go

FROM debian:bookworm-slim
RUN apt-get update && apt-get install -y ca-certificates libssl3 && rm -rf /var/lib/apt/lists/*

WORKDIR /tgbot
COPY --from=builder /tgbot/csc1038bot .
COPY --from=builder /tgbot/sqlite/ ./sqlite/

CMD ["./csc1038bot"]