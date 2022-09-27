FROM node:16-alpine

ENV PROJECT_ROOTDIR /app/

WORKDIR $PROJECT_ROOTDIR

COPY package.json yarn.lock $PROJECT_ROOTDIR

RUN apk update && \
    apk add --no-cache chromium && \
    rm -f /var/cache/apk/*

RUN yarn install

COPY . $PROJECT_ROOTDIR

EXPOSE 3000
ENV HOST 0.0.0.0

CMD ["yarn", "dev"]
