package main

import (
	"fmt"
	"net/http"
)

func getGreeting() string {
	return "Hello, Kontur!"
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, getGreeting())
}

func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8080", nil)
}
