APP_NAME = service-customer
MAIN_PATH = ./cmd
BUILD_DIR = ./bin
TEST_COMPOSE_FILE = docker-compose.test.yml
TEST_DB_DSN = "postgres://postgres:postgres@localhost:5431/test_db?sslmode=disable"

.PHONY: all
all: build

.PHONY: build
build: tidy
	go build -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_PATH)

.PHONY: run
run:
	go run $(MAIN_PATH)

.PHONY: test
test: mock test-env-up
	TEST_DATABASE_URL=$(TEST_DB_DSN) go test ./... -v -count=1
	@$(MAKE) test-env-down

.PHONY: test-env-up
test-env-up:
	docker-compose -f $(TEST_COMPOSE_FILE) up -d

.PHONY: test-env-down
test-env-down:
	docker-compose -f $(TEST_COMPOSE_FILE) down

.PHONY: mock
mock:
	go generate ./...

.PHONY: tidy
tidy: mock
	go fmt ./...
	go mod tidy