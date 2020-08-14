GO         ?= go
GOMOD      ?= $(GO) mod
GOBUILD    ?= $(GO) build
GOHOSTOS   ?= $(shell $(GO) env GOHOSTOS)
GOHOSTARCH ?= $(shell $(GO) env GOHOSTARCH)

GO_VERSION        ?= $(shell $(GO) version)
GO_VERSION_NUMBER ?= $(word 3, $(GO_VERSION))

NPM ?= npm

.PHONY: all
all: godeps assets suspenders

.PHONY: godeps
godeps:
	$(GOMOD) download

.PHONY: devassets
devassets:
	npm run build:dev --prefix=server

.PHONY: assets
assets:
	npm install --prefix=server
	npm run build:prod --prefix=server

suspenders:
	$(GOBUILD)

.PHONY: run
run: assets
	$(GO) run suspenders.go

.PHONY: clean
clean:
	$(GO) clean
	rm -rf server/static

.PHONY: superclean
superclean: clean
	rm -rf server/node_modules
