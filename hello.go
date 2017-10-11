package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Starting a server listening at port 8080")
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}
