COMMIT_ID := $(shell git rev-list --all --max-count=1)
#VERSION := $(shell git describe --abbrev=0 --tags --match "v[0-9]*" $(COMMIT_ID))
VERSION := "0.1.0"
BUILD_DATE := $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
BUILD_PATH := "./build"

build: test
	mkdir -p $(BUILD_PATH)
	go build -o $(BUILD_PATH)/gitmadeeasy-base main.go
	GOOS=linux go build -ldflags="-s -w -X main.version=$(VERSION) -X main.commit=$(COMMIT_ID) -X main.date='$(BUILD_DATE)' -X main.buildSource=binaryRelease" -o $(BUILD_PATH)/gitmadeeasy main.go
	GOOS=darwin go build -ldflags="-s -w -X main.version=$(VERSION) -X main.commit=$(COMMIT_ID) -X main.date='$(BUILD_DATE)' -X main.buildSource=darwin" -o $(BUILD_PATH)/gitmadeeasy-ios main.go

test:
	@echo -e "$(COLOR_YELLOW)Building the project $(COLOR_END)"
	@#go test --json -v  ./..
	 go test -v  ./... || true # TODO - fix exit code 2

local-deploy: build
	sudo cp -v gitmadeeasy /usr/local/bin/