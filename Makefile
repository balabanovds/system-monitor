lint:
	golangci-lint run ./...

test-linux:
	./scripts/go14_linux.sh test -race ./...

test-darwin:
	go test -race  ./...

test: lint test-darwin test-linux

dl-client:
	cd client && npm install && cd ..

generate: dl-client
	protoc -I=./schema --go_out=plugins=grpc:./internal/api/ ./schema/metric_service.proto
	cd client && ./generate.sh && ./generate_classic.sh && cd ..

build: test generate
	docker-compose -f ./deployments/docker-compose.yml -p smonitor up -d --build

dev: test generate
	docker-compose -f ./deployments/docker-compose.yml -f ./deployments/docker-compose.dev.yml -p smonitor up -d --build

up: test generate
	docker-compose -f ./deployments/docker-compose.yml -p smonitor up -d

down:
	docker-compose -f ./deployments/docker-compose.yml -p smonitor down

logs:
	docker-compose -f ./deployments/docker-compose.yml -p smonitor logs
