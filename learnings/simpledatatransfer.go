package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"net"
	"os"
)

type Message struct {
	Body string
}

var (
	addr = flag.String("addr", "127.0.0.1:443", "Please provide the address:port")
)

func main() {
	flag.Parse()
	conn, err := net.Dial("tcp", *addr)
	if err != nil {
		panic(err)
	}
	x := bufio.NewScanner(os.Stdin)
	gg := json.NewEncoder(conn)

	for x.Scan() {
		m := Message{Body: x.Text()}
		err := gg.Encode(m)
		if err != nil {
			panic(err)
		}
	}

}
