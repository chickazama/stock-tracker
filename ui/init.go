package ui

import (
	"html/template"
	"log"
)

const (
	tmplGlobPath = "./templates/*.go.html"
)

var (
	tmpl *template.Template
)

func init() {
	var err error
	tmpl, err = template.ParseGlob(tmplGlobPath)
	if err != nil {
		log.Fatal(err.Error())
	}
}
