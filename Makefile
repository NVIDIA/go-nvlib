# Copyright (c) 2021, NVIDIA CORPORATION.  All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

include versions.mk

DOCKER ?= docker

PCI_IDS_URL ?= https://pci-ids.ucw.cz/v2.2/pci.ids

CHECK_TARGETS := lint
TARGETS := binary build all check fmt assert-fmt generate lint vet test coverage
DOCKER_TARGETS := $(patsubst %,docker-%, $(TARGETS))
.PHONY: $(TARGETS) $(DOCKER_TARGETS) vendor check-vendor

GOOS := linux

build:
	GOOS=$(GOOS) go build ./...

all: check build binary
check: $(CHECK_TARGETS)

vendor:
	go mod tidy
	go mod vendor
	go mod verify

check-vendor: vendor
	git diff --exit-code HEAD -- go.mod go.sum vendor

# Apply go fmt to the codebase
fmt:
	go list -f '{{.Dir}}' $(MODULE)/... \
		| xargs gofmt -s -l -w

assert-fmt:
	go list -f '{{.Dir}}' $(MODULE)/... \
		| xargs gofmt -s -l > fmt.out
	@if [ -s fmt.out ]; then \
		echo "\nERROR: The following files are not formatted:\n"; \
		cat fmt.out; \
		rm fmt.out; \
		exit 1; \
	else \
		rm fmt.out; \
	fi

generate:
	go generate $(MODULE)/...

lint:
	golangci-lint run ./...

## goimports: Apply goimports -local to the codebase
goimports:
	find . -name \*.go \
			-not -name "zz_generated.deepcopy.go" \
			-not -path "./vendor/*" \
			-not -path "./pkg/nvidia.com/resource/clientset/versioned/*" \
		-exec goimports -local $(MODULE) -w {} \;

vet:
	go vet $(MODULE)/...

COVERAGE_FILE := coverage.out
test: build
	go test -v -coverprofile=$(COVERAGE_FILE) $(MODULE)/...

coverage: test
	cat $(COVERAGE_FILE) | grep -v "_mock.go" > $(COVERAGE_FILE).no-mocks
	go tool cover -func=$(COVERAGE_FILE).no-mocks

update-pcidb:
	wget $(PCI_IDS_URL) -O $(CURDIR)/pkg/pciids/default_pci.ids

build-image: $(DOCKERFILE_DEVEL)
	$(DOCKER) build \
		--progress=plain \
		--build-arg GOLANG_VERSION="$(GOLANG_VERSION)" \
		--build-arg CLIENT_GEN_VERSION="$(CLIENT_GEN_VERSION)" \
		--build-arg CONTROLLER_GEN_VERSION="$(CONTROLLER_GEN_VERSION)" \
		--build-arg GOLANGCI_LINT_VERSION="$(GOLANGCI_LINT_VERSION)" \
		--build-arg MOQ_VERSION="$(MOQ_VERSION)" \
		--tag $(BUILDIMAGE) \
		-f $(DOCKERFILE_DEVEL) \
		.

$(DOCKER_TARGETS): docker-%:
	@echo "Running 'make $(*)' in container image $(BUILDIMAGE)"
	$(DOCKER) run \
		--rm \
		-e GOCACHE=/tmp/.cache/go \
		-e GOMODCACHE=/tmp/.cache/gomod \
		-e GOLANGCI_LINT_CACHE=/tmp/.cache/golangci-lint \
		-v $(PWD):/work \
		-w /work \
		--user $$(id -u):$$(id -g) \
		$(BUILDIMAGE) \
			make $(*)

# Start an interactive shell using the development image.
PHONY: .shell
.shell:
	$(DOCKER) run \
		--rm \
		-ti \
		-e GOCACHE=/tmp/.cache \
		-v $(PWD):$(PWD) \
		-w $(PWD) \
		--user $$(id -u):$$(id -g) \
		$(BUILDIMAGE)
