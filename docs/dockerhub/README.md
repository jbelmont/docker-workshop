# Docker Workshop - Docker Hub

Section:

* [Docker Hub](#docker-hub)
* [Using Docker Hub](#using-docker-hub)
* [Adding images to Docker Hub](#adding-images-to-docker-hub)

## Docker Hub
==========

Think of Docker Hub being to docker what Github is to source control systems for `git`

Docker Hub is the central repository or "store" for sharing images.

Private Repositories for Docker Hub are also available to store images:

* [AWS Container Registry](https://aws.amazon.com/ecr/)
* [Azure Container Registry](https://azure.microsoft.com/en-us/services/container-registry/)
* [Artifactory](https://www.jfrog.com/artifactory/) (paid) When inside a firewall
* [Nexus](https://www.sonatype.com/nexus-repository-oss) (free) for on-premise Docker repositories.

## Using Docker Hub
----------------

Throughout the workshop we used `docker pull SOME_IMAGE` and each time we did this we pulled content from docker hub.

* Browse [https://hub.docker.com/](https://hub.docker.com/) to find other images to download.
* Browse [Docker Store](https://hub.docker.com/explore/) as the new image explorer
* Each image usually has a Dockerfile showing how that image was built and instructions on using the image.
    * Environment variables to set and more

## Adding images to Docker Hub
--------------------

(You can use AWS, Azure, or other private registries if you want to keep private images)

1. Create an account on [https://hub.docker.com/](https://hub.docker.com/).

2. Run `docker login` in order to login into the hub and push images.

3. Note the profile picture section has your username as well as `My Profile`, `Logout` and more sections.

4. Tag an image in the form `username/imagename:version`
    1. Notice I tag my images with `jbelmont/someimage`:
    2. Remember we tagged an image like this `docker tag hello-world jbelmont/hello-world`
    2. Note the registry details are in the image name.
    3. This makes it more difficult to move images between repositories or to build automation though

5. `docker push username/imagename:version` substituting the details of the image you tagged above.
    1. `docker push jbelmont/hello-world` we pushed our newly tagged image with this command

Previous | Next
:------- | ---:
← [Docker Swarm](../docker-swarm/README.md) | [Docker Compose](../docker-compose/README.md) →
