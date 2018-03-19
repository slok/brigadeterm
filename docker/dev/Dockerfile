FROM golang:1.10-alpine

RUN apk --no-cache add \
    g++ \
    git \
    bash

# Mock creator
RUN go get -u github.com/vektra/mockery/.../

# Create user
ARG uid=1000
ARG gid=1000
RUN addgroup -g $gid brigadeterm && \
    adduser -D -u $uid -G brigadeterm brigadeterm && \
    chown brigadeterm:brigadeterm -R /go


USER brigadeterm
WORKDIR /go/src/github.com/slok/brigadeterm