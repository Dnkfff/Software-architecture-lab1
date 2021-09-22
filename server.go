package main
import (
    "time"
    "os" // file usage package
    "fmt" // formatted input output package
    "net/http" // HTTP protocol package
    "log"
)
 

func main() {
    type Time struct {
        Time string `json:"time"`
    }
    fmt.Fprintln(os.Stdout, "hello cold")
}
