FROM golang:1.8.1-alpine

RUN set -ex && apk add --no-cache bash bzr git curl mercurial subversion openssh-client ca-certificates jq
RUN go get github.com/constabulary/gb/...
