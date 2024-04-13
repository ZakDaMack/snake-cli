all: build install

build:
	go build -o bin/snake cmd/main.go

run: build
	./bin/snake

install:
	mkdir -p ~/.local/bin
	cp bin/snake ~/.local/bin/

global-install:
	sudo cp bin/snake /usr/local/bin/

