# Golang client for Linode generated with OpenAPI Generator

This is experimental. Don't use it.

This repo contains Golang code generated from [a custom version of the Linode API spec](patched.openapi.yaml) based on the [official spec](https://www.linode.com/docs/api/openapi.yaml).

Contents:

- `Makefile` includes tasks to fetch the API spec, apply patches, and generate a client
- `fetched.openapi.yaml` the latest fetched openapi spec
- `patched.openapi.yaml` the latest patched openapi spec
- `patches/*.patch` patch files to apply against the Linode API spec
- `clients/` generated clients
- `types/` generated types
- `examples/` hand crafted examples to demonstrate usage (none yet)

## Build

To build the client, run `make`.
