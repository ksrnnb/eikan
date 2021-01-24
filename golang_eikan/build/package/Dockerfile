FROM golang:1.15-alpine3.12

RUN apk update                      \
    && apk --no-cache add           \
    gcc                             \
    git                             \
    musl-dev                        \
    vim
# gitは go getに必要

WORKDIR /var/eikan