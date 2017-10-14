include github.com/tj/make/golang

.DEFAULT_GOAL := generate

# Generate themes.
generate:
	@go generate ./...
.PHONY: generate

# Clean build artifacts.
clean:
	@echo "==> Clean"
	@rm -fr build
.PHONY: clean
