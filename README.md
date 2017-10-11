Example project to get the hang of the Go language, Docker containers, continuous integration with CircleCI.



## Todo

* -Docker hello world-
* CircleCI integration
* Production deployment
* Hello http server


## Usage

Assuming you have docker installed, get golang docker image, build and run:

* $ docker pull golang
* $ docker build -t my-golang-app .
* $ docker run -it --rm --name my-running-app my-golang-app

You should see "hello world" in the console.

 See also
 * The Go Programming Language https://golang.org/
 * Docker https://duckduckgo.com/?q=docker&t=canonical&ia=web
 * CircleCI https://circleci.com/
 * golang docker image https://hub.docker.com/_/golang/
