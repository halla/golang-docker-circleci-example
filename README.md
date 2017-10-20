Example project to get the hang of the Go language, Docker containers, continuous integration with CircleCI.

This is built on Ubuntu Xenial (16.04). On Windows/Mac some tweaks probably needed.

Build status: [![CircleCI](https://circleci.com/gh/halla/golang-docker-circleci-example/tree/master.svg?style=svg)](https://circleci.com/gh/halla/golang-docker-circleci-example/tree/master)

## Core Todo

* ~~Docker hello world~~
* ~~Hello http server~~
* ~~Unit test case~~
* ~~Development workflow with code reload/compile~~
* ~~Make a HTTP-request~~
* ~~Parse JSON~~
* ~~CircleCI integration~~
* ~~Production deployment~~
* ~~Contiuous deployment~~

## More Todo
* ~~Use a goroutine~~
* ~~Use a channel~~
* ~~Use an HTML template~~
* Serve static files
* docker-compose.yml



## Usage

### Development

Assuming you have docker installed, clone this repo, cd to the project directory, build and run:

* $ docker build -t hello-golang .
* $ docker run --volume=$PWD:/go/src/app -p 8080:3000 --rm --name my-running-app hello-golang

You should see "Starting server..." in the console, and "Hello world" in your browser if you navigate to localhost:8080

> build -t tags the container with the given name
> --volume mounts the current directory to the docker container for code reloading
> -p binds your host port to docker port
> --rm removes the container after exit
> --name assigns a name

If you modify the source code, you should see the changes when refreshing your browser window.

### Deployment

To deploy the project using the image in Docker Hub registry

* install Docker on your server
* docker pull ahalla/golang-docker-circleci-example
* docker run -d -p 8080:3000 --rm --name golang-docker-circleci-example ahalla/golang-docker-circleci-example

This will run the container as a background job, listening to the OS port 8080. The container will be removed automatically when killed.

### Contiuous deployment

To automatically update the running application, you can do something along the lines of:

```bash
#/bin/bash
echo "Pulling the latest version from registry"
docker pull ahalla/golang-docker-circleci-example
echo "Killing the running container"
docker kill golang-docker-circleci-example
echo "Preparing to run.."
docker run -d -p 8080:3000 --rm --name golang-docker-circleci-example ahalla/golang-docker-circleci-example
```

Here we give the container a name so we can kill it later. This script can be either triggered from the CI or run
in cron for periodical updates (you should probably add a check whether a new version was actually pulled or not).

This is a naive solution, but probably enough to start with.


## Technology links

 * The Go Programming Language https://golang.org/
 * Docker https://www.docker.com/
 * CircleCI https://circleci.com/
 * golang docker image https://hub.docker.com/_/golang/
 * gin for code reloading https://github.com/codegangsta/gin

## Tutorials

* Install Docker on Ubuntu 16.04 https://www.digitalocean.com/community/tutorials/how-to-install-and-use-docker-on-ubuntu-16-04
* Development with Go and Docker https://medium.com/developers-writing/docker-powered-development-environment-for-your-go-app-6185d043ea35
