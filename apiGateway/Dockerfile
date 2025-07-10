FROM golang:1.23.3 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN go build -o ApiGateway

FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/ApiGateway .

EXPOSE 9001

CMD ["./ApiGateway"]