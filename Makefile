@PHONY: all

all: clean build

build: ranking

ranking:
	go build -v ./cmd/ranking

clean:
	rm -f ranking

run: ranking
	./ranking < sample_input.txt

