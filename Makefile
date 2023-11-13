
# Git information
GIT_COMMIT = $(shell git rev-parse HEAD)
#GIT_COMMIT_HASH    = $(shell git rev-parse --short HEAD)
GIT_COMMIT_HASH    = $(shell git rev-parse HEAD)
GIT_TAG    = $(shell git describe --tags --abbrev=0 --exact-match 2>/dev/null)
GIT_TREESTATE  = $(shell test -n "`git status --porcelain`" && echo "dirty" || echo "clean")
BUILDDATE = $(shell date -u +'%Y-%m-%dT%H:%M:%SZ')

LDFLAGS += -X github.com/zdz1715/go-pkg-version.version=$(GIT_TAG)
LDFLAGS += -X github.com/zdz1715/go-pkg-version.gitCommit=$(GIT_COMMIT_HASH)
LDFLAGS += -X github.com/zdz1715/go-pkg-version.gitTreeState=$(GIT_TREESTATE)
LDFLAGS += -X github.com/zdz1715/go-pkg-version.buildDate=$(BUILDDATE)

##@ General

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "Usage:\n  make \033[36m<target>\033[0m\n"} /^[%a-zA-Z_._0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)


##@ Build

.PHONY: build
build: ## Build binary.
	go build -ldflags "$(LDFLAGS)" -o not-use-command-version example/not-use-command/main.go
