FROM golang:1.15-alpine3.12

RUN apk update                      \
    && apk --no-cache add           \
    gcc                             \
    musl-dev                        \
    vim

WORKDIR /var/eikan