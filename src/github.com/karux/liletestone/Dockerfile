FROM golang:alpine AS build-env
WORKDIR /usr/local/go/src/github.com/karux/liletestone
COPY . /usr/local/go/src/github.com/karux/liletestone
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh
RUN go get ./...
RUN go build -o build/liletestone ./liletestone


FROM alpine:latest
RUN apk add --no-cache ca-certificates
COPY --from=build-env /usr/local/go/src/github.com/karux/liletestone/build/liletestone /bin/liletestone
CMD ["liletestone", "up"]
