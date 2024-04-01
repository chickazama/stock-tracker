package ui

import (
	"log"
	"net/http"
)

type IndexViewModel struct {
	Title string
}

type IndexHandler struct {
	tmplName string
}

func NewIndexHandler() *IndexHandler {
	ret := new(IndexHandler)
	ret.tmplName = "Index"
	return ret
}

func (h *IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.get(w, r)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *IndexHandler) get(w http.ResponseWriter, r *http.Request) {
	viewData := IndexViewModel{Title: "Home"}
	err := tmpl.ExecuteTemplate(w, h.tmplName, viewData)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}
