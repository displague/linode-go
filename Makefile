.PHONY: all gen patch fetch

SPEC_URL:=https://www.linode.com/docs/api/openapi.yaml
SPEC_FETCHED_FILE:=fetched.openapi.yaml
SPEC_PATCHED_FILE:=patched.openapi.yaml
IMAGE=openapitools/openapi-generator-cli

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
		--model-package types \
		-o /local/v1 \
		-i /local/${SPEC_PATCHED_FILE}