APP_NAME = service-customer
MAIN_PATH = ./cmd
BUILD_DIR = ./bin

.PHONY: all
all: build

.PHONY: build
build: tidy
	go build -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_PATH)

.PHONY: run
run:
	go run $(MAIN_PATH)

.PHONY: test
test: mock
	go test ./... -v

.PHONY: mock
mock:
	go generate ./...

.PHONY: tidy
tidy: mock
	go fmt ./...
	go mod tidy