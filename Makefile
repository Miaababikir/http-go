.DEFAULT_GOAL := run

.PHONY:run

build:
	go build -o bin/http-go cmd/main.go

run: build
	./bin/http-go

clean:
	go clean
