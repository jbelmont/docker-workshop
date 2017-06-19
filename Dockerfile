FROM golang
FROM mhart/alpine-node:8.1.2

LABEL Marcel Belmont "marcelbelmont@gmail.com"

# Set Environment variables
ENV appDir /var/www/app

# Set the work directory
RUN mkdir -p ${appDir}
WORKDIR ${appDir}

# Add application files
ADD . ${appDir}

# Install mgo mongodb driver for golang
RUN go get gopkg.in/mgo.v2

# Install npm dependencies and install ava globally
RUN npm install
RUN npm install -g ava

# Add main node execution command
CMD ["npm", "run", "dev"]
