.DEFAULT_GOAL := build

clean:
	rm -rf bin

build: clean
	go build -o bin/deepn ./cmd/main.go