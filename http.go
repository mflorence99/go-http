package main

import (
    "fmt"
		"net/http"
		"os"
)

func main() {

    fmt.Printf("My host=%s port=%s\n\n", os.Getenv("HTTP_HOST"), os.Getenv("HTTP_PORT"))

    http.HandleFunc("/", HelloServer)
    // NOTE: listen on all interfaces
    // @see https://stackoverflow.com/questions/40063120
    http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("HTTP_PORT")), nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("Received request...\n")
    fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
