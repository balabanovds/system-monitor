install-lint-deps:
	(which golangci-lint > /dev/null) || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.30.0

lint: install-lint-deps
	golangci-lint run ./...

test-linux-docker:
	./scripts/go14_linux.sh test -race ./...

test: lint
	go test -race -count 1  ./...

dl-client:
	cd client && npm install && cd ..

generate: dl-client
	protoc -I=./schema --go_out=plugins=grpc:./internal/api/ ./schema/metric_service.proto
	cd client && ./generate.sh && ./generate_classic.sh && cd ..

build:
	go build -o system-monitor ./cmd

rebuild: test generate
	docker-compose -f ./deployments/docker-compose.yml -p smonitor up -d --build

dev: test generate
	docker-compose -f ./deployments/docker-compose.yml -f ./deployments/docker-compose.dev.yml -p smonitor up -d --build

up: test generate
	docker-compose -f ./deployments/docker-compose.yml -p smonitor up -d

down:
	docker-compose -f ./deployments/docker-compose.yml -p smonitor down

logs:
	docker-compose -f ./deployments/docker-compose.yml -p smonitor logs
