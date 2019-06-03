# base image
FROM node:10.16 as react-build

# ADD yarn.lock /yarn.lock
# ADD package.json /package.json

# ENV NODE_PATH=/node_modules
# ENV PATH=$PATH:/node_modules/.bin
# RUN yarn

# # set working directory
# WORKDIR /app
# ADD . /app

# RUN yarn
# RUN yarn build

# Create a work directory and copy over our dependency manifest files.
RUN mkdir /app
WORKDIR /app
COPY /src /app/src
COPY ["package.json", "package-lock.json*", "./"]

# If you're using yarn:
#  yarn build
RUN npm install --silent && mv node_modules ../

# Expose PORT 3000 on our virtual machine so we can run our server
EXPOSE 3000