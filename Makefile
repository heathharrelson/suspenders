BIN = suspenders
OUT_DIR = dist

GITHUB_URL = github.com/heathharrelson/suspenders
VERSION ?= $(shell cat VERSION)
COMMIT ?= $(shell git rev-parse --short HEAD)
BUILD_DATE ?= $(shell date -u +'%Y-%m-%dT%H:%M:%SZ')

ALL_ARCH = amd64 arm arm64 ppc64le s390x
ALL_PLATFORMS = $(addprefix linux/,$(ALL_ARCH))
ALL_BINARIES ?= $(addprefix $(OUT_DIR)/$(BIN)-, $(addprefix linux-,$(ALL_ARCH)))

GOOS ?= $(shell uname -s | tr A-Z a-z)
GOARCH ?= $(shell go env GOARCH)
LDFLAGS = \
	-X $(GITHUB_URL)/buildinfo.Version=$(VERSION) \
	-X $(GITHUB_URL)/buildinfo.Commit=$(COMMIT) \
	-X $(GITHUB_URL)/buildinfo.BuildDate=$(BUILD_DATE)

NPM_FLAGS = --prefix=ui

DOCKER_REPO ?= heathharrelson/suspenders
DOCKERHUB_URL ?= https://hub.docker.com/r/heathharrelson/suspenders
MANIFEST_BASE_TAG ?= $(DOCKER_REPO):$(VERSION)
MANIFEST_ARCH_TAGS ?= $(addprefix $(MANIFEST_BASE_TAG)-, $(ALL_ARCH))

all: assets suspenders

godeps: go.mod go.sum
	go mod download

jsdeps:
	npm install $(NPM_FLAGS)

devassets: jsdeps
	npm run build:dev $(NPM_FLAGS)

assets: jsdeps
	npm run build $(NPM_FLAGS)

$(BIN): godeps suspenders.go server.go
	go build -ldflags '$(LDFLAGS)' -o $(BIN) $(GITHUB_URL)

crossbuild: $(ALL_BINARIES)

$(OUT_DIR)/$(BIN): $(OUT_DIR)/$(BIN)-$(GOOS)-$(GOARCH)
	cp $(OUT_DIR)/$(BIN)-$(GOOS)-$(GOARCH) $(OUT_DIR)/$(BIN)

$(OUT_DIR)/$(BIN)-%:
	@echo ">> building $(OUT_DIR)/$(BIN)-$*"
	GOARCH=$(word 2,$(subst -, ,$(*:.exe=))) \
	GOOS=$(word 1,$(subst -, ,$(*:.exe=))) \
	CGO_ENABLED=0 \
	go build -ldflags '$(LDFLAGS)' -o $(OUT_DIR)/$(BIN)-$* $(GITHUB_URL)

image: $(OUT_DIR)/$(BIN)-$(GOOS)-$(GOARCH) Dockerfile
	docker build \
		--build-arg BINARY=$(BIN)-$(GOOS)-$(GOARCH) \
		--label "org.opencontainers.image.created=$(BUILD_DATE)" \
		--label "org.opencontainers.image.url=$(DOCKERHUB_URL)" \
		--label "org.opencontainers.image.source=https://$(GITHUB_URL)" \
		--label "org.opencontainers.image.version=$(VERSION)" \
		--label "org.opencontainers.image.revision=$(COMMIT)" \
		--label "org.opencontainers.image.licences=Apache-2.0" \
		-t $(DOCKER_REPO):$(VERSION)-$(GOARCH) \
		.
ifeq ($(GOARCH), amd64)
	docker tag $(DOCKER_REPO):$(VERSION)-$(GOARCH) $(DOCKER_REPO):$(VERSION)
endif

image-%:
	$(MAKE) GOOS=linux GOARCH=$* image
	docker push $(DOCKER_REPO):$(VERSION)-$*

manifest:
	DOCKER_CLI_EXPERIMENTAL=enabled \
	docker manifest create $(MANIFEST_BASE_TAG) $(MANIFEST_ARCH_TAGS)

manifest-%:
	DOCKER_CLI_EXPERIMENTAL=enabled \
	docker manifest annotate --arch "$*" "$(DOCKER_REPO):$(VERSION)" "$(DOCKER_REPO):$(VERSION)-$*"

docker: assets crossbuild $(addprefix image-,$(ALL_ARCH)) manifest $(addprefix manifest-,$(ALL_ARCH))
	DOCKER_CLI_EXPERIMENTAL=enabled \
	docker manifest push -p "$(DOCKER_REPO):$(VERSION)"

clean:
	go clean
	rm -rf dist
	rm -rf ui/dist
	rm -rf ui/static

superclean: clean
	rm -rf ui/node_modules

.PHONY: all assets clean crossbuild devassets docker godeps image image-% manifest manifest-% jsdeps superclean