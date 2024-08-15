package main

import "net/http"

func SayHelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/hello", SayHelloHandler)
	http.ListenAndServe(":8080", nil)
}
