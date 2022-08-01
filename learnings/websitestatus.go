package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	c := make(chan bool)
	for {
		go ping("https://google.com", c)
		time.Sleep(5 * time.Second)
		status := <-c
		if status {
			break
		}
	}
}

func ping(url string, c chan bool) {
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		c <- false
	}
	code := resp.StatusCode
	if code != 200 {
		log.Println(url + " is down")
		c <- false
	} else {
		log.Println(url + " is up")
		c <- true
	}
}
