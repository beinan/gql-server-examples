FROM golang:1.11

ADD . /go/src/github.com/beinan/gql-server-examples/gateway
WORKDIR /go/src/github.com/beinan/gql-server-examples/gateway

RUN go get github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

RUN go build
RUN go install

CMD watcher -run github.com/beinan/gql-server-examples/gateway -watch github.com/beinan/gql-server-examples/gateway