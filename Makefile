.DEFAULT_GOAL := run

.PHONY:run

build:
	go build -o bin/http-go src/*.go

run: build
	./bin/http-go

clean:
	go clean
