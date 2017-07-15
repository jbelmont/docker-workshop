# Docker Workshop - Docker Best Practices

## Sections:
* [Best practices for writing Dockerfiles](#best-practices-for-writing-dockerfiles)
* [Read Best Practices README](#read-best-practices-readme)
* [Bread Crumb Navigation](#bread-crumb-navigation)

## Best practices for writing Dockerfiles

[Best Practices for Dockerfiles Documentation](https://docs.docker.com/engine/userguide/eng-image/dockerfile_best-practices/)

#### Containers should be ephemeral

* The container produced by the image your Dockerfile defines should be as ephemeral as possible.
* By “ephemeral,” we mean that it can be stopped and destroyed and a new one built and put in place with an absolute minimum of set-up and configuration.

####  Use a .dockerignore file

* In most cases, it’s best to put each Dockerfile in an empty directory.
* Then, add to that directory only the files needed for building the Dockerfile.
* To increase the build’s performance, you can exclude files and directories by adding a .dockerignore file to that directory as well.

#### Avoid installing unnecessary packages

* In order to reduce complexity, dependencies, file sizes, and build times, you should avoid installing extra or unnecessary packages just because they might be “nice to have.”

#### Each container should have only one concern

* Decoupling applications into multiple containers makes it much easier to scale horizontally and reuse containers

#### Minimize the number of layers

* You need to find the balance between readability (and thus long-term maintainability) of the Dockerfile and minimizing the number of layers it uses.
* Be strategic and cautious about the number of layers you use.

#### Sort multi-line arguments

* Whenever possible, ease later changes by sorting multi-line arguments alphanumerically.
* This will help you avoid duplication of packages and make the list much easier to update.
* This also makes PRs a lot easier to read and review

```dockerfile
RUN apt-get update && apt-get install -y \
  bzr \
  cvs \
  git \
  mercurial \
  subversion
```

**Try to run all the `RUN` COMMANDS in one line because doing multiple RUN commands makes image larger**

## Read Best Practices README

Github user FuriKuri wrote an excellent README compiling best practices from blogs and other articles [HERE](https://github.com/FuriKuri/docker-best-practices)

## Bread Crumb Navigation
_________________________

Previous | Next
:------- | ---:
← [Continuous Integration](../ci/README.md) | [README](../../README.md) →
