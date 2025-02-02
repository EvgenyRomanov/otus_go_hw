BIN_CALENDAR := "./bin/calendar"
BIN_SCHEDULER := "./bin/calendar_scheduler"
BIN_SENDER := "./bin/calendar_sender"
DOCKER_CALENDAR_IMG="calendar:develop"
DOCKER_SCHEDULER_IMG="scheduler:develop"
DOCKER_SENDER_IMG="sender:develop"

GIT_HASH := $(shell git log --format="%h" -n 1)
LDFLAGS := -X main.release="develop" -X main.buildDate=$(shell date -u +%Y-%m-%dT%H:%M:%S) -X main.gitHash=$(GIT_HASH)

build-calendar:
	go build -v -o $(BIN_CALENDAR) -ldflags "$(LDFLAGS)" ./cmd/calendar

run-calendar: build
	$(BIN_CALENDAR) -config ./configs/calendar_config.toml

build-scheduler:
	go build -v -o $(BIN_SCHEDULER) -ldflags "$(LDFLAGS)" ./cmd/scheduler

run-scheduler: build-scheduler
	$(BIN_SCHEDULER) -config ./configs/scheduler_config.toml

build-sender:
	go build -v -o $(BIN_SENDER) -ldflags "$(LDFLAGS)" ./cmd/sender

run-sender: build-sender
	$(BIN_SENDER) -config ./configs/sender_config.toml

build: build-calendar build-scheduler build-sender

build-img-calendar:
	docker build \
		--build-arg=LDFLAGS="$(LDFLAGS)" \
		-t $(DOCKER_CALENDAR_IMG) \
		-f build/calendar/Dockerfile .

build-img-scheduler:
	docker build \
		--build-arg=LDFLAGS="$(LDFLAGS)" \
		-t $(DOCKER_SCHEDULER_IMG) \
		-f build/scheduler/Dockerfile .

build-img-sender:
	docker build \
		--build-arg=LDFLAGS="$(LDFLAGS)" \
		-t $(DOCKER_SENDER_IMG) \
		-f build/sender/Dockerfile .

build-img: build-img-calendar build-img-scheduler build-img-sender

run-img-sender: build-img-sender
	docker run $(DOCKER_SENDER_IMG)

run-img-scheduler: build-img-scheduler
	docker run $(DOCKER_SCHEDULER_IMG)

run-img-calendar: build-img-calendar
	docker run $(DOCKER_CALENDAR_IMG)

up:
	docker compose -f "./build/docker-compose.yml" up --no-deps --build

down:
	docker compose -f "./build/docker-compose.yml" down

integration-tests:
	docker compose -f "./build/docker-compose.yml" --profile test up -d --no-deps --build
	docker compose -f "./build/docker-compose.yml" --profile test exec -it app-integration-test go test -v ./test/integration_test.go -tags integration
	docker compose -f "./build/docker-compose.yml" --profile test down

generate:
	rm -rf internal/server/pb/* && \
	protoc \
        --proto_path=api/ \
        --go_out=internal/server/pb \
        --go-grpc_out=internal/server/pb \
        api/*.proto

version-calendar: build
	$(BIN_CALENDAR) version

scheduler-version: build
	$(BIN_SCHEDULER) version

sender-version: build
	$(BIN_SENDER) version

test:
	go test -race ./internal/...

install-lint-deps:
	(which golangci-lint > /dev/null) || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.63.4

lint: install-lint-deps
	golangci-lint run ./...

.PHONY: build-calendar run-calendar build-scheduler run-scheduler build-sender run-sender build build-img-calendar build-img-scheduler build-img-sender build-img up down generate version-calendar scheduler-version sender-version test install-lint-deps lint run-img-scheduler run-img-calendar run-img-sender integration-tests
