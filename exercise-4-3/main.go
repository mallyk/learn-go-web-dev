package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

var tpl *template.Template

type data struct {
	Method        string
	URL           *url.URL
	Submissions   url.Values
	Header        http.Header
	Host          string
	ContentLength int64
	Name          string
}

func index(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	data := data{
		Method:        req.Method,
		URL:           req.URL,
		Submissions:   req.Form,
		Header:        req.Header,
		Host:          req.Host,
		ContentLength: req.ContentLength,
	}

	tpl.ExecuteTemplate(w, "index.gohtml", data)
}

func dog(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	data := data{
		Method:        req.Method,
		URL:           req.URL,
		Submissions:   req.Form,
		Header:        req.Header,
		Host:          req.Host,
		ContentLength: req.ContentLength,
		Name:          "Doggo",
	}

	tpl.ExecuteTemplate(w, "index.gohtml", data)
}

func me(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	data := data{
		Method:        req.Method,
		URL:           req.URL,
		Submissions:   req.Form,
		Header:        req.Header,
		Host:          req.Host,
		ContentLength: req.ContentLength,
		Name:          "Kyle",
	}

	tpl.ExecuteTemplate(w, "index.gohtml", data)
}

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {

	http.Handle("/", http.HandlerFunc(index))
	http.Handle("/dog/", http.HandlerFunc(dog))
	http.Handle("/me/", http.HandlerFunc(me))

	http.ListenAndServe(":8080", nil)
}
