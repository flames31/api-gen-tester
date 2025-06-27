# syntax=docker/dockerfile:1

# Step 1: Build the Go binary
FROM golang:1.23.3-bullseye AS builder

WORKDIR /app

# Copy go.mod and go.sum files first (for caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app .

# Step 2: Create a minimal image with the binary
FROM debian:bullseye-slim

# Run updates and install certs for HTTPS
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

WORKDIR /root/
RUN mkdir -p /root/logs

# Copy the binary and other files from the builder stage
COPY --from=builder /app/app .
COPY --from=builder /app/logs /app/logs
COPY --from=builder /app/sample.json .
COPY --from=builder /app/.env .

# Set binary as entrypoint
ENTRYPOINT ["./app"]