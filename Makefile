.PHONY: setup
setup:
	@go install golang.org/dl/gotip@latest
	@gotip download

.PHONY: test
test:
	@gotip test ./... -v -count=1