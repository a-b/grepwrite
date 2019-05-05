test:
	@go build
	@cat sample_output.txt | ./grepwriter
