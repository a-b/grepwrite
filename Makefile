test: build
	cat test/sample_output.txt | ./grepwrite

clean:
	rm ./grepwrite

grepwrite:
	go build

build: grepwrite

all: test build
