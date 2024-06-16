GIT_COMMIT := $(shell git rev-parse --short --verify HEAD)
GIT_SHA := $(shell git rev-parse --verify HEAD)
VERSION := $(shell cat .version)

.PHONY: bundle clean test

clean:
	rm -rf bin

bundle: clean
	@echo "Creating bundle for version $(GIT_SHA) on commit $(GIT_COMMIT)"
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-s -w -X main.Version=$(GIT_COMMIT)" -o bin/main-linux-amd64-$(GIT_SHA)
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -a -installsuffix cgo -ldflags="-s -w -X main.Version=$(GIT_COMMIT)" -o bin/main-linux-arm64-$(GIT_SHA)
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -a -installsuffix cgo -ldflags="-s -w -X main.Version=$(GIT_COMMIT)" -o bin/main-darwin-arm64-$(GIT_SHA)
	echo $(GIT_SHA) > .version

	@git add bin/*
	@git add .version
	@git commit -m "Bundling version $(VERSION) on commit $(GIT_COMMIT)"

test:
	@echo "Testing commit: $(GIT_COMMIT)"
	go test -v ./...

action:
	@echo "Running action"
	@echo "Commit: $(GIT_COMMIT)"


	INPUT_GCP_PROJECT_ID=project-1234 \
    INPUT_ENVIRONMENT=production \
    go run main.go

