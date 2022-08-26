
FROM golang:1.19.0-alpine3.16

RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /go/src/github.com/ghgsnaidu/go-db-demo

# We want to populate the module cache based on the go.{mod,sum} files.
#COPY go.mod .
#COPY go.sum .
#
#RUN echo $GOPATH
#RUN go mod tidy

COPY . .
RUN go mod tidy
WORKDIR /go/src/github.com/ghgsnaidu/go-db-demo/cmd/main
ENV DBURL="root:root@tcp(172.17.0.2:3306)/mydb"
# Build the Go app
RUN go build   -o app
#-o ./go-db-demo/cmd/main


# This container exposes port 3000 to the outside world

EXPOSE 3000
# Run the binary program produced by `go install`
CMD ["./app"]