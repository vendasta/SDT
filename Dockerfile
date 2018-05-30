FROM alpine:3.5

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

COPY gopath/bin/server /server

ENTRYPOINT ["/server"]
