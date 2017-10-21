package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"time"
)

// Project type is used to extract fields from Github project json object.
// Field names have to be uppercase (exported) for json decoder to work.
type Project struct {
	Name     string `json:"name"`
	PushedAt string `json:"pushed_at"`
}

type Page struct {
	Title string
	Body  string
}

const githubProjectBaseURL = "https://api.github.com/repos/"

const projectName = "halla/golang-docker-circleci-example"

// slices are popular data structures in go
var projectNames = []string{
	"golang/go",
	"docker/docker-ce",
}

// there is no timeout by default!
var httpClient = &http.Client{Timeout: 10 * time.Second}

// main is the entry point to the application by convention
func main() {
	fmt.Println("Starting a server listening at port 8080")
	http.HandleFunc("/", hello)
	port := os.Getenv("PORT")
	http.ListenAndServe(":"+port, nil)
}

func FormatDate(rfc3339date string) string {
	t1, err := time.Parse(time.RFC3339, rfc3339date)
	if err != nil {
		return "Error parsing date: " + rfc3339date
	}
	return t1.Format(time.RFC1123)
}

// lowercase functions are internal to the package
func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>golang-docker-circleci-example</h1>")

	page := new(Page)
	page.Title = "golang-docker-circleci-example"
	page.Body = "Body here"

	// synchronous http request
	project := new(Project)
	projectURL := githubProjectBaseURL + projectName
	fmt.Println("Requesting info for " + projectURL)

	getJSON(projectURL, project)
	io.WriteString(w, Message(project))

	// async http requests, response through channel
	ch := make(chan *Project)
	for _, projectName := range projectNames {
		go getProjectAsync(githubProjectBaseURL+projectName, ch) // start an async go routine
	}
	funcMap := template.FuncMap{
		"FormatDate": FormatDate,
	}
	t, err := template.New("project-template.html").Funcs(funcMap).ParseFiles("project-template.html")
	if err != nil {
		fmt.Println("Error parsing project template.")
	}
	for range projectNames {
		project = <-ch
		t.Execute(w, project)
		fmt.Println("Received project info for " + project.Name)
	}
	fmt.Println(project)

}

// Uppercased functions are exported
func Message(project *Project) string {
	return "<p>" + project.Name + ": Latest commit: " + project.PushedAt + "</p>"
}

func getJSON(url string, target interface{}) error {
	resp, err := httpClient.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(target)
}

func getProjectAsync(url string, ch chan<- *Project) {
	fmt.Println("Requesting info for " + url)
	resp, err := httpClient.Get(url)
	target := new(Project)
	if err != nil {
		target.Name = "Unknown"
		target.PushedAt = "Unknown"
		ch <- target
		return
	}
	json.NewDecoder(resp.Body).Decode(target)
	ch <- target
	defer resp.Body.Close()
	return
}
