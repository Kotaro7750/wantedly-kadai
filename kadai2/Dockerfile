FROM golang:latest as builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
WORKDIR /app
COPY . .
RUN go get github.com/gin-contrib/cors
RUN go get github.com/gin-gonic/gin
RUN go build main.go


CMD /app/main $PORT
