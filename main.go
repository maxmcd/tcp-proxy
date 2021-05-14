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
		log.Println("New connection", conn.RemoteAddr())
		if err != nil {
			log.Println("error accepting connection", err)
			continue
		}
		go func() {
			defer conn.Close()
			conn2, err := net.Dial("tcp", *remoteAddr)
			if err != nil {
				log.Println("error dialing remote addr", err)
				return
			}
			defer conn2.Close()
			closer := make(chan struct{}, 2)
			go copy(closer, conn2, conn)
			go copy(closer, conn, conn2)
			<-closer
			log.Println("Connection complete", conn.RemoteAddr())
		}()
	}
}

func copy(closer chan struct{}, dst io.Writer, src io.Reader) {
	_, _ = io.Copy(dst, src)
	closer <- struct{}{} // connection is closed, send signal to stop proxy
}
