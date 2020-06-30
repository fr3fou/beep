BINARY_NAME=megolovania
MUSIC_FILE=memory.gsc


.PHONY: build run play

all: build run play

re-run: run play

build: 
	go build -o $(BINARY_NAME) ./cmd/generate/

run:
	./$(BINARY_NAME) $(MUSIC_FILE)

play:
	ffplay $(MUSIC_FILE) -autoexit -showmode 1 -f f64le -ar 48000
