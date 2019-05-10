package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.gmao"))
}

func main() {
	// err := template.ParseFiles("tpl.gohtml")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	err := tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	//add additional templates
	// err = template.ParseFiles("two.gmao", "vespa.gmao")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	//execute specific template when multiple in container
	err = tpl.ExecuteTemplate(os.Stdout, "vespa.gmao", nil)
	if err != nil {
		log.Fatalln(err)
	}

	glob()
}

func glob() {
	// err := template.ParseGlob("*.gmao")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	err := tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "vespa.gmao", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "two.gmao", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
