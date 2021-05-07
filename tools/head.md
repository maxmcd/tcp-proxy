# TCP Proxy

A simple tcp proxy written in go.

Use like:

```bash
# with docker
docker run -p 20000:20000 maxmcd/tcp-proxy tcp-proxy -l 0.0.0.0:2000 -r 100.111.145.101:2000

# install with go
go get github.com/maxmcd/tcp-proxy

# and then use it
tcp-proxy -l 0.0.0.0:2000 -r 100.111.145.101:2000
```

Full source:

```go
