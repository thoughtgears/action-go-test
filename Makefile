GIT_COMMIT := $(shell git rev-parse --short --verify HEAD)
GIT_SHA := $(shell git rev-parse --verify HEAD)
VERSION := $(shell cat .version)

.PHONY: bundle clean test

clean:
	rm -rf bin

bundle: clean
	@echo "Creating bundle for version $(VERSION) on commit $(GIT_COMMIT)"
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-s -w -X main.Version=$(GIT_SHA)" -o bin/main-linux-amd64-v1
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -a -installsuffix cgo -ldflags="-s -w -X main.Version=$(GIT_SHA)" -o bin/main-linux-arm64-v1
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -a -installsuffix cgo -ldflags="-s -w -X main.Version=$(GIT_SHA)" -o bin/main-darwin-arm64-v1
	echo $(GIT_SHA) > .version

	@git add bin/*
	@git add .version
	@git commit -m "Bundling version $(VERSION) on commit $(GIT_COMMIT)"

test:
	@echo "Testing commit: $(GIT_COMMIT)"
