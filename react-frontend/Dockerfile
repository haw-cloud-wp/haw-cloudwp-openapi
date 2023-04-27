FROM node as repo
WORKDIR /app
COPY package.json .
RUN npm install

FROM repo as build
WORKDIR /app
COPY . .
RUN rm nginx.conf
RUN npm run build

FROM nginx as server
RUN rm /etc/nginx/conf.d/default.conf
COPY nginx.conf /etc/nginx/conf.d
COPY --from=build /app/build /usr/share/nginx/html