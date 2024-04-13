build:
	go build -o bin/snake cmd/main

run: build
	./bin/snake