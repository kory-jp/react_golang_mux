FROM node:16-alpine
RUN mkdir -p /app/client
COPY . /app/client
WORKDIR /app/client/client-app
RUN npm install --save prop-types
RUN npm install -g create-react-app
RUN yarn install && \
  yarn build
EXPOSE 3000
CMD ["yarn", "start"]