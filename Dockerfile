FROM golang:latest AS builder

WORKDIR /tgbot

COPY bot ./bot
COPY sqlite ./sqlite

WORKDIR /tgbot/bot
RUN GO111MODULE=off go build -o csc1038bot .

FROM debian:bookworm-slim
RUN apt-get update && apt-get install -y ca-certificates libssl3 && rm -rf /var/lib/apt/lists/*

WORKDIR /tgbot
COPY --from=builder /tgbot/bot/csc1038bot .

CMD ["./csc1038bot"]