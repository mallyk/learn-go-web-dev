package main

import (
	"log"
	"net/http"
	"text/template"
)

func main() {
	fs := http.FileServer(http.Dir("."))
	http.HandleFunc("/", foo)
	http.Handle("/public/", fs)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("./templates/index.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}
