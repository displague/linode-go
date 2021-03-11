.PHONY: all gen patch fetch

SPEC_URL:=https://www.linode.com/docs/api/openapi.yaml
SPEC_FETCHED_FILE:=fetched.openapi.yaml
SPEC_PATCHED_FILE:=patched.openapi.yaml
IMAGE=openapitools/openapi-generator-cli
GIT_ORG=displague
GIT_REPO=linode-go
PACKAGE_MAJOR=v1

SWAGGER=docker run --rm -v $(CURDIR):/local ${IMAGE}

all: pull fetch patch gen

pull:
	docker pull ${IMAGE}

fetch:
	curl -o ${SPEC_FETCHED_FILE} ${SPEC_URL}

patch:
	# patch is idempotent, always starting with the fetched
	# fetched file to create the patched file.
	ARGS="-o ${SPEC_PATCHED_FILE} ${SPEC_FETCHED_FILE}"; \
	for diff in $(shell find patches/*.patch | sort -n); do \
		patch --no-backup-if-mismatch -N -t $$ARGS $$diff; \
		ARGS=${SPEC_PATCHED_FILE}; \
	done

gen:
	${SWAGGER} generate -g go \
		--package-name ${PACKAGE_MAJOR} \
		--model-package types \
		--api-package models \
		--git-user-id ${GIT_ORG} \
		--git-repo-id ${GIT_REPO} \
		-o /local/${PACKAGE_MAJOR} \
		-i /local/${SPEC_PATCHED_FILE}
