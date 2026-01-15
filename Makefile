test:
	go mod tidy
	go test -v ./...

install:
	go install

lint:
	golangci-lint run ./...

build:
	go build -o bin/gengiff ./main.go