TIMESTAMP ?= $(shell date +%Y%m%d%H%M%S)
VERSION=v0.0.1
GOOS="linux"
GOARCH="amd64"


.PHONY: build-agent
build-agent:
	@echo ">> building AGENT binaries..."
	@cd cmd/agent && GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o agent.$(GOARCH).$(VERSION).$(TIMESTAMP)

.PHONY: build-seeker
build-seeker:
	@echo ">> building SEEKER binaries..."
	@cd cmd/seeker && GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o seeker.$(GOARCH).$(VERSION).$(TIMESTAMP)

.PHONY: build-all
build-all: build-seeker build-agent
