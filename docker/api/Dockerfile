FROM golang:1.10.2-alpine3.7

COPY src/api /go/src/api/

WORKDIR /go/src/api/

RUN apk update \
    && apk add --no-cache git \
    && go get -u github.com/codegangsta/gin \
    && go get -u github.com/golang/dep/cmd/dep \
    && go get -u github.com/lib/pq \
    && go get -u github.com/jinzhu/gorm \
    && dep ensure

EXPOSE 5000
CMD gin -i run