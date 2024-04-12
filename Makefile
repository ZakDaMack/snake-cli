build:
	go build -o bin/ cmd/snake

run: @build
	./bin/snake