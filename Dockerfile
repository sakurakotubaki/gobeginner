# ビルドステージ
FROM golang:1.23 AS builder
WORKDIR /app
COPY . .
RUN go build -o main

# 実行ステージ
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
CMD ["./main"]