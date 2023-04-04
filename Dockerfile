FROM golang

WORKDIR /go/src/github.com/olivere/go-container-debugging

EXPOSE 8080 2345

RUN go install github.com/go-delve/delve/cmd/dlv@v1.20.1

ADD . /go/src/github.com/olivere/go-container-debugging

CMD ["dlv", "debug", "github.com/olivere/go-container-debugging", "--headless", "--listen=:2345", "--api-version=2", "--log"]
