Example project to get the hang of the Go language, Docker containers, continuous integration with CircleCI.

This is built on Ubuntu Xenial (16.04). On Windows/Mac some tweaks probably needed.

Build status: [![CircleCI](https://circleci.com/gh/halla/golang-docker-circleci-example/tree/master.svg?style=svg)](https://circleci.com/gh/halla/golang-docker-circleci-example/tree/master)

## Todo

* ~~Docker hello world~~
* ~~Hello http server~~
* ~~Unit test case~~
* ~~Development workflow with code reload/compile~~
* Use a goroutine
* Use a channel
* ~~Make a HTTP-request~~
* ~~Parse JSON~~
* Use an HTML template
* Serve static files
* ~~CircleCI integration~~
* Production deployment


## Usage

Assuming you have docker installed, clone this repo, cd to the project directory, build and run:

* $ docker build -t hello-golang .
* $ docker run --volume=$PWD:/go/src/app -p 8080:3000 --rm --name my-running-app hello-golang

You should see "Starting server..." in the console, and "Hello world" in your browser if you navigate to localhost:8080

> build -t tags the container with the given name
> --volume mounts the current directory to the docker container for code reloading
> -p binds your host port to docker port
> --rm removes the container after exit
> --name assigns a name



## Technology links

 * The Go Programming Language https://golang.org/
 * Docker https://www.docker.com/
 * CircleCI https://circleci.com/
 * golang docker image https://hub.docker.com/_/golang/
 * gin for code reloading https://github.com/codegangsta/gin

## Tutorials

* Install Docker on Ubuntu 16.04 https://www.digitalocean.com/community/tutorials/how-to-install-and-use-docker-on-ubuntu-16-04
* Development with Go and Docker https://medium.com/developers-writing/docker-powered-development-environment-for-your-go-app-6185d043ea35
