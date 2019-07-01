FROM golang:latest

WORKDIR /go/src/github.com/rconway/goapp
COPY ./goapp/. .

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 8080/tcp

CMD ["goapp"]
