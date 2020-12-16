.PHONY: all clean

GO_CMD=go
BINARY_NAME=Week04

GO_BUILD=$(GO_CMD) build
GO_CLEAN=$(GO_CMD) clean
GO_TEST=$(GO_CMD) test
GO_GET=$(GO_CMD) get

# 注入编译信息
LD_FLAGS += -X "${BINARY_NAME}/global.Version=$(shell git describe --tags `git rev-list --tags --max-count=1`)"
LD_FLAGS += -X "${BINARY_NAME}/global.BuildTime=$(shell date '+%Y-%m-%d %H:%M:%S')"
LD_FLAGS += -X "${BINARY_NAME}/global.GitHash=$(shell git rev-parse HEAD)"
LD_FLAGS += -X "${BINARY_NAME}/global.GoVersion=$(shell go version | awk '{print $3}')"

all: swag clean build

build:
	CGO_ENABLED=0 $(GO_BUILD) -o $(BINARY_NAME) \
		-ldflags '$(LD_FLAGS)'

swag:
	swag init

build-sqlite:
	go build -tags sqlite3 -o $(BINARY_NAME) main.go

clean:
	$(GO_CLEAN)
	rm -f $(BINARY_NAME)

run:
	$(GO_BUILD) -o $(BINARY_NAME) ./...
	./$(BINARY_NAME)

# Cross compilation
linux:
	swag init
	GOOS=linux GOARCH=amd64 $(GO_BUILD) -o $(BINARY_NAME) \
		-ldflags '$(LD_FLAGS)'

#.PHONY: test
#test:
#	go test -v ./... -cover

#.PHONY: docker
#docker:
#	docker build . -t Week04:latest