package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Mcleod-Key", "this is from mcleod")
	res.Header().Set("Content-Type", "text/html; chaset=utf-8")
	fmt.Fprintln(res, "<h1>any code you want in this func</h1>")
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
