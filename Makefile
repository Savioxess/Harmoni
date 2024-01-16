BINARY=bin/harmoni
PROJECT=cmd/harmoni

build:
	go build -o ${BINARY} ${PROJECT}/main.go ${PROJECT}/ScaleInput.go ${PROJECT}/ScaleNotesDisplay.go ${PROJECT}/ChordsDisplay.go

run:	build
	./${BINARY}
