# Start from the official Go image as a builder
FROM golang:1.22 AS builder

# Set environment variables
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

# Create app directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire project
COPY . .

# Build the Go app
RUN go build -o email-client ./cmd

# Final lightweight stage
FROM alpine:latest

# Add certificate bundle (optional but useful for net/http)
RUN apk --no-cache add ca-certificates

# Set working directory
WORKDIR /app

# Copy built binary from builder
COPY --from=builder /app/email-client .

# Set the entrypoint
ENTRYPOINT ["./email-client"]
