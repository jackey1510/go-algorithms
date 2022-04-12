package main

import (
	"net/http"
)

func greet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This is an example server.\n"))
}

func main() {
	http.HandleFunc("/", greet)
	http.ListenAndServeTLS(":443", "server.crt", "server.key", nil)
}
