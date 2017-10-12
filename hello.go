package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Starting a server listening at port 8080")
	http.HandleFunc("/", hello)
	port := os.Getenv("PORT")
	http.ListenAndServe(":"+port, nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, Message())
}

// Uppercased functions are exported
func Message() string {
	return "Hello world!"
}
