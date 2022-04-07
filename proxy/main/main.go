package main

import (
	"flag"
	"log"
	"net"
	"net/http"

	"github.com/utopizza/go-learn/proxy"
)

func startHTTPProxy() {
	var addr = flag.String("addr", "127.0.0.1:8080", "The addr of the application.")
	flag.Parse()

	handler := &proxy.HTTPProxy{}

	log.Println("Starting proxy server on", *addr)
	if err := http.ListenAndServe(*addr, handler); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func startTCPProxy() {
	p := &proxy.TCPProxy{}

	// tcp连接，监听8080端口
	l, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Panic(err)
	}

	// 死循环，每当遇到连接时，调用handle
	for {
		client, err := l.Accept()
		if err != nil {
			log.Panic(err)
		}

		go p.Handle(client)
	}
}

func main() {
	//startHTTPProxy()
	startTCPProxy()
}
