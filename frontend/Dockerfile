# ---------------
# Build Yarn base image
# ---------------
FROM node:lts-alpine3.14 as yarn
RUN mkdir /app
WORKDIR /app
RUN npm install -g --force npm@latest typescript@latest yarn@latest

# ---------------
# Build Package repo
# ---------------
FROM yarn as repo
COPY package.json .
RUN yarn install

# ---------------
# Build webapp
# ---------------
FROM repo as build
COPY . .
RUN yarn build

# ---------------
# Copy in vanilla nginx image
# ---------------
FROM nginx:latest as server
RUN rm /etc/nginx/conf.d/default.conf
COPY nginx.conf /etc/nginx/conf.d
COPY --from=build /app/build /usr/share/nginx/html