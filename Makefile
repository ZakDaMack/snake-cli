all: build install

build:
	go build -o bin/snake cmd/main.go

run: build
	./bin/snake

install:
	sudo cp bin/snake /usr/local/bin/
