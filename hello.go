package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const projectURL = "https://api.github.com/repos/halla/golang-docker-circleci-example"

// there is no timeout by default!
var httpClient = &http.Client{Timeout: 10 * time.Second}

func main() {

	fmt.Println("Starting a server listening at port 8080")
	http.HandleFunc("/", hello)
	port := os.Getenv("PORT")
	http.ListenAndServe(":"+port, nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	project := new(Project)
	getJSON(projectURL, project)
	fmt.Println(project)
	io.WriteString(w, "<h1>golang-docker-circleci-example</h1>")
	io.WriteString(w, Message(project))
}

// Project type is used to extract fields from Github project json object.
// Field names have to be uppercase (exported) for json decoder to work.
type Project struct {
	PushedAt string `json:"pushed_at"`
}

// Uppercased functions are exported
func Message(project *Project) string {
	return "Hello world! Latest commit: " + project.PushedAt
}

func getJSON(url string, target interface{}) error {
	resp, err := httpClient.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(target)
}
