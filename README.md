Mini webserver on gooland

for run 
  docker build -t erik/go-test:latest .
  docker run -tid -p 8181:8181 erik/go-test 
