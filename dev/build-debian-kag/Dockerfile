FROM golang:1.8.1

RUN apt update && apt install -y bash grep bzr git curl mercurial subversion openssh-client ca-certificates jq
RUN go get github.com/constabulary/gb/...
