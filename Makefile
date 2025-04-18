.PHONY: build run clean

build:
	go build -o email-client ./cmd

run:
	go run cmd/main.go

clean:
	rm -f email-client
