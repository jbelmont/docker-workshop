FROM mhart/alpine-node:8.1.2

LABEL maintainer "marcelbelmont@gmail.com"

# Set Environment variables
ENV appDir /var/www/app

RUN apk add --no-cache make gcc g++ python bash go bzr git mercurial subversion openssh-client ca-certificates

# Set the work directory
RUN mkdir -p ${appDir}
WORKDIR ${appDir}

# Add application files
ADD . ${appDir}

# Install npm dependencies and install ava globally
RUN npm install
RUN npm install -g ava webpack webpack-dev-server yarn
RUN go get gopkg.in/mgo.v2

# Add main node execution command
CMD ["npm", "run", "dev:docker"]
