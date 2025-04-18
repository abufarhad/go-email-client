.PHONY: build run clean image


APP_NAME = email-client
DOCKER_IMAGE = $(APP_NAME):latest

build:
	docker build -t $(DOCKER_IMAGE) .

run:
	docker run --rm -it $(DOCKER_IMAGE)

clean:
	docker rmi -f $(DOCKER_IMAGE)

image:
	docker images | grep $(APP_NAME)
