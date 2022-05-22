package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", Hello)
	http.HandleFunc("/update", updatevdo)
	http.ListenAndServe(":7900", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {

	videos := getVideos()

	videoBytes, err := json.Marshal(videos)

	if err != nil {
		panic(err)
	}
	w.Write(videoBytes)

}

func updatevdo(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			panic(err)
		}

		var vidoes []video
		err = json.Unmarshal(body, &vidoes)
		if err != nil {
			w.WriteHeader(400)
			fmt.Fprintf(w, "Bad Request")
		}

		saveVideos(vidoes)

	} else {
		w.WriteHeader(405)
		fmt.Fprint(w, "Method not supported")
	}
}
