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

# Release binaries to GitHub.
release: build
	@echo "==> Releasing"
	@goreleaser -p 1 --rm-dist -config .goreleaser.yml
	@echo "==> Complete"
.PHONY: release
