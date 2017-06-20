FROM mhart/alpine-node:8.0.0

LABEL maintainer "marcelbelmont@gmail.com"

# Set Environment variables
ENV appDir /var/www/app

RUN apk add --no-cache make gcc g++ python bash go bzr git mercurial subversion openssh-client ca-certificates

# Set the work directory
RUN mkdir -p ${appDir}
WORKDIR ${appDir}

COPY . ${appDir}
COPY package.json ${appDir}/package.json

# Install npm dependencies and install ava globally
RUN npm install
RUN go get gopkg.in/mgo.v2
RUN go get github.com/gorilla/mux
RUN go get github.com/jbelmont/docker-workshop/routes
RUN go get github.com/jbelmont/docker-workshop/handlers
RUN go get github.com/jbelmont/docker-workshop/model

# Add main node execution command
CMD ["npm", "run", "dev:docker"]
