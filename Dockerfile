FROM mhart/alpine-node:8.0.0

LABEL maintainer "marcelbelmont@gmail.com"

# Set Environment variables
ENV appDir /var/www/app

RUN apk add --no-cache make gcc g++ python bash bzr git subversion openssh-client ca-certificates

# Set the work directory
RUN mkdir -p ${appDir}
WORKDIR ${appDir}

COPY . ${appDir}
COPY package.json ${appDir}/package.json

# Install npm dependencies and install ava globally
RUN npm install
RUN npm rebuild node-sass

# Add main node execution command
CMD ["npm", "run", "dev:docker"]
