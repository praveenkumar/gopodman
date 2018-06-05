# Go and compilation related variables
BUILD_DIR ?= out

ORG := github.com/praveenkumar
REPOPATH ?= $(ORG)/gopodman
GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)

ifeq ($(GOOS),windows)
	IS_EXE := .exe
endif
gopodman_BINARY ?= $(GOPATH)/bin/gopodman$(IS_EXE)

$(BUILD_DIR)/$(GOOS)-$(GOARCH):
	mkdir -p $(BUILD_DIR)/$(GOOS)-$(GOARCH)

$(BUILD_DIR)/darwin-amd64/gopodman: vendor $(BUILD_DIR)/$(GOOS)-$(GOARCH) ## Cross compiles the darwin executable and places it in $(BUILD_DIR)/darwin-amd64/gopodman
	CGO_ENABLED=0 GOARCH=amd64 GOOS=darwin go build --installsuffix "static" -o $(BUILD_DIR)/darwin-amd64/gopodma

$(BUILD_DIR)/linux-amd64/gopodman: vendor $(BUILD_DIR)/$(GOOS)-$(GOARCH) ## Cross compiles the linux executable and places it in $(BUILD_DIR)/linux-amd64/gopodman
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build --installsuffix "static" -o $(BUILD_DIR)/linux-amd64/gopodman

$(BUILD_DIR)/windows-amd64/gopodman.exe: vendor $(BUILD_DIR)/$(GOOS)-$(GOARCH) ## Cross compiles the windows executable and places it in $(BUILD_DIR)/windows-amd64/gopodman
	CGO_ENABLED=0 GOARCH=amd64 GOOS=windows go build --installsuffix "static" -o $(BUILD_DIR)/windows-amd64/gopodman.exe

.PHONY: vendor
vendor:
	dep ensure -v

$(BUILD_DIR):
	mkdir -p $(BUILD_DIR)

.PHONY: clean
clean:
	rm -rf $(BUILD_DIR)

.PHONY: cross ## Cross compiles all binaries
cross: $(BUILD_DIR)/darwin-amd64/gopodman $(BUILD_DIR)/linux-amd64/gopodman $(BUILD_DIR)/windows-amd64/gopodman.exe

.PHONY: build
build: $(BUILD_DIR)
	go build -installsuffix "static" -o $(BUILD_DIR)/gopodman
	chmod +x $(BUILD_DIR)/gopodman
