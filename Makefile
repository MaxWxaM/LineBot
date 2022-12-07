.PHONY: docker build check run help start
BIN_FILE=linebot
docker:
	docker-compose up -d
wire:
	wire ./...
build:
	@go build -o "./bin/${BIN_FILE}" ./init
check:
	@go fmt ./init
	@go vet ./init
run:
	./bin/"${BIN_FILE}"
start: docker wire build run