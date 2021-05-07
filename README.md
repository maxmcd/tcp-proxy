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
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
)

var localAddr *string = flag.String("l", "localhost:9999", "local address")
var remoteAddr *string = flag.String("r", "localhost:80", "remote address")

func main() {
	flag.Parse()
	fmt.Printf("Listening: %v\nProxying: %v\n\n", *localAddr, *remoteAddr)

	listener, err := net.Listen("tcp", *localAddr)
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		log.Println("Got new connection", conn.RemoteAddr())
		if err != nil {
			log.Println("error accepting connection", err)
			continue
		}
		go func() {
			conn2, err := net.Dial("tcp", *remoteAddr)
			if err != nil {
				log.Println("error dialing remote addr", err)
				return
			}
			go io.Copy(conn2, conn)
			io.Copy(conn, conn2)
			log.Println("Connection complete", conn.RemoteAddr())
			conn2.Close()
			conn.Close()
		}()
	}
}
```
