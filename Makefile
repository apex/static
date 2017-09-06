include github.com/tj/make/golang

# Generate test site.
#
# Just a temporary hack for now, you'll need to clone apex/up.
build: clean
	@echo "==> Generate"
	@go generate ./...
	@echo "==> Build"
	@go run cmd/static-docs/main.go -in ../up/docs -title Up -subtitle "Deploy serverless apps in seconds."
.PHONY: build

# Generate themes.
generate:
	@go generate ./...
.PHONY: generate

# Clean build artifacts.
clean:
	@echo "==> Clean"
	@rm -fr build
.PHONY: clean
