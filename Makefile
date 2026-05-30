.PHONY: format
format:
	gofmt -w -s -l .

.PHONY: format-check
format-check:
	gofmt -s -d .
