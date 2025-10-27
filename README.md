# simple pandoc server

## building

*for the latest building guide see the [Wiki](https://github.com/dozro/simple-pandoc-server/wiki/building)*

### build using go-task

Simply clone the Repo (newer then the merge of https://github.com/dozro/simple-pandoc-server/pull/4) and make sure you have Task installed.  
An installation guide can be found here: [Installation guide for go-task](https://taskfile.dev/docs/installation).   

The Taskfile, where the build configuration is defined, in the [Taskfile.yml](./Taskfile.yml).

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
