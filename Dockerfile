# consider using alpine-version in production
FROM golang:1.9

WORKDIR /go/src/app
COPY . .

RUN go-wrapper download # "go get -d -v ./..."
RUN go-wrapper install # "go install -v ./..."

RUN go get github.com/codegangsta/gin

# Gin assumes PORT environment var is set
ENV PORT 8080

CMD ["gin", "run"]
# CMD ["go-wrapper", "run"] # ["app"]
