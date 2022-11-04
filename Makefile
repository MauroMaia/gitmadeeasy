COMMIT_ID := $(shell git rev-list --all --max-count=1)
#VERSION := $(shell git describe --abbrev=0 --tags --match "v[0-9]*" $(COMMIT_ID))
VERSION := "0.1.0"
BUILD_DATE := $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

build:
	go build -o gitmadeeasy-base main.go
	GOOS=linux go build -ldflags="-s -w -X main.version=$(VERSION) -X main.commit=$(COMMIT_ID) -X main.date='$(BUILD_DATE)' -X main.buildSource=binaryRelease" -o gitmadeeasy main.go
	GOOS=darwin go build -ldflags="-s -w -X main.version=$(VERSION) -X main.commit=$(COMMIT_ID) -X main.date='$(BUILD_DATE)' -X main.buildSource=darwin" -o gitmadeeasy-ios main.go

test:
	@echo -e "$(COLOR_YELLOW)Building the project $(COLOR_END)"
	@#go test --json -v  ./..
	 go test -v  ./...

local-deploy: build
	sudo cp -v gitmadeeasy /usr/local/bin/