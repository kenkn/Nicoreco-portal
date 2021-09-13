FROM node:14.9.0-alpine3.10

CMD [ "/bin/sh" ]

WORKDIR /app

RUN apk update && \
    npm install -g vue-cli && \
    npm i axios && \
    npm install bootstrap && \
    npm install @popperjs/core &&\
    npm install vee-validate@next &&\
    npm install yup


ADD ./ui/package.json /app

RUN npm install