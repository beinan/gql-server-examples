FROM golang:1.11

ADD . /go/src/github.com/beinan/gql-server-examples/user-service
WORKDIR /go/src/github.com/beinan/gql-server-examples/user-service

RUN go get github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

RUN go build
RUN go install

CMD watcher -run github.com/beinan/gql-server-examples/user-service -watch github.com/beinan/gql-server-examples/user-service