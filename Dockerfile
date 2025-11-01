# ========= Stage 1: Build =========
FROM golang:1.25-bookworm AS builder

WORKDIR /app

# Copy go.mod and go.sum first for dependency caching
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copy the rest of your source code
COPY . .

# Build the Go app binary
RUN go build -v -o server .

# ========= Stage 2: Run =========
FROM debian:bookworm

# Install certificates for HTTPS (if needed)
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

# Copy built binary from builder stage
COPY --from=builder /app/server /usr/local/bin/server

# Copy .env (optional, or you can mount it)
COPY .env /usr/local/bin/.env

# Expose port (optional, matches your .env)
EXPOSE 8080

# Command to run your app
CMD ["server"]
