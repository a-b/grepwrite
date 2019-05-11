GOBIN=bin

test: build
	cat test/sample_input.txt | bin/gw

clean:
	find bin -mindepth 1 -type f -name gw -delete
	find tmp -mindepth 1 ! -name .gitkeep -delete

grepwrite:
	go build -o bin/gw cmd/gw/main.go

build: grepwrite

all: clean build
