# Go parameters
GOCMD=go
GOX=gox
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOINSTALL=$(GOCMD) install
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=waitformysql
SOURCE_PATH=./cmd/$(BINARY_NAME)/main.go
BUILD_PATH=./build

all: deps clean build
build:
	$(GOX) -os="linux darwin windows" -arch="amd64" \
	-output="$(BUILD_PATH)/$(BINARY_NAME)_{{.OS}}_{{.Arch}}" -verbose ./...
clean:
	$(GOCLEAN)
	rm -rf $(BUILD_PATH)
deps:
	$(GOGET) -u github.com/go-sql-driver/mysql
	$(GOGET) -u github.com/mitchellh/gox