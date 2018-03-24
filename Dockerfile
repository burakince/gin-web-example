FROM golang:1.10.0-alpine AS build

RUN apk --no-cache add git
RUN go get -d -v github.com/gin-gonic/gin github.com/gin-contrib/location
ADD . /go/src/github.com/burakince/gin-web-example
RUN go install github.com/burakince/gin-web-example

FROM alpine:latest

RUN apk --no-cache add --update \
  ca-certificates
RUN mkdir /http-server
WORKDIR /http-server

COPY --from=build /go/bin/gin-web-example /http-server/gin-web-example
COPY --from=build /go/src/github.com/burakince/gin-web-example/templates /http-server/templates

ENTRYPOINT ["/http-server/gin-web-example"]
