BINARY_NAME=harmoni

build:
	go build -o bin/harmoni cmd/harmoni/main.go cmd/harmoni/ScaleInput.go cmd/harmoni/ScaleNotesDisplay.go

run:	build
	./bin/harmoni
