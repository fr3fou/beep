BINARY_NAME=megolovania
MUSIC_FILE=memory.gsc
MUSIC_FILE_WAV=memory.wav

.PHONY: build run play wav

all: build run play

re-run: run play

build: 
	go build -o $(BINARY_NAME) ./cmd/generate/

run:
	./$(BINARY_NAME) $(MUSIC_FILE)

play:
	ffplay $(MUSIC_FILE) -autoexit -showmode 1 -f f64le -ar 48000

wav:
	ffmpeg -f f64le -ar 48000 -i $(MUSIC_FILE) $(MUSIC_FILE_WAV)    