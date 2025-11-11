# simple pandoc server

[![Dependabot Updates](https://github.com/dozro/simple-pandoc-server/actions/workflows/dependabot/dependabot-updates/badge.svg)](https://github.com/dozro/simple-pandoc-server/actions/workflows/dependabot/dependabot-updates)
[![Docker](https://github.com/dozro/simple-pandoc-server/actions/workflows/docker-publish.yml/badge.svg)](https://github.com/dozro/simple-pandoc-server/actions/workflows/docker-publish.yml)
[![go-lint-test-vet.yml](https://github.com/dozro/simple-pandoc-server/actions/workflows/go-lint-test-vet.yml/badge.svg)](https://github.com/dozro/simple-pandoc-server/actions/workflows/go-lint-test-vet.yml)


simple-pandoc-server is a lightweight HTTP server wrapper for Pandoc that allows document conversions via API-calls.  

## Motivation

an easy way to convert files saved in latex for example to html for presentation in Gitea.
A guide on how to do this using this service can be found over in the [Wiki](https://github.com/dozro/simple-pandoc-server/wiki/Use-as-Renderer-for-Gitea).

## building

*for the latest building guide see the [Wiki](https://github.com/dozro/simple-pandoc-server/wiki/building)*

### build using go-task

#### Requirements

- having [Go Task](https://taskfile.dev/docs/installation) installed with version 3 or newer
- having [Go](https://go.dev/dl/) installed with version 1.24 or newer
- having [pandoc](https://pandoc.org/installing.html) installed

#### Guide

Clone the repository at or after [PR #4](https://github.com/dozro/simple-pandoc-server/pull/4) was merged (for Task support) and make sure you have Task installed.  

The Taskfile is defined in the [Taskfile.yml](./Taskfile.yml).

Then run

```sh
task
```

It will build the application to `out/simple-pandoc-server`. Simply run this executable and there you go.

### build using docker

If you have Task installed (see the point above) then you can run

```sh
# to build the docker image (it will be named sps)
# TODO: make the naming dependent on env variables
task docker-build
# to build and run the docker image
task docker-run
```

If you don't have Task installed you can build the docker container by yourself using the [Dockerfile](./Dockerfile).

```sh
docker build -t sps .
```
