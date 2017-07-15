# Docker Workshop - Docker Compose

## Sections:

* [Install Docker Compose](#install-docker-compose)
* [What is Docker Compose](#what-is-docker-compose)
* [Docker Compose Commands](#docker-compose-commands)
* [docker-compose.yml](#docker\-compose\.yml)
* [Docker Compose build files](#docker-compose-build-files)
* [Run the Golang API](#run-the-golang-api)
* [Bread Crumb Navigation](#bread-crumb-navigation)

## Install Docker Compose
_________________________

Download and Install [Docker Compose](https://docs.docker.com/compose/install/)

Check that docker-compose is installed with this command:

```bash
$ docker-compose version
# If docker-compose is installed you should see this output
docker-compose version 1.14.0, build c7bdf9e
docker-py version: 2.3.0
CPython version: 2.7.12
OpenSSL version: OpenSSL 1.0.2j  26 Sep 2016
```

## What is Docker Compose
_________________________

* Compose is a tool for defining and running multi-container Docker applications.
* With Compose, you use a Compose file to configure your application's services.
* Then, using a single command, you create and start all the services from your configuration.
* To learn more about all the features of Compose see [the list of features](https://docs.docker.com/compose/overview/#features)

Using Compose is basically a three-step process:

1. Define your app’s environment with a `Dockerfile` so it can be reproduced anywhere.
2. Define the services that make up your app in `docker-compose.yml` so they can be run together in an isolated environment.
3. Lastly, run `docker-compose up` and Compose will start and run your entire app.

## Docker Compose Commands
_________________________

Check the available commands of Docker Compose. Type in your terminal:

<details>

<pre>
$ docker-compose
# output like this
...
Usage:
  docker-compose [-f <arg>...] [options] [COMMAND] [ARGS...]
  docker-compose -h|--help

Options:
  -f, --file FILE             Specify an alternate compose file (default: docker-compose.yml)
  -p, --project-name NAME     Specify an alternate project name (default: directory name)
  --verbose                   Show more output
  -v, --version               Print version and exit
  -H, --host HOST             Daemon socket to connect to

  --tls                       Use TLS; implied by --tlsverify
  --tlscacert CA_PATH         Trust certs signed only by this CA
  --tlscert CLIENT_CERT_PATH  Path to TLS certificate file
  --tlskey TLS_KEY_PATH       Path to TLS key file
  --tlsverify                 Use TLS and verify the remote
  --skip-hostname-check       Don't check the daemon's hostname against the name specified
                              in the client certificate (for example if your docker host
                              is an IP address)
  --project-directory PATH    Specify an alternate working directory
                              (default: the path of the Compose file)
...
</pre>

</details>

* Whenever you don't remember a command, just type docker-compose
* For more information on a particular command type `docker-compose help COMMAND`
    * (e.g. `docker-compose help start`)

| Command | Reference | Description |
| --- | --- | --- |
| build | [Documentation](https://docs.docker.com/compose/reference/build/) | Build or rebuild services |
| bundle | [Documentation](https://docs.docker.com/compose/reference/bundle/) | Generate a Docker bundle from the Compose file |
| config | [Documentation](https://docs.docker.com/compose/reference/config/) | Validate and view the Compose file |
| create | [Documentation](https://docs.docker.com/compose/reference/create/) | Create services |
| down | [Documentation](https://docs.docker.com/compose/reference/down/) | Stop and remove containers, networks, images, and volumes |
| events | [Documentation](https://docs.docker.com/compose/reference/events/) | Receive real time events from containers |
| exec | [Documentation](https://docs.docker.com/compose/reference/exec/) | Execute a command in a running container |
| help | [Documentation](https://docs.docker.com/compose/reference/help/) | Get help on a command |
| images | No Documention | List images |
| kill | [Documentation](https://docs.docker.com/compose/reference/kill/) | Kill containers |
| logs | [Documentation](https://docs.docker.com/compose/reference/logs/) | View output from containers |
| pause | [Documentation](https://docs.docker.com/compose/reference/pause/) | Pause services |
| port | [Documentation](https://docs.docker.com/compose/reference/port/) | Print the public port for a port binding |
| ps | [Documentation](https://docs.docker.com/compose/reference/ps/) | List containers |
| pull | [Documentation](https://docs.docker.com/compose/reference/pull/) | Pull service images |
| push | [Documentation](https://docs.docker.com/compose/reference/push/) | Push service images |
| restart | [Documentation](https://docs.docker.com/compose/reference/restart/) | Restart services |
| rm | [Documentation](https://docs.docker.com/compose/reference/rm/) | Create services |
| run | [Documentation](https://docs.docker.com/compose/reference/run/) | Run a one-off command |
| scale | [Documentation](https://docs.docker.com/compose/reference/scale/) | Set number of containers for a service |
| start | [Documentation](https://docs.docker.com/compose/reference/start/) | Start services |
| stop | [Documentation](https://docs.docker.com/compose/reference/stop/) | Stop services |
| top | [Documentation](https://docs.docker.com/compose/reference/top/) | Display the running processes |
| unpause | [Documentation](https://docs.docker.com/compose/reference/unpause/) | Unpause services |
| up | [Documentation](https://docs.docker.com/compose/reference/up/) | Create and start containers |
| version | [Documentation](https://docs.docker.com/compose/reference/version/) | Show the Docker-Compose version information |

## docker-compose.yml
_________________________

The `docker-compose.yml` file is a [YAML](http://yaml.org/) file defining [services](https://docs.docker.com/compose/compose-file/#service-configuration-reference), [networks](https://docs.docker.com/compose/compose-file/#network-configuration-reference) and [volumes](https://docs.docker.com/compose/compose-file/#volume-configuration-reference). The default path for a Compose file is `./docker-compose.yml`.

<details>

<h4>Key-value pairs (Scalars)</h4>

<ul>
    <li>YAML keeps data stored as a map containing keys and values associated to those keys.</li>
    <li>This map is in no particular order, so you can reorder it at will. Each pair is in the format KEY: VALUE. </li>
</ul>

 For example:

<pre>
name: 'John J'
age: 32
</pre>

<ul>
    <li>Note the 'quotes' around the value. When the value is a text string</li>
    <li>The quotes are used to make sure any special characters are not given special meaning</li>
    <li>Instead all the values in quote are the value.</li>
    <li>So even though they are optional, using them is highly recommended.</li>
</ul>

YAML will consider that lines prefixed with more spaces than the parent key are contained inside it;
Moreover, all lines must be prefixed with the same amount of spaces to belong to the same map. So this works:

<pre>
prop:
    subprop:
        value: '(%person%) %name%'
        value2: '* %fruit% %rank%'
</pre>

Alternative:

<pre>
prop:
        subprop:
                    value: '(%person%) %name%'
                    value2: '* %fruit% %rank%'
</pre>

This example won't work:

<pre>
formatting:
from-game:
chat: '(%sender%) %message%'
action: '* %sender% %message%'
</pre>

<h4>Alternative YAML format</h4>

YAML supports an alternative syntax to store key-value maps, useful for compressing small maps into a single line.

The syntax is: {KEY: VALUE, KEY: VALUE, ...}. The above example would become:

<code>
formatting: {from-game: {chat: '(%sender%) %message%', action: '* %sender% %message%'}}
</code>

<h4>Lists</h4>

<ul>
    <li>Lists are used to store a collection of ordered values.</li>
    <li>The values are not associated with a key</li>
    <li>only with a positional index obtained from the order in which they are specified (item 1, item 2, etc.)</li>
</ul>

<h4>Block Sequences</h4>

<pre>
- 'item 1'
- 'item 2'
</pre>

<h4>Inline Sequences</h4>

<code>
items: ['item 1', 'item 2']
</code>

You can have:
* Maps inside maps
* Lists inside maps
* Maps inside lists

<h4>Anchors</h4>

<pre>
item:
  - method: UPDATE
    where: &FREE_ITEMS
      - Some Coat
      - Dress Shoes
    SellPrice: 1.5
    BuyPrice: 2.5

stuff:
  - method: MERGE
    item-merge: {name: Some Value Here}
    items: *FREE_ITEMS
</pre>

<ul>
    <li>Any YAML node can be anchored and referenced elsewhere as an alias node.</li>
    <li>To anchor a particular value or set of values, use ``&name of anchor``.</li>
    <li>To reference an anchor, use ``*name of anchor``.</li>
</ul>

</details>

A service definition contains configuration which will be applied to each container started for that service, much like passing command-line parameters to `docker run`. Likewise, network and volume definitions are analogous to `docker network create` and `docker volume create`.

Options specified in the `Dockerfile` (e.g., `CMD`, `EXPOSE`, `VOLUME`, `ENV`) are respected by default - you don’t need to specify them again in `docker-compose.yml`.


## Docker Compose build files
_________________________

We've already created a simple app in `code/docker-compose` that uses Golang and redis.

1. Go to `code/docker-compose` folder

2. Review `Dockerfile`:

```dockerfile
FROM golang:1.8.3

LABEL maintainer "marcelbelmont@gmail.com"

# Set Environment variables
ENV appDir /var/www/app
ENV appMain /var/www/app/main

# Set the work directory
RUN mkdir -p ${appDir}
ADD . ${appDir}
WORKDIR ${appDir}

RUN go build

CMD ["go", "run", "main.go"]
```

1. We use the base image of `golang:1.8.3`.
2. Sets `LABEL` of maintainer
3. The `Dockerfile` then creates the directory where our code will be stored, `/var/www/app`
4. Copies all the code we have in the host machine and runs the command that will keep the container running.
5. Builds Go binary
6. Then gives main execution command

* Review `docker-compose.yml`:

```yml
version: '3'
services:
  redis:
    image: redis:alpine
  api:
    build: .
    ports:
      - "3000:3000"
    depends_on:
      - redis
```

* The `docker-compose.yml` file describes the services that make your app.
* In this example those services are an api and an in memory database.
* The compose file also describes which Docker images these services use, how they link together services
* The compose file also describes any volumes the services might need mounted inside the containers.
* Finally, the `docker-compose.yml` file describes which ports these services expose.
* See the docker-compose.yml [reference](https://docs.docker.com/compose/compose-file/)
* In this case, we defined two services, `redis` that uses `redis:alpine` and `api`, our golang app.
* We linked the two of them, and `api` depends on `redis` as you can see in `depends_on`.
* Also, our golang app listens to the port `3000` so we linked host's port 3000 to the docker container 3000 port.


## Run the Golang API
_________________________

### Build the images

The docker-compose yml script can build all the images at once by running the command:

```bash
$ docker-compose build
```

The `docker-compose build` reads the `docker-compose.yml` file and builds all the services defined in there.

### Run a command against a service
We can run a one-time command against a service. For example, the following command starts the `api` service and runs `go run main.go` as its command.

```bash
$ docker-compose run api
```

Commands you use with `run` start in new containers with the same configuration as defined by the service's configuration. This means the container has the same volumes, links, as defined in the configuration file. There two differences though.

First, the command passed by `run` overrides the command defined in the service configuration.
For example, if the `api` service configuration is started with `go`, then `docker-compose run api sh` overrides it with `sh`.

The second difference is the `docker-compose run` command does not create any of the ports specified in the service configuration. This prevents the port collisions with already open ports.

### Start services

We can run `docker-compose up` that builds, (re)creates, starts, and attaches to containers for a service. Unless they are already running, this command also starts any linked services.

Type in your terminal:

```bash
$ docker-compose up
```

This instructs Compose to run the services defined in the `docker-compose.yml` in containers, using the `redis` image and the `api` service's image and configuration.

The docker-compose up command aggregates the output of each container. When the command exits, all containers are stopped.

If we want, we can run the containers in background with `-d` flag:

```bash
$ docker-compose up -d
```

At this point, your Node app should be running at port `8088` on your Docker host. If you are using a Docker Machine VM, you can use the `docker-machine ip MACHINE_NAME` to get the IP address.

### Logs

We can see the log output from services running:

```bash
$ docker-compose logs
```

If we want to review the logs of a specific service, e.g. `api`:

```bash
$ docker-compose logs api
```

### List containers

We can run `ps` like in `docker ps` to list containers and their status:

```bash
$ docker-compose ps
        Name                       Command               State           Ports
---------------------------------------------------------------------------------------
golangexample_api_1     go run main.go                   Up      0.0.0.0:3000->3000/tcp
golangexample_redis_1   docker-entrypoint.sh redis ...   Up      0.0.0.0:6379->6379/tcp
```

* Notice here that when we run `docker-compose ps` command there are 2 services `redis` and `api`
* Remember that in our `docker-compose.yml` script we have 2 services named `redis` and `api`

### Stop containers

```bash
$ docker-compose stop
```

* This command stops running containers without removing them.
* This command will stop all the containers
* We could also stop an individual service like this `docker-compose stop redis`

```bash
$ docker-compose ps
        Name                       Command               State            Ports
----------------------------------------------------------------------------------------
golangexample_api_1     go run main.go                   Up       0.0.0.0:3000->3000/tcp
golangexample_redis_1   docker-entrypoint.sh redis ...   Exit 0
```

Now we have only 1 service running `api`

The containers can be started again with `docker-compose start`.

### Start container

Starts existing containers for a service, e.g. `api`:

```bash
$ docker-compose start api
```

### Ping Container

When can also run ping inside a container to make sure that the services are communicating correctly

```bash
$ docker-compose exec api /bin/sh
# ping redis
PING redis (172.19.0.2): 56 data bytes
64 bytes from 172.19.0.2: icmp_seq=0 ttl=64 time=0.106 ms
64 bytes from 172.19.0.2: icmp_seq=1 ttl=64 time=0.146 ms
64 bytes from 172.19.0.2: icmp_seq=2 ttl=64 time=0.149 ms
```

* Notice here that we are getting a packet back with the IP_ADDRESS: `172.19.0.2`
* Pretty cool how Docker is smart enough to alias `redis` service to that IP_ADDRESS for us right !!!

### Remove containers

```bash
$ docker-compose rm
```

**If you run `docker-compose rm -f` it will forcibly remove the containers with no verification**

The previous command removes __stopped__ service containers.

If we want to stop and remove them:

```bash
$ docker-compose down
```

## Bread Crumb Navigation
_________________________

Previous | Next
:------- | ---:
← [Volumes](../volumes/README.md) | [Docker Network](../docker-network/README.md) →
