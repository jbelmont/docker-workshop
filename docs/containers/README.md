# Docker Workshop - Containers

[Containers](https://docs.docker.com/get-started/#prerequisites)

Content:

* [Definition of a Container](#definition-of-a-container)
* [Containers vs Virtual Machines](#containers-vs-virtual-machines)

## Definition of a Container

### A brief explanation of containers:

* An image is a lightweight, stand-alone, executable package that includes everything
    * needed to run a piece of software, including the code, a runtime, libraries, environment variables
    * and config files.

* A container is a runtime instance of an image—what the image becomes in memory when actually executed.
* It runs completely isolated from the host environment by default
    * only accessing host files and ports if configured to do so.

* Containers run apps natively on the host machine’s kernel.
* They have better performance characteristics than virtual machines that only get virtual access to host
    * resources through a hypervisor.
* Containers can get native access, each one running in a discrete process, taking no more memory than
    * any other executable.

* Using containers, everything required to make a piece of software run is packaged into isolated containers.
* Unlike VMs, containers do not bundle a full operating system - only libraries and settings required to make
    * the software work are needed.
* This makes for efficient, lightweight, self-contained systems and guarantees that software will always run
    * the same, regardless of where it’s deployed.

## Containers vs Virtual Machines

**Consider this diagram comparing virtual machines to containers:**

### Virtual Machine Diagram

![Virtual Machines](../../images/virtualmachine.png)

* Virtual machines run guest operating systems—note the OS layer in each box.
* This is resource intensive, and the resulting disk image and application state is an entanglement of OS
    * settings, system-installed dependencies, OS security patches, and other easy-to-lose,
    * hard-to-replicate ephemera.

### Container Diagram

![Container Diagram](../../images/container.png)

* Containers can share a single kernel, and the only information that needs to be in a container image is
    * the executable and its package dependencies, which never need to be installed on the host system.
* These processes run like native processes, and you can manage them individually by running commands like
    * docker ps—just like you would run ps on Linux to see active processes.
* Finally, because they contain all their dependencies, there is no configuration entanglement;
    * a containerized app “runs anywhere.”

Previous | Next
:------- | ---:
← [README](../../README.md) | [Docker Basics](../docker-basics/README.md) →
