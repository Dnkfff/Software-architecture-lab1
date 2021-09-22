package main
import (
    "time"
    "os" // file usage package
    "net/http" // HTTP protocol package
    "log"
)
 
func TimeResponseHandler(w http.ResponseWriter, r *http.Request) {
    test := &Time{Time: time.Now().Format(time.RFC3339)} // get current time with our format   
}
func main() {
    type Time struct {
        Time string `json:"time"`
    }
    http.HandleFunc("/time",)
    
}
