.DEFAULT_GOAL := build

# Gets service_name and version from server.toml
SERVICE_NAME := $(shell sed -n 's/^ *SERVICE_NAME.*=.*"\([^"]*\)".*/\1/p' server.toml)
VERSION := $(shell sed -n 's/^ *VERSION.*=.*"\([^"]*\)".*/\1/p' server.toml)


build:
	@echo "Building for darwin/amd64"
	@GOARCH=amd64 GOOS=darwin go build -o ./build-output/$(SERVICE_NAME)-darwin main.go
	@echo "Building for linux/amd64"
	@GOARCH=amd64 GOOS=linux go build -o ./build-output/$(SERVICE_NAME)-linux main.go
	@echo "Build successful"

build-docker:
	docker build -t "$(SERVICE_NAME):$(VERSION)" -f ./Dockerfile .
 module name, and add new dependencies
build-docker-debug:
	docker build -t "$(SERVICE_NAME):$(VERSION)-debug" -f ./Dockerfile.debug .

format:
	@go fmt ./

run:
	@go run main.go

run-docker-debug:
	@docker run -p 8080:8080 -p 40000:40000 --security-opt="apparmor=unconfined" --cap-add=SYS_PTRACE -it "$(SERVICE_NAME):$(VERSION)-debug"

clean:
	@go clean
	@rm -r ./build-output/*

clean-docker:
	docker rmi -f ` docker images --all --format '{{.Repository}}:{{.Tag}}' | grep $(SERVICE_NAME)`

deps:
	@go mod download

tidy:
	@go mod tidy

vet:
	@go vet

.PHONY: build build-docker build-docker-debug format run run-docker-debug clean clean-docker deps tidy vet
