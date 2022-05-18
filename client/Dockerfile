FROM node:16-alpine as builder
RUN mkdir -p /app/client
WORKDIR /app/client

COPY ./client-app/package.json .
COPY ./client-app/yarn.lock .
COPY ./client-app/tsconfig.json .
RUN yarn install --production
COPY ./client-app/ .
RUN yarn build

FROM nginx
COPY --from=builder /app/client/build /usr/share/nginx/html
COPY nginx.conf /etc/nginx/nginx.conf
EXPOSE 3000