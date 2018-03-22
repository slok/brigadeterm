FROM golang:1.10-alpine

ARG ostype=Linux

RUN apk --no-cache add \
    g++ \
    git \
    bash

# Mock creator
RUN go get -u github.com/vektra/mockery/.../

# Create user
ARG uid=1000
ARG gid=1000

RUN bash -c 'if [ ${ostype} == Linux ]; then addgroup -g $gid brigadeterm; else addgroup brigadeterm; fi && \
    adduser -D -u $uid -G brigadeterm brigadeterm && \
    chown brigadeterm:brigadeterm -R /go'

USER brigadeterm
WORKDIR /go/src/github.com/slok/brigadeterm