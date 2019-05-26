package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/abundance", abundance)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func set(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "my-cookie",
		Value: "some value",
		Path:  "/",
	})
	fmt.Fprintln(w, "COOKIE WRITTEN - CHECK YOUR BROWSER")
}

func read(w http.ResponseWriter, req *http.Request) {
	readCookies(w, req, "my-cookie")
	readCookies(w, req, "general")
	readCookies(w, req, "specific")
}

func readCookies(w http.ResponseWriter, req *http.Request, cookieName string) {
	c, err := req.Cookie(cookieName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	fmt.Fprintln(w, "YOUR COOKIE:", c)
}

func abundance(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "general",
		Value: "some other value about general things",
	})

	http.SetCookie(w, &http.Cookie{
		Name:  "specific",
		Value: "some other value about specific things",
	})

	fmt.Fprintln(w, "COOKIES WRITTEN -CHECK YOUR BROWSER")
}
