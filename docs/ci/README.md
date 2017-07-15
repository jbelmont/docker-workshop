# Docker Workshop - Continuous Integration

## Sections:

* [Using Docker in Travis CI](#using-docker-in-travis-ci)
* [Docker Compose in Travis CI](#docker-compose-in-travis-ci)
* [Bread Crumb Navigation](#bread-crumb-navigation)

## Using Docker in Travis CI

Read [Travis CI docs](https://docs.travis-ci.com/user/docker/)

```dockerfile
sudo: required

services:
  - docker
```

This will use the `Ubuntu` Trusty Image which is currently in Beta in Travis CI

Here is a sample `.travis.yml` script for Golang

[Go Travis CI environment docs](https://docs.travis-ci.com/user/languages/go/)

```yml
sudo: required

language: go

go:
  - 1.8.x

env:
    - IMAGE_NAME="golang:thing"
    - GO_VERSION="1.8.3"

services:
    - docker

branches:
  only:
    - master


before_install:
  - go get github.com/mattn/goveralls
  - go get github.com/stretchr/testify/assert

before_script:
  - git config --global user.email "marcelbelmont@gmail.com"
  - git config --global user.name "jbelmont"

script:
    - docker build -t "${IMAGE_NAME}" -q someFolder/
    - docker run -d -P --name myGoProgram -v /var/www/app go:1.8.3 go run main.go

after_success:
  - if [ "$TRAVIS_BRANCH" == "master" ]; then
    docker login -u="$DOCKER_USERNAME" -p="$DOCKER_PASSWORD";
    docker push USER/REPO;
    fi
```

**This could be separated out into a series of shell scripts as well which you can invoke in the scripts**

## Docker Compose in Travis CI

* The Docker Compose tool is also installed in the Docker enabled environment.
* So you can run docker-compose commands in the scripts block as well

## Bread Crumb Navigation
_________________________

Previous | Next
:------- | ---:
← [Dockerhub](../dockerhub/README.md) | [Docker Best Practices](../docker-best-practices/README.md) →
