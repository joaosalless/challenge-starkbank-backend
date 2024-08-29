CMD_DIR=./cmd
MOCKS_DIR=./src/mocks
MIGRATIONS_DIR=./src/migrations
INTERFACES_DIR=./src/interfaces
MOCKGEN=$(shell go env GOPATH)/bin/mockgen

.PHONY: all run api cron test mocks clean docker-build docker-up docker-down

all: build

api:
	@echo "Starting a API..."
	go run $(CMD_DIR)/api/main.go

cron:
	@echo "Starting o Cron Job..."
	go run $(CMD_DIR)/cron/main.go

test:
	@echo "Running tests..."
	go test ./src/...

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
