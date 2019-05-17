package main

import (
	"io"
	"net/http"
)

func index(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello")
}

func dog(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "doggo")
}

func me(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Kyle")
}

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)

	http.ListenAndServe(":8080", nil)
}
