package main

import (
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/contact", contact)
	http.HandleFunc("/about", about)
	http.HandleFunc("/apply", apply)
	http.ListenAndServe(":8080", nil)
}

func about(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "about.gohtml", nil)
}

func contact(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "contact.gohtml", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func apply(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		tpl.ExecuteTemplate(w, "applyProcess.gohtml", nil)
		return
	}

	tpl.ExecuteTemplate(w, "apply.gohtml", nil)
}
