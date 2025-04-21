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

# ✅ Install required tools
RUN apk add --no-cache \
    ca-certificates \
    make \
    ncurses

WORKDIR /app

# ✅ Copy binaries from previous stages
COPY --from=terminal-builder /app/email-client .
COPY --from=server-builder /app/server .

# ✅ Copy static HTML/JS frontend
RUN mkdir -p /app/web/static
COPY ./web/static /app/web/static

# ✅ Set execute permissions
RUN chmod +x /app/server /app/email-client

RUN ls -la /app

# ✅ Expose the WebSocket/HTTP server port
EXPOSE 8080

# ✅ Start the server
ENTRYPOINT ["/app/server"]
