FROM golang:1.8.3

LABEL maintainer "marcelbelmont@gmail.com"

# Set Environment variables
ENV appDir /var/www/app
ENV appMain /var/www/app/main

# Set the work directory
RUN mkdir -p ${appDir}
ADD . ${appDir}
WORKDIR ${appDir}

RUN go get github.com/garyburd/redigo/redis
RUN go build

CMD ["go", "run", "main.go"]
