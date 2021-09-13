package main

import (
	"github.com/tymoteuszgach/lunch-and-go/hello"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", hello.Handler)
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
