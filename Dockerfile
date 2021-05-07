FROM golang

WORKDIR /opt

COPY main.go go.mod ./

RUN env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
		go build -tags netgo -ldflags '-w' -o tcp-proxy . \
    && strip tcp-proxy

FROM alpine
COPY --from=0 /opt/tcp-proxy /bin/tcp-proxy

