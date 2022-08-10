package main

import (
    "fmt"
    "net/http"
)

func main() {
//    http.HandleFunc("/", HelloServer)
  //  http.ListenAndServe(":8080", nil)
if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
