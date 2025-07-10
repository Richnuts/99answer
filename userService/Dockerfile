FROM golang:1.23.3 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN go build -o user_service

FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/user_service .

EXPOSE 9001

CMD ["./user_service"]