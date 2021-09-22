package main
import (
    "time"
    "log"
    "net/http" // HTTP protocol package
    "encoding/json"
)
 
func TimeResponseHandler(w http.ResponseWriter, r *http.Request) {
    test := &Time{Time: time.Now().Format(time.RFC3339)} // get current time with our format   
    res := json.Marshal(test)
    w.Header().Set("content-type", "aplication/json")
    w.WriteHeader(200)
    w.Write(res)
}

func main() {
    type Time struct {
        Time string `json:"time"`
    }
    http.HandleFunc("/time", TimeResponseHandler) {
        log.Println("Starting HTTP server")
        err := http.ListenAndServe(":3000", nil)
        if err != nil {
            log.Fatal("ListenAndServe: ", err)
        }
    }
}
