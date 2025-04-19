# ---------- Stage 1: Build terminal app ----------
FROM golang:1.22 AS terminal-builder
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o email-client ./cmd

# ---------- Stage 2: Build WebSocket server ----------
FROM golang:1.22 AS server-builder
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o server ./web


# ---------- Final Stage ----------
FROM alpine:latest

RUN apk add --no-cache ca-certificates

WORKDIR /app

# Copy both binaries and make sure they're executable
COPY --from=terminal-builder /app/email-client .
COPY --from=server-builder /app/server .
COPY ./web/static /app/web/static


# Ensure binary has exec perms (just in case)
RUN chmod +x /app/server /app/email-client

EXPOSE 8080
ENTRYPOINT ["/app/server"]

