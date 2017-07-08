# Docker Workshop - Dockerfile

# Sections:

* [What is a Dockerfile](#what-is-a-dockerfile)
* [Usage of docker build](#usage-of-docker-build)
* [Understand how CMD and ENTRYPOINT interact](#understand-how-cmd-and-entrypoint-interact)

## What is a Dockerfile

* Docker can build images automatically by reading the instructions from a `Dockerfile`.
* A `Dockerfile` is a text document that contains all the commands a user could call on the command line to assemble an image.
* Using `docker build` users can create an automated build that executes several command-line instructions in succession.

## Usage of docker build

* The `docker build` command builds an image from a `Dockerfile` and a `context`.
* The build’s context is the files at a specified location `PATH` or `URL`
    * The `PATH` is a directory on your local filesystem.
    * The `URL` is a Git repository location.
* A context is processed recursively
    * So, a `PATH` includes any subdirectories and the `URL` includes the repository and its submodules.

Remember in the previous document `docker basics` we did the following:

```bash
$ docker build -t hello-world .
Sending build context to Docker daemon  5.632kB
Step 1/2 : FROM nginx:1.8-alpine
 ---> c0dddb65129b
Step 2/2 : ADD site /usr/share/nginx/html
 ---> Using cache
 ---> 37e37aa0cac0
Successfully built 37e37aa0cac0
```

Dockerfile Reference Table:

| Command | Reference | Description |
| --- | --- | --- |
| ADD | [Doc](https://docs.docker.com/engine/reference/builder/#add) | Copies files, directories or remote file URLs from `src` and adds them container |
| CMD | [Doc](https://docs.docker.com/engine/reference/builder/#copy) | Provide defaults for an executing container |
| ENTRYPOINT | [Doc](https://docs.docker.com/engine/reference/builder/#entrypoint) | Allows you to configure a container that will run as an executable |
| ENV | [Doc](https://docs.docker.com/engine/reference/builder/#env) | Instruction sets the environment variable `key` to the value `value` |
| EXPOSE | [Doc](https://docs.docker.com/engine/reference/builder/#expose) | Instruction informs Docker that the container listens on the specified network ports at runtime |
| COPY | [Doc](https://docs.docker.com/engine/reference/builder/#copy) | Instruction copies new files or directories from `src` and adds them to the filesystem of the container at the path `dest` |
| LABEL | [Doc](https://docs.docker.com/engine/reference/builder/#label) | Instruction adds metadata to an image. A `LABEL` is a key-value pair |
| VOLUME | [Doc](https://docs.docker.com/engine/reference/builder/#volume) | Creates a shared volume that can be shared among containers or by the host machine |
| WORKDIR | [Doc](https://docs.docker.com/engine/reference/builder/#workdir) | Set the default working directory for the container |
| FROM | [Doc](https://docs.docker.com/engine/reference/builder/#from) | The base image to use in the build. This is mandatory and must be the first command in the file. |
| USER | [Doc](https://docs.docker.com/engine/reference/builder/#user) | Sets the default user within the container |
| ARG | [Doc](https://docs.docker.com/engine/reference/builder/#arg) | Instruction defines a variable that users can pass at build-time to the builder with the docker build command using the `--build-arg` `varname=value` flag |
| ONBUILD | [Doc](https://docs.docker.com/engine/reference/builder/#onbuild) | A command that is triggered when the image in the `Dockerfile` is used as a base for another image |
| STOPSIGNAL | [Doc](https://docs.docker.com/engine/reference/builder/#stopsignal) | Instruction sets the system call signal that will be sent to the container to exit |
| HEALTHCHECK | [Doc](https://docs.docker.com/engine/reference/builder/#healthcheck) | Instruction tells Docker how to test a container to check that it is still working |
| SHELL | [Doc](https://docs.docker.com/engine/reference/builder/#shell) | Instruction allows the default `shell` used for the shell form of commands to be overridden |
| MAINTAINER | [Doc](https://docs.docker.com/engine/reference/builder/#maintainer-deprecated) | Instruction sets the Author field of the generated images. |

[Dockerfile Project](http://dockerfile.github.io/)

Go Dockerfile:

```dockerfile
# Pull base image.
FROM dockerfile/go

LABEL maintainer="marcelbelmont@gmail.com"

# This is deprecated so you should use the LABEL command instead
MAINTAINER Marcel Belmont "marcelbelmont@gmail.com"

ENV GOVERSION 1.8.3

# Add scripts.
ADD bin/go-build /usr/local/bin/go-build
ADD bin/go-run /usr/local/bin/go-run

# Add executable permission to scripts.
RUN chmod +x /usr/local/bin/go-*

# Set instructions on build.
ONBUILD ADD . /gopath/src/app/
ONBUILD RUN go-build

# Define default command.
CMD ["go-run"]

# Expose port.
EXPOSE 8080
```

Node.js Dockerfile:

```dockerfile
#
# Node.js Dockerfile
#
# https://github.com/dockerfile/nodejs
#

# Pull base image.
FROM dockerfile/python

# Install Node.js
RUN \
  cd /tmp && \
  wget http://nodejs.org/dist/node-latest.tar.gz && \
  tar xvzf node-latest.tar.gz && \
  rm -f node-latest.tar.gz && \
  cd node-v* && \
  ./configure && \
  CXX="g++ -Wno-unused-local-typedefs" make && \
  CXX="g++ -Wno-unused-local-typedefs" make install && \
  cd /tmp && \
  rm -rf /tmp/node-v* && \
  npm install -g npm && \
  printf '\n# Node.js\nexport PATH="node_modules/.bin:$PATH"' >> /root/.bashrc

# Define working directory.
WORKDIR /data

# Define default command.
CMD ["bash"]
```

## Understand how CMD and ENTRYPOINT interact

[Understand how CMD and ENTRYPOINT interact](https://docs.docker.com/engine/reference/builder/#understand-how-cmd-and-entrypoint-interact)

* Both `CMD` and `ENTRYPOINT` instructions define what command gets executed when running a container.
* There are few rules that describe their co-operation.
* Dockerfile should specify at least one of `CMD` or `ENTRYPOINT` commands.
* ENTRYPOINT should be defined when using the container as an executable.
* `CMD` should be used as a way of defining default arguments for an ENTRYPOINT command or for executing an ad-hoc command in a container.
* `CMD` will be overridden when running the container with alternative arguments.

EntryPoint Dockerfile:

```dockerfile
FROM ubuntu:16.04

COPY debug-mode.sh /root

RUN \
uptime

ENTRYPOINT ["/usr/bin/uptime"]
CMD ["-V"]
```

ARG Dockerfile:

```dockerfile
# ARG is the only command that can precede FROM which is ordinarily first COMMAND in Dockerfile
ARG  CODE_VERSION=latest
FROM base:${CODE_VERSION}
CMD  /code/run-app

FROM extras:${CODE_VERSION}
CMD  /code/run-extras
```

RethinkDB dockerfile:

```dockerfile
#
# RethinkDB Dockerfile
#
# https://github.com/dockerfile/rethinkdb
#

# Pull base image.
FROM dockerfile/ubuntu

# Install RethinkDB.
RUN \
  echo "deb http://download.rethinkdb.com/apt `lsb_release -cs` main" > /etc/apt/sources.list.d/rethinkdb.list && \
  wget -O- http://download.rethinkdb.com/apt/pubkey.gpg | apt-key add - && \
  apt-get update && \
  apt-get install -y rethinkdb python-pip && \
  rm -rf /var/lib/apt/lists/*

# Install python driver for rethinkdb
RUN pip install rethinkdb

# Define mountable directories.
VOLUME ["/data"]

# Define working directory.
WORKDIR /data

# Define default command.
CMD ["rethinkdb", "--bind", "all"]

# Expose ports.
#   - 8080: web UI
#   - 28015: process
#   - 29015: cluster
EXPOSE 8080
EXPOSE 28015
EXPOSE 29015
```

USER dockerfile:

```dockerfile
FROM oracle

USER root

RUN echo 'export ORACLE_SID=XE' >> /etc/bash.bashrc

USER oracle

RUN sqplus @my_schema.sql
```

STOPSIGNAL dockerfile:

```dockerfile
FROM busybox
CMD sleep 3600
STOPSIGNAL SIGKILL
```

HEALTHCHECK dockerfile:

```dockerfile
FROM alpine:3.3

MAINTAINER tomwillfixit

RUN apk update && apk add curl && rm -rf /var/cache/apk/*

COPY helloworld.bin /

EXPOSE 80

HEALTHCHECK --interval=5s --timeout=3s --retries=3 \
      CMD curl -f http://localhost:80 || exit 1

ENTRYPOINT ["/helloworld.bin"]
```

SHELL dockerfile:

```dockerfile
FROM microsoft/windowsservercore

# Executed as cmd /S /C echo default
RUN echo default

# Executed as cmd /S /C powershell -command Write-Host default
RUN powershell -command Write-Host default

# Executed as powershell -command Write-Host hello
SHELL ["powershell", "-command"]
RUN Write-Host hello

# Executed as cmd /S /C echo hello
SHELL ["cmd", "/S"", "/C"]
RUN echo hello
```

Previous | Next
:------- | ---:
← [Docker Basics](../docker-basics/README.md) | [Volumes](../volumes/README.md) →
