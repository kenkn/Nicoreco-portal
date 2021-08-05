FROM node:14.9.0-alpine3.10

CMD [ "/bin/sh" ]

WORKDIR /app

RUN apk update && \
    npm install -g vue-cli && \
    npm i axios

ADD ./ui/package.json /app


RUN npm install --save --legacy-peer-deps sass-loader@10.1.1 sass webpack && \
    npm install