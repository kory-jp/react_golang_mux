FROM node:16-alpine
# RUN mkdir -p /app/client/client-app
# WORKDIR /app/client/client-app
# COPY client-app/package.json .
# COPY client-app/tsconfig.json .
# COPY client-app/yarn.lock .
# RUN yarn install --production
# COPY . .
# EXPOSE 3000
# CMD ["yarn", "build"]

RUN mkdir -p /app/client
WORKDIR /app/client
COPY ./client-app .
EXPOSE 3000
CMD ["yarn", "build"]