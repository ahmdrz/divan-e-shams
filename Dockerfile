FROM golang:1.10-alpine

LABEL maintainer zibaeiahmadreza@gmail.com

ENV GIN_MODE=release

RUN mkdir -p $GOPATH/src/github.com/ahmdrz/divan-e-shams && mkdir -p /app
ADD . $GOPATH/src/github.com/ahmdrz/divan-e-shams
WORKDIR $GOPATH/src/github.com/ahmdrz/divan-e-shams
RUN go build -i -o /app/main
EXPOSE 8080

CMD ["/app/main"]