FROM ubuntu:18.04

RUN apt-get update && apt-get install -y golang-go
WORKDIR /go
COPY ./server.go ./ 

EXPOSE 8181
CMD ["/usr/bin/go", "run", "server.go"]
