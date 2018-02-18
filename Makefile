APPNAME ?= postit-lambda-api
ENV ?= dev
ENV_NO ?= 1

SOURCE_FILES?=$$(go list ./... | grep -v /vendor/)

default: build package deploy

# Install all the build and lint dependencies
setup:
	go get -u github.com/alecthomas/gometalinter
	go get -u github.com/golang/dep/cmd/dep
	go get -u github.com/pierrre/gotestcover
	go get -u golang.org/x/tools/cmd/cover
	go get -u github.com/golang/protobuf/...
	go get -u -v github.com/twitchtv/twirp/...
	gometalinter --install
	dep ensure
.PHONY: setup

# Run all the tests
test:
	@gotestcover $(TEST_OPTIONS) -covermode=atomic -coverprofile=coverage.txt $(SOURCE_FILES) -run $(TEST_PATTERN) -timeout=2m
.PHONY: test

# Run all the linters
lint:
	gometalinter --vendor ./...
.PHONY: lint

# generate service code
generate:
	protoc --proto_path=$GOPATH/src:. --twirp_out=. --go_out=. ./service.proto
.PHONY: generate

# build the lambda binary
build:
	GOOS=linux GOARCH=amd64 go build -o main ./cmd/lambda
.PHONY: build

clean:
	rm ./main
.PHONY: clean

# Run all the tests and code checks
ci: setup test lint
.PHONY: ci

# run lambda locally using sam-local
local-lambda: build
	go get -u github.com/awslabs/aws-sam-local
	aws-sam-local local start-api --template deploy.sam.yml -n .env.json
.PHONY: local

install:
	go install ./cmd/postit-server
	go install ./cmd/postit

# package up the lambda and upload it to S3
package:
	echo "Running as: $(shell aws sts get-caller-identity --query Arn --output text)"
	aws cloudformation package \
		--template-file deploy.sam.yml \
		--output-template-file deploy.out.yml \
		--s3-bucket $(S3_BUCKET) \
		--s3-prefix sam
.PHONY: package

# deploy the lambda
deploy:
	aws cloudformation deploy \
		--template-file deploy.out.yml \
		--capabilities CAPABILITY_IAM \
		--stack-name $(APPNAME)-$(ENV)-$(ENV_NO) \
		--parameter-overrides EnvironmentName=$(ENV) EnvironmentNumber=$(ENV_NO)
.PHONY: deploy