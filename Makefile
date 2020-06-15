BINARY_NAME=megolovania

.PHONY: build run play

all: build run play

re-run: run play

build: 
	go build -o $(BINARY_NAME) main.go

run:
	./$(BINARY_NAME)

play:
	ffplay megolovania.bin -autoexit -showmode 1 -f f64le -ar 48000
