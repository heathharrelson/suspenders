GO         ?= go
GOMOD      ?= $(GO) mod
GOBUILD    ?= $(GO) build
GOHOSTOS   ?= $(shell $(GO) env GOHOSTOS)
GOHOSTARCH ?= $(shell $(GO) env GOHOSTARCH)

GO_VERSION        ?= $(shell $(GO) version)
GO_VERSION_NUMBER ?= $(word 3, $(GO_VERSION))

NPM ?= npm

.PHONY: all
all: assets suspenders

.PHONY: godeps
godeps: go.mod go.sum
	$(GOMOD) download

.PHONY: devassets
devassets:
	npm run build:dev --prefix=ui

.PHONY: assets
assets:
	npm install --prefix=ui
	npm run build:prod --prefix=ui

suspenders: godeps suspenders.go server.go
	$(GOBUILD)

.PHONY: run
run: assets
	$(GO) run suspenders.go

.PHONY: clean
clean:
	$(GO) clean
	rm -rf ui/static

.PHONY: superclean
superclean: clean
	rm -rf ui/node_modules
