package main

import (
	"log"
	"net/http"
	"regexp"

	"github.com/chickazama/stock-tracker/router"
	"github.com/chickazama/stock-tracker/ui"
)

func main() {
	mux := http.NewServeMux()
	serveStaticFiles(mux)
	r := router.New()
	r.AddHandler(regexp.MustCompile(`^/$`), ui.NewIndexHandler())
	mux.Handle("/", r)
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func serveStaticFiles(mux *http.ServeMux) {
	fsRoot := http.Dir("./static/")
	fs := http.FileServer(fsRoot)
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
}
