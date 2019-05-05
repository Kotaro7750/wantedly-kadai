FROM golang:latest
WORKDIR /go/src
ADD . /go/src/
EXPOSE 8080
RUN apt-get update -y && go get github.com/gin-gonic/gin && go get github.com/gin-contrib/cors
CMD ["/usr/local/go/bin/go", "run", "/go/src/main.go"]