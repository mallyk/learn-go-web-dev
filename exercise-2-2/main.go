package main

import (
	"html/template"
	"log"
	"os"
)

type hotel struct {
	Name    string
	Address string
	City    string
	Zip     int
	Region  string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	caliHotels := []hotel{
		hotel{
			Name:    "Marriott",
			Address: "1234 5th St",
			City:    "San Diego",
			Zip:     12345,
			Region:  "Southern",
		},
		hotel{
			Name:    "Bobs Hotel",
			Address: "5555 6th St",
			City:    "Sacramento",
			Zip:     54321,
			Region:  "Northern",
		},
	}

	err := tpl.Execute(os.Stdout, caliHotels)
	if err != nil {
		log.Fatalln(err)
	}
}
