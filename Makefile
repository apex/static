include github.com/tj/make/golang

# Generate themes.
generate:
	@go generate ./...
.PHONY: generate
