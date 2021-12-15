.DEFAULT_GOAL := all

## parameters

NAME              ?= orbitdb
PROJECT           ?= debridgefinance
VERSION           ?= development
ENV               ?= dev
IMAGE_REPOSITORY  ?= docker.io/$(PROJECT)/$(NAME)
IMAGE_TAG         ?= $(VERSION)
DOCKER_USER       ?=
DOCKER_PASSWORD   ?=
SEMVER    		  ?= v0.1.0


K8S_CLUSTER   ?= kubernetes-$(ENV)
K8S_NAMESPACE ?= $(PROJECT)-$(ENV)

AUTH_COOKIE_SECRET ?=
AUTH_CLIENT_ID     ?=
AUTH_CLIENT_SECRET ?=

PARALLEL_JOBS ?= 8
NIX_OPTS      ?=

export GOFLAGS ?= -mod=vendor
.PHONY: all
all: test lint build # test, check and build all cmds

.PHONY: help
help: # print defined targets and their comments
	@grep -Po '^[a-zA-Z%_/\-\s]+:+(\s.*$$|$$)' $(MAKEFILE_LIST)  \
		| sort                                                   \
		| sed 's|:.*#|#|;s|#\s*|#|'                              \
		| column -t -s '#' -o ' | '

.PHONY: build
build: # build orbitdb binary
	go build -o orbitdb ./main.go


.PHONY: docker-build
docker-build: # build docker image
	docker build -t $(IMAGE_REPOSITORY):$(SEMVER) .