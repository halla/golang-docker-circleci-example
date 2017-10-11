Example project to get the hang of the Go language, Docker containers, continuous integration with CircleCI.



## Todo

* ~~Docker hello world~~
* ~~Hello http server~~
* Development workflow with code reload/compile
* CircleCI integration
* Production deployment


## Usage

Assuming you have docker installed, get golang docker image, build and run:

* $ docker pull golang
* $ docker build -t my-golang-app .
* $ docker run -p 8080:8080 --rm --name my-running-app my-golang-app

You should see "Starting server..." in the console, and "Hello world" in your browser if you navigate to localhost:8080


'''
build -t tags the container with the given name
-p binds your host port to docker port
--rm removes the container after exit
--name assigns a name
'''


 See also
 * The Go Programming Language https://golang.org/
 * Docker https://duckduckgo.com/?q=docker&t=canonical&ia=web
 * CircleCI https://circleci.com/
 * golang docker image https://hub.docker.com/_/golang/
