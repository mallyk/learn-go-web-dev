package main

import (
	"net/http"

	"gopkg.in/mgo.v2"

	"github.com/julienschmidt/httprouter"
	"github.com/mallyk/learn-go-web-dev/035-mongo/controllers"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r)

}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}

	return s
}
