FROM golang:1.17

# Mac用
# PROTOCOL := "tcp(nicoreco:3306)"

RUN apk update && \
    apk upgrade && \
    apk add git && \
    go get -u github.com/gofiber/fiber && \
    go get -u github.com/dgrijalva/jwt-go

RUN go get github.com/cespare/reflex
ENV CGO_ENABLED=0

WORKDIR /go/src/api                                                                                             
COPY ./api ./

RUN go mod download