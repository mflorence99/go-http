go build http.go
docker build . -t $DOCKER_USERNAME/go-http
# docker push $DOCKER_USERNAME/go-http
