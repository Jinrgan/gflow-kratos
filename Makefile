GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)
APP_RELATIVE_PATH=$(shell a=`basename $$PWD` && cd .. && b=`basename $$PWD` && echo $$b/$$a)
INTERNAL_PROTO_FILES=$(shell find internal -name *.proto)
API_PROTO_FILES=$(shell find api -name *.proto)
PROTO_PATH=./api/$(DOMAIN)/v1
CONF_PROTO_PATH=./app/$(DOMAIN)/internal/conf

.PHONY: init
# init env
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2@latest
	go install github.com/google/gnostic/cmd/protoc-gen-openapi@v0.6.1

.PHONY: errors
# generate errors code
errors:
	protoc --proto_path=$(PROTO_PATH) \
               --proto_path=./third_party \
               --go_out=paths=source_relative:$(PROTO_PATH) \
               --go-errors_out=paths=source_relative:$(PROTO_PATH) \
               $(DOMAIN)_error.proto

.PHONY: config
# generate internal proto
config:
	protoc --proto_path=$(CONF_PROTO_PATH) \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:$(CONF_PROTO_PATH) \
	       conf.proto

.PHONY: api
# generate api proto
api:
	protoc --proto_path=$(PROTO_PATH) \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:$(PROTO_PATH) \
 	       --go-http_out=paths=source_relative:$(PROTO_PATH) \
 	       --go-grpc_out=paths=source_relative:$(PROTO_PATH) \
 	       --openapi_out==paths=source_relative:$(PROTO_PATH) \
	       $(DOMAIN).proto

.PHONY: build
# build
build:
	mkdir -p bin/ && go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./...

.PHONY: generate
# generate
generate:
	go generate ./...

.PHONY: all
# generate all
all:
	make api;
	make errors;
	make config;
	make generate;

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
