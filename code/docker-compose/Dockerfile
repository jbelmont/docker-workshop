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

RUN go build -o main .
CMD [${appMain}]
