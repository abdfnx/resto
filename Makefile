.PHONY: build

TAG=$(shell git describe --abbrev=0 --tags)
DATE=$(shell go run ./build/date.go)

build:
		@go mod tidy && \
		go build -ldflags "-X main.version=$(TAG) -X main.versionDate=$(DATE)" -o resto

install: resto
		@mv resto /usr/local/bin

brc: # build resto container
		@docker build -t restohq/resto . && \
		docker push restohq/resto

brcwc: # build resto container with cache
		@docker pull restohq/resto:latest && \
		docker build -t restohq/resto --cache-from restohq/resto:latest . && \
		docker push restohq/resto

bfrc: # build full resto container
		@cd container && \
		docker build -t restohq/resto-full . && \
		docker push restohq/resto-full

bfrcwc: # build full resto container with cache
		@docker pull restohq/resto-full:latest && \
		docker build -t restohq/resto-full --cache-from restohq/resto-full:latest . && \
		docker push restohq/resto-full

ghrs:
		@node scripts/gh-resto/gh-rs.js
