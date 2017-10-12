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
* Make a HTTP-request
* ~~CircleCI integration~~
* Production deployment


## Usage

Assuming you have docker installed, get golang docker image, build and run:

* $ docker pull golang
* $ docker build -t hello-golang .
* $ docker run --volume=$PWD:/go/src/app -p 8080:3000 --rm --name my-running-app hello-golang

You should see "Starting server..." in the console, and "Hello world" in your browser if you navigate to localhost:8080

> build -t tags the container with the given name
> --volume mounts the current directory to the docker container for code reloading
> -p binds your host port to docker port
> --rm removes the container after exit
> --name assigns a name



 See also
 * The Go Programming Language https://golang.org/
 * Docker https://www.docker.com/
 * CircleCI https://circleci.com/
 * golang docker image https://hub.docker.com/_/golang/
