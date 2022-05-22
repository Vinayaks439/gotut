package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net"
)

var (
	addr = flag.String("addr", "127.0.0.1:8080", "IP:port")
)

type Message struct {
	Body string
}

func main() {
	flag.Parse()
	listen, err := net.Listen("tcp", *addr)
	if err != nil {
		panic(err)
	}
	for {
		c, err := listen.Accept()
		if err != nil {
			panic(err)
		}
		go parseclients(c)
	}
}

func parseclients(c net.Conn) {
	encode := json.NewDecoder(c)
	for {
		var message []Message
		err := encode.Decode(&message)
		if err != nil {
			panic(err)
		}
		fmt.Println(message)
	}
}
