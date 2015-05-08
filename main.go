package main

import (
	"net/http"
)

func handler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Demo of Face++\n"))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":80", nil)
}
