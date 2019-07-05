package main
import (
    "fmt"
    "net/http"
)
type msg string
func (m msg) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
   fmt.Fprint(resp, m) 
}
func main() {
    msgHandler := msg("<html>" +
                      "<head>" +
                      "<meta http-equiv=\"Content-Type\" content=\"text/html; charset=utf-8\">" +
                      "<title>Тег FRAME</title>" +
                      "</head>"+
                      "<body>"+
		      "<img src=\"https://www.sumologic.com/wp-content/uploads/2015/08/DevOps-infinity-loop.png\" alt=\"devops\">"+
		      "</body>"+
                      "</frameset>"  +
                      "</frameset>" +
                      "</html>")
    fmt.Println("Server is listening...")
    http.ListenAndServe("0.0.0.0:8181", msgHandler)
}
