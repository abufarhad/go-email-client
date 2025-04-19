APP_NAME=email-client
APP_SRC=./cmd

SERVER_NAME=server
SERVER_SRC=./web

PORT=8080
IMAGE_NAME=$(APP_NAME)-web

.PHONY: build run docker-build docker-run clean

# Build both binaries into project root
build:
	go build -o $(APP_NAME) $(APP_SRC)
	go build -o $(SERVER_NAME) $(SERVER_SRC)

# Run the WebSocket server locally
run: build
	./$(SERVER_NAME)

# Build Docker image
docker-build:
	docker build -t $(IMAGE_NAME) .

# Run Docker container
docker-run:
	docker run -p $(PORT):8080 $(IMAGE_NAME)

# Clean up binaries
clean:
	docker rmi -f $(IMAGE_NAME)
