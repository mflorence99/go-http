package main

import (
    "fmt"
		"net/http"
		"os"
)

func main() {
    http.HandleFunc("/", HelloServer)
    http.ListenAndServe(":5000", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %s -- config=%s secret=%s", r.URL.Path[1:], os.Getenv("HTTP_CONFIG"), os.Getenv("HTTP_CONFIG"))
}
