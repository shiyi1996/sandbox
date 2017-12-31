FROM alpine:latest
MAINTAINER ShiYi <shiyi@fightcoder.com>

RUN apk upgrade --update; \
apk add python; \
apk add g++; \
apk add git; \
apk add go;

ENV GOPATH=/workspace:/fightcoder-sandbox:/fightcoder-sandbox/deps
ADD . /fightcoder-sandbox

RUN cd /fightcoder-sandbox/;go build;mv fightcoder-sandbox sandbox

WORKDIR /workspace

CMD while true; do sleep 1; done