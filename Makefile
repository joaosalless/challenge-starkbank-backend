CMD_DIR=./cmd
MOCKS_DIR=./src/mocks
MIGRATIONS_DIR=./src/migrations
INTERFACES_DIR=./src/interfaces
MOCKGEN=$(shell go env GOPATH)/bin/mockgen

.PHONY: all run api schedule test mocks clean docker-build docker-up docker-down

all: build

api:
	@echo "Starting API..."
	go run $(CMD_DIR)/api/main.go

schedule:
	@echo "Starting schedule..."
	go run $(CMD_DIR)/schedule/main.go

test:
	@echo "Running tests..."
	go test ./src/...

coverage:
	@echo "Generating test coverage report..."
	go test ./src/... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

mocks:
	@echo "Generating mocks..."
	@for file in $(INTERFACES_DIR)/*.go; do \
		mock_name=$$(basename $$file .go)_mock.go; \
		$(MOCKGEN) -source=$$file -destination=$(MOCKS_DIR)/$$mock_name -package=mocks && \
		echo "Generated $(MOCKS_DIR)/$$mock_name" ; \
	done

clean:
	@echo "Cleaning generated files..."
	go clean

build:
	@echo "Building project..."
	go build -o bin/api $(CMD_DIR)/api/main.go
	go build -o bin/cron $(CMD_DIR)/cron/main.go

docker-build:
	@echo "Build docker image..."
	docker-compose build

docker-up:
	@echo "Up containers..."
	docker-compose up --build

docker-down:
	@echo "Down containers..."
	docker-compose down
