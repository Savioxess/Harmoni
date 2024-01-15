BINARY_NAME=harmoni

build:
	go build -o bin/harmoni cmd/harmoni/main.go

run:	build
	./bin/harmoni
