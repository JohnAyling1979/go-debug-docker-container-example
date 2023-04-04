FROM golang

WORKDIR /go/src/go-container-debugging

EXPOSE 8080 2345

RUN go install github.com/go-delve/delve/cmd/dlv@v1.20.1

CMD ["dlv", "debug", "go-container-debugging", "--headless", "--listen=:2345", "--api-version=2", "--log"]
