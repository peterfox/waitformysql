# Go parameters
GOCMD=go
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
	mkdir $(BUILD_PATH)
	$(GOBUILD) -o $(BUILD_PATH)/$(BINARY_NAME) -v $(SOURCE_PATH)
clean:
	$(GOCLEAN)
	rm -rf $(BUILD_PATH)
deps:
	$(GOGET) -u github.com/go-sql-driver/mysql