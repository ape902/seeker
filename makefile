TIMESTAMP ?= $(shell date +%Y%m%d%H%M%S)
VERSION=v0.0.1
GOOS="linux"
GOARCH="amd64"


.PHONY: build-agent
build-agent:
	@echo ">> building AGENT binaries..."
	@GOOS=$(GOOS) GOARCH=$(GOARCH) go build -trimpath -ldflags "-s -w -X github.com/ape902/seeker/pkg/tools/versionx.Version=$(VERSION) -X github.com/ape902/seeker/pkg/tools/versionx.BuildTime=$(TIMESTAMP)" -o agent.$(GOARCH).$(VERSION).$(TIMESTAMP) cmd/agent/agent.go

.PHONY: build-seeker
build-seeker:
	@echo ">> building SEEKER binaries..."
	@GOOS=$(GOOS) GOARCH=$(GOARCH) go build -trimpath -ldflags "-s -w -X github.com/ape902/seeker/pkg/tools/versionx.Version=$(VERSION) -X github.com/ape902/seeker/pkg/tools/versionx.BuildTime=$(TIMESTAMP)" -o seeker.$(GOARCH).$(VERSION).$(TIMESTAMP) cmd/seeker/seeker.go

.PHONY: build-all
build-all: build-seeker build-agent
