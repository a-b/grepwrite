test:
	@go build
	@cat test/sample_output.txt | ./grepwrite
