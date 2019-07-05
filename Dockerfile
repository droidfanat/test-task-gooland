FROM ubuntu:18.04

RUN apt-get update && apt-get install -y golang-go
WORKDIR /go
COPY ./server.go ./ 

CMD ["/usr/bin/go", "run", "server.go"]
