# Build stage
FROM golang:1.24 AS builder

WORKDIR /app

# Copy go mod files
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy source code
COPY . ./

# Build the application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o /app/ordersystem ./cmd/ordersystem

# Run stage
FROM debian:bookworm-slim

WORKDIR /app

# Copy binary and config from builder
COPY --from=builder /app/ordersystem .
COPY /cmd/ordersystem/.env .

RUN chmod +x /app/ordersystem

CMD ["./ordersystem"]
