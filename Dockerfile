FROM ubuntu:18.04

RUN apt-get update && apt-get install -y golang-go
WORKDIR /go
COPY ./server.go ./ 
RUN go build server.go

EXPOSE 8181
CMD ["/go/server"]
