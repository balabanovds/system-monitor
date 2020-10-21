FROM node:lts-alpine as build
WORKDIR /app
COPY . . 
RUN rm -rf node_modules && npm install 
RUN npm run build

FROM node:lts-alpine
RUN npm install -g http-server
WORKDIR /app
COPY --from=build /app/dist .
EXPOSE $CLIENT_PORT
CMD [ "http-server", "." ]