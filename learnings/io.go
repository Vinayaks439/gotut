package main

import (
	"bufio"
	"encoding/json"
	"os"
)

type Message struct {
	Body string
}

func main() {
	x := bufio.NewScanner(os.Stdin)
	gg := json.NewEncoder(os.Stdout)

	for x.Scan() {
		m := Message{Body: x.Text()}
		err := gg.Encode(m)
		if err != nil {
			panic(err)
		}
	}

}
