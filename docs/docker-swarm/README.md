# Docker Workshop - Docker Swarm

## Sections:

* [What is Docker Swarm](#what-is-docker-swarm)
* [Docker Swarm Nodes](#docker-swarm-nodes)
* [Online Docker Sandbox](#online-docker-sandbox)
* [Local Docker Swarm](#local-docker-swarm)
* [Activate Docker Swarm](#activate-docker-swarm)
* [Docker Swarm Services](#docker-swarm-services)
* [Kill the Docker Swarm](#kill-the-docker-swarm)
* [Bread Crumb Navigation](#bread-crumb-navigation)

## What is Docker Swarm
_________________________

* Docker Swarm is native clustering for Docker.
* It turns a pool of Docker hosts into a single, virtual Docker host
* It lets you handle a cluster of machines as a single docker daemon, with automatic failover,
    * container scheduling, routing and more.

## Docker Swarm Nodes
_________________________

* You need a machine cluster to get docker swarm going.
* Try online with [`play-with-docker`](http://play-with-docker.com/) to try it online.

**Prerequisites**

* Install [`docker-machine`](https://docs.docker.com/machine/) and [`Virtualbox`](https://www.virtualbox.org/).

* If you have `Docker for mac` or `Docker for windows` you probably already have it installed;
* `Linux` users should get `docker-machine` separately.

* Docker Swarm is a resource intensive so brace yourself.
* If you're just trying it out, the online route is probably what you want.
* If you'd like your swarm to be persistent and/or experiment than locally is for you

## Online Docker Sandbox
_________________________

Go to [`play-with-docker`](http://play-with-docker.com/) and create three nodes with the:

**`+ ADD NEW INSTANCE` button**

Skip to the [next section](#activate-docker-swarm).

## Local Docker Swarm
_________________________

Make sure you have `docker-machine` installed by doing:

```bash
$ docker-machine --version
## Output should appear like this
docker-machine version 0.12.0, build 45c69ad
```

Create the cluster master node using `docker-machine create`:

```bash
$ docker-machine create -d virtualbox manager
# This will take a minute or so
Running pre-create checks...
Creating machine...
(manager) Copying /Users/marcelbelmont/.docker/machine/cache/boot2docker.iso to /Users/marcelbelmont/.docker/machine/machines/m
anager/boot2docker.iso...
(manager) Creating VirtualBox VM...
(manager) Creating SSH key...
(manager) Starting the VM...
(manager) Check network to re-create if needed...
(manager) Waiting for an IP...
Waiting for machine to be running, this may take a few minutes...
Detecting operating system of created instance...
Waiting for SSH to be available...
Detecting the provisioner...
Provisioning with boot2docker...
Copying certs to the local machine directory...
Copying certs to the remote machine...
Setting Docker configuration on the remote daemon...
Checking connection to Docker...
Docker is up and running!
To see how to connect your Docker Client to the Docker Engine running on this virtual machine, run: docker-machine env manager
```

We can ssh into this like we did in the `docker-machine` lesson:

```bash
$ docker-machine ssh manager
# Output like this
docker-machine ssh manager
                        ##         .
                  ## ## ##        ==
               ## ## ## ## ##    ===
           /"""""""""""""""""\___/ ===
      ~~~ {~~ ~~~~ ~~~ ~~~~ ~~~ ~ /  ===- ~~~
           \______ o           __/
             \    \         __/
              \____\_______/
 _                 _   ____     _            _
| |__   ___   ___ | |_|___ \ __| | ___   ___| | _____ _ __
| '_ \ / _ \ / _ \| __| __) / _` |/ _ \ / __| |/ / _ \ '__|
| |_) | (_) | (_) | |_ / __/ (_| | (_) | (__|   <  __/ |
|_.__/ \___/ \___/ \__|_____\__,_|\___/ \___|_|\_\___|_|
Boot2Docker version 17.06.0-ce, build HEAD : 0672754 - Thu Jun 29 00:06:31 UTC 2017
Docker version 17.06.0-ce, build 02c1d87

docker@manager:~$ exit
```


Now let us create some worker nodes for the swarm:

```bash
$ docker-machine create -d virtualbox w1
# output here
$ docker-machine create -d virtualbox w2
# output the same here
Running pre-create checks...
Creating machine...
(w2) Copying /Users/marcelbelmont/.docker/machine/cache/boot2docker.iso to /Users/marcelbelmont/.docker/machine/machines/w2/boo
t2docker.iso...
(w2) Creating VirtualBox VM...
(w2) Creating SSH key...
(w2) Starting the VM...
(w2) Check network to re-create if needed...
(w2) Waiting for an IP...
Waiting for machine to be running, this may take a few minutes...
Detecting operating system of created instance...
Waiting for SSH to be available...
Detecting the provisioner...
Provisioning with boot2docker...
Copying certs to the local machine directory...
Copying certs to the remote machine...
Setting Docker configuration on the remote daemon...
Checking connection to Docker...
Docker is up and running!
To see how to connect your Docker Client to the Docker Engine running on this virtual machine, run: docker-machine env w2
```

We need to talk to the docker daemon in the manager node.

This time we will connect using a different command, it is actually the last line when we provision machine:

```bash
$ eval $(docker-machine env manager)
```

* This will set some environment variables telling the docker client both how to find the docker daemon
    * and where to find docker daemon so that the docker client can communicate with the docker daemon

```bash
$ env | grep DOCKER
# You should see output like this
DOCKER_TLS_VERIFY=1
DOCKER_HOST=tcp://192.168.99.100:2376
DOCKER_CERT_PATH=/Users/marcelbelmont/.docker/machine/machines/manager
DOCKER_MACHINE_NAME=manager
```

* These commands run on the manager docker daemon you will not see them using `docker images` or `docker ps`
* You ordinarily see them when you communicate with the local docker daemon
* You will need to open a new terminal to go back to the local docker daemon or `unset` the environment variables

## Activate Docker Swarm
_________________________

Copy the manager node HOST IP address:

```bash
$ env | grep DOCKER
DOCKER_TLS_VERIFY=1
DOCKER_HOST=tcp://192.168.99.100:2376
DOCKER_CERT_PATH=/Users/marcelbelmont/.docker/machine/machines/manager
DOCKER_MACHINE_NAME=manager
```

**If you are using `play-with-docker` you will see it in that node's shell prompt**

[docker swarm init docs](https://docs.docker.com/engine/reference/commandline/swarm_init/)

```bash
$ docker swarm init --advertise-addr 192.168.99.100
# Output like this should appear
Swarm initialized: current node (p4vlwh6vt9eyiuzxkkvs3tx98) is now a manager.

To add a worker to this swarm, run the following command:

docker swarm join --token SWMTKN-1-4wpg1wf5xlqpwc16iekegpeb28scexyxmy4g94zuk7kyb03v5y-bumbgd89f0hxvdujr0cl5g6pa 192.168.99.
100:2377

To add a manager to this swarm, run 'docker swarm join-token manager' and follow the instructions.
```


This set the node's docker daemon to swarm mode and output the `swarm join` command you'll need for other nodes to join this swarm. Copy it to your clipboard; you'll need it soon.

Verify the swarm status by doing.

```bash
$ docker info
# Output like should appear
...
Swarm: active
 NodeID: p4vlwh6vt9eyiuzxkkvs3tx98
 Is Manager: true
 ClusterID: u38ogdxglizq9ecvmbui091u8
 Managers: 1
 Nodes: 1
 Orchestration:
  Task History Retention Limit: 5
 Raft:
  Snapshot Interval: 10000
  Number of Old Snapshots to Retain: 0
  Heartbeat Tick: 1
  Election Tick: 3
 Dispatcher:
  Heartbeat Period: 5 seconds
 CA Configuration:
  Expiry Duration: 3 months
  Force Rotate: 0
 Root Rotation In Progress: false
 Node Address: 192.168.99.100
 Manager Addresses:
  192.168.99.100:2377
...
```

**Notice that the swarm is active and there is a manager and it gives you the manager IP address**


Run the following command to see the swarm nodes:

```bash
$ docker node ls
ID                            HOSTNAME            STATUS              AVAILABILITY        MANAGER STATUS
p4vlwh6vt9eyiuzxkkvs3tx98 *   manager             Ready               Active              Leader
```

The next step is to make both the worker nodes join the docker swarm cluster:

1. Run this commmand `docker-machine ssh w1` to ssh into worker `w1`
2. Run `docker swarm join --token SWMTKN-1-4wpg1wf5xlqpwc16iekegpeb28scexyxmy4g94zuk7kyb03v5y-bumbgd89f0hxvdujr0cl5g6pa 192.168.99.
100:2377` to make w1 join the swarm.
3. If the command worked you will see this output `This node joined a swarm as a worker.`
4. Run `exit` to leave this worker
5. Run this commmand `docker-machine ssh w2` to ssh into worker `w2`
6. Run `docker swarm join --token SWMTKN-1-4wpg1wf5xlqpwc16iekegpeb28scexyxmy4g94zuk7kyb03v5y-bumbgd89f0hxvdujr0cl5g6pa 192.168.99.
100:2377` to make w2 join the swarm
7. If the command worked you will see this output `This node joined a swarm as a worker.`
8. Run `exit` to leave this worker

Let us verify that the docker swarm has 3 nodes now:

```bash
$ docker node ls
ID                            HOSTNAME            STATUS              AVAILABILITY        MANAGER STATUS
p4vlwh6vt9eyiuzxkkvs3tx98 *   manager             Ready               Active              Leader
svj88yvnpq25mrynlecv4z99s     w2                  Ready               Active
ul4dhx5i23v3dcsj77cvvavph     w1                  Ready               Active
```

Notice that before we only had 1 node but now we have 3 nodes since the 2 workers joined

## Docker Swarm Services
_________________________

We will create a service that pings `google.com`

```bash
$ docker service create --name ping_google --replicas 1 alpine ping google.com
# this will output an ID
v483vd7v286n6wic0rnd37dri
```

Now to print information about this service we just created run this command:

```bash
$ docker service inspect --pretty ping_google
# Output like this should appear
ID:             v483vd7v286n6wic0rnd37dri
Name:           ping_google
Service Mode:   Replicated
 Replicas:      1
Placement:
UpdateConfig:
 Parallelism:   1
 On failure:    pause
 Monitoring Period: 5s
 Max failure ratio: 0
 Update order:      stop-first
RollbackConfig:
 Parallelism:   1
 On failure:    pause
 Monitoring Period: 5s
 Max failure ratio: 0
 Rollback order:    stop-first
ContainerSpec:
 Image:         alpine:latest@sha256:1072e499f3f655a032e88542330cf75b02e7bdf673278f701d7ba61629ee3ebe
 Args:          ping google.com
Resources:
Endpoint Mode:  vip
```

To check service status run this command:

```bash
$ docker service ps ping_google
# Output like this should appear
ID                  NAME                IMAGE               NODE                DESIRED STATE       CURRENT STATE           ERROR               PORTS
t0teszvht02q        ping_google.1       alpine:latest       w2                  Running             Running 2 minutes ago
```

Now let us scale the service and get more replicas by getting this command:

```bash
$ docker service scale ping_google=5
# Output like this
ping_google scaled to 5
```

Now to see which nodes the new replicas are running in do this command:

```bash
$ docker service ps ping_google
ID                  NAME                IMAGE               NODE                DESIRED STATE       CURRENT STATE
  ERROR               PORTS
t0teszvht02q        ping_google.1       alpine:latest       w2                  Running             Running 6 minutes ago

i5ksu8mecbtd        ping_google.2       alpine:latest       manager             Running             Running about a minute ago

ib1cqm3l8j9v        ping_google.3       alpine:latest       w2                  Running             Running about a minute ago

u309fwgz7trk        ping_google.4       alpine:latest       w1                  Running             Running about a minute ago

qzt180qbuw1a        ping_google.5       alpine:latest       manager             Running             Running about a minute ago
```

**Notice that since we scaled this to 5, we get 5 Nodes now**

**This is now a local Docker Swarm**

## Kill the Docker Swarm
_________________________

If you are using `play-with-docker` you can just click the delete button in any worker node.

```bash
$ docker-machine rm w1
$ docker service ps ping_google
ID                  NAME                IMAGE               NODE                DESIRED STATE       CURRENT STATE            ER
ROR               PORTS
t0teszvht02q        ping_google.1       alpine:latest       w2                  Running             Running 10 minutes ago

i5ksu8mecbtd        ping_google.2       alpine:latest       manager             Running             Running 5 minutes ago

ib1cqm3l8j9v        ping_google.3       alpine:latest       w2                  Running             Running 5 minutes ago

2k7470xhl9l3        ping_google.4       alpine:latest       manager             Running             Running 34 seconds ago

u309fwgz7trk         \_ ping_google.4   alpine:latest       w1                  Shutdown            Running 5 minutes ago

qzt180qbuw1a        ping_google.5       alpine:latest       manager             Running             Running 5 minutes ago
```

**Notice that node `w1` was shut down after we removed the `w1` worker and the service was rescheduled**

You can run `docker-machine rm w2` and maybe `docker-machine rm manager`

If you get errors after running this then do the following:

1. `docker-machine regenerate-certs default`
2. `docker-machine create -d virtualbox default`
3. `eval $(docker-machine env default)`

Now the environment should be reset back to default

## Bread Crumb Navigation
_________________________

Previous | Next
:------- | ---:
← [Docker Machine](../docker-machine/README.md) | [Dockerhub](../dockerhub/README.md) →
