FROM golang:1.11

ADD . /go/src/github.com/beinan/gql-server-examples/vlog/vlog-server
WORKDIR /go/src/github.com/beinan/gql-server-examples/vlog/vlog-server

RUN go get github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

RUN go build
RUN go install

CMD GOFLAGS=-mod=vendor watcher -run github.com/beinan/gql-server-examples/vlog/vlog-server -watch github.com/beinan/gql-server-examples/vlog/vlog-server