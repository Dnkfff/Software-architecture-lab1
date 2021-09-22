package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http" // HTTP protocol package
	"time"
)

func TimeResponseHandler(w http.ResponseWriter, req *http.Request) {
	output := &Time{Time: time.Now().Format(time.RFC3339)} // get current time with our format
	res, _ := json.Marshal(output)

	fmt.Fprintf(w, "%s\n", res)
	//w.Header().Set("content-type", "aplication/json") // downloading page content
	//w.WriteHeader(200)
	//w.Write(res)
}

type Time struct {
	Time string `json:"time"`
}

func main() {
	log.Printf("HTTP server started")
	http.HandleFunc("/time", TimeResponseHandler)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
