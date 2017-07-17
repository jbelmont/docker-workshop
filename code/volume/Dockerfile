FROM mhart/alpine-node:8.0.0

LABEL Marcel Belmont "marcelbelmont@gmail.com"

# Set Environment variables
ENV appDir /var/www/app
ENV volumeDir /var/www/app/public
ENV appNodeModules /var/www/app/node_modules

# Set the work directory
RUN mkdir -p ${appDir}
WORKDIR ${appDir}

COPY package.json ${appDir}
RUN npm install

COPY . ${appDir}

# Define mountable directories.
VOLUME [${volumeDir}, ${appNodeModules}]

EXPOSE 3006 3006

# Add main node execution command
CMD ["npm", "start"]
