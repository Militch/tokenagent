# syntax=docker/dockerfile:1

FROM golang:1.18.0-alpine AS Builder

RUN sed -i "s/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g" /etc/apk/repositories
ENV GOPROXY=https://proxy.golang.com.cn,direct

RUN apk add git build-base

WORKDIR /usr/src/app

COPY go.mod go.sum ./ 

RUN go mod download -x && go mod verify

COPY . .
RUN go build -a -o app .

FROM golang:1.18.0-alpine

COPY --from=Builder /usr/src/app/app /usr/local/bin

WORKDIR /etc/tokenagent
COPY tokenagent.yaml ./config.yml

CMD ["app", "daemon", "-C /etc/tokenagent/config.yml"]
