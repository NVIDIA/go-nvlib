# Copyright (c) NVIDIA CORPORATION.  All rights reserved.
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

MODULE := github.com/NVIDIA/go-nvlib

GIT_COMMIT ?= $(shell git describe --match="" --dirty --long --always --abbrev=40 2> /dev/null || echo "")
GIT_TAG ?= $(patsubst v%,%,$(shell git describe --tags 2>/dev/null))
PARTS := $(subst -, ,$(GIT_TAG))
VERSION ?= $(word 1,$(PARTS))

# vVERSION represents the version with a guaranteed v-prefix
vVERSION := v$(VERSION:v%=%)

GOLANG_VERSION ?= 1.23.5

ifeq ($(IMAGE),)
REGISTRY ?= nvidia
IMAGE=$(REGISTRY)/go-nvlib
endif
IMAGE_TAG ?= $(GOLANG_VERSION)

# these variables are only needed when building a local image
# by default, the k8s-test-infra image is used
CLIENT_GEN_VERSION ?= v0.26.1
CONTROLLER_GEN_VERSION ?= v0.9.2
GOLANGCI_LINT_VERSION ?= v1.52.0
MOQ_VERSION ?= v0.3.4

BUILDIMAGE ?= ghcr.io/nvidia/k8s-test-infra:devel-go$(GOLANG_VERSION)
DOCKERFILE_DEVEL := "images/devel/Dockerfile"
K8S_TEST_INFRA := "https://github.com/NVIDIA/k8s-test-infra.git"
