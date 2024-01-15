BINARY_NAME=harmoni
PROJECT=cmd/harmoni

build:
	go build -o bin/harmoni ${PROJECT}/main.go ${PROJECT}/ScaleInput.go ${PROJECT}/ScaleNotesDisplay.go ${PROJECT}/ChordsDisplay.go

run:	build
	./bin/harmoni
