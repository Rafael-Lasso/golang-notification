package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func WebServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World com Go! %s", time.Now())
}


func SendSMS(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website with Golang!\n")
}

func main() {
	http.HandleFunc("/", WebServer)
	http.HandleFunc("/getRoot", SendSMS)
	http.ListenAndServe(":8080", nil)
}
