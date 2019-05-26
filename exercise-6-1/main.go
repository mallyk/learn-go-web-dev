package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func set(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("counting-my-cookie")
	if err != nil {
		http.SetCookie(w, &http.Cookie{
			Name:  "counting-my-cookie",
			Value: "0",
		})
		fmt.Println("cookie wasn't found")
	} else {
		count, err := strconv.Atoi(c.Value)
		if err != nil {
			count = 0
			c.Value = strconv.Itoa(count)
			fmt.Println("value was not an integer")
		}
		count++
		c.Value = strconv.Itoa(count)
		http.SetCookie(w, c)
	}

	fmt.Fprintln(w, "COOKIE WRITTEN - CHECK YOUR BROWSER")
}

func read(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("counting-my-cookie")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	fmt.Fprintln(w, "YOUR COOKIE:", c)
}
