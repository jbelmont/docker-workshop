# Docker Workshop

The Workshop is separated in nine sections:

* [Container Philosophy](docs/containers/README.md)
* [Docker Basics](docs/docker-basics/README.md)
* [Dockerfile Usage](docs/dockerfile/README.md)
* [Volumes](docs/volumes/README.md)
* [Docker Compose](docs/docker-compose/README.md)
* [Docker Network](docs/docker-network/README.md)
* [Docker Machine](docs/docker-machine/README.md)
* [Docker Swarm](docs/docker-swarm/README.md)
* [DockerHub](docs/dockerhub/README.md)
* [Docker Best Practices](docs/docker-best-practices/README.md)

### Preparations:

* [Install Docker](https://docs.docker.com/engine/installation/)
* Clone this repo: `git clone https://github.com/jbelmont/docker-workshop.git`
* Docker Images:

```
docker pull mongo:3.4.5
docker pull golang:1.8.3
docker pull redis:3.2.9-alpine
docker pull mhart/alpine-node:8.0.0
```

### Assumptions:

* During this workshop the following ports will be used: `8081`, `8080`, `3000-3005`.

If they are not available on your machine, adjust the CLI commands accordingly.

### Group Details

Join [Code Craftsmanship Saturdays Meetup](https://www.meetup.com/Code-Craftsmanship-Saturdays/events/240767853/)

Join [Code Craftsmanship Saturdays Github Org](https://github.com/Code-Craftsmanship-Saturdays)

Prerequisites
-------------

Since Docker leverages the Operating System's virtualization technologies, the install requirements for Docker are specific.

OS X requirements:

- 2010 or newer model with Intel's MMU virtualization
- OS X El Capitan 10.11 or newer

Windows requirements:

- 64-bit Windows
- Windows 10 Pro, Enterprise or Education (not Home, not Windows 7 or 8) to install Hyper-V
- Windows 10 Anniversary Update or better
- Access to your machine's BIOS to turn on virtualization

[Install Docker](https://docs.docker.com/engine/installation/)

![Docker Logo](images/docker.png)

## Gitbook

Go to the gitbook [Here](https://jbelmont.github.io/docker-workshop/)
