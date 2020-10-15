lint:
	golangci-lint run ./...

test-linux:
	./scripts/go14_linux.sh test -race -count 100 ./...

test-darwin:
	go test -race -count 100 ./...

test: lint 
	go test -race -count 100 ./...

generate:
	protoc -I=./schema --go_out=plugins=grpc:./internal/api/ ./schema/metric_service.proto
	cd client && ./generate.sh && cd ..
