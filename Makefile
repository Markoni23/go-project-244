test:
	go mod tidy
	go test -v ./...

install:
	go install

lint:
	golangci-lint run ./...

build:
	go build -o bin/gengiff ./cmd/gendiff/main.go

run:
	go run cmd/gendiff/main.go