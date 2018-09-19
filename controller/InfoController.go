package controller

import (
	"net/http"
	"text/template"
)

type Page struct {
	Title string
	Count int
}
 
func InfoController(w http.ResponseWriter, req *http.Request){
	// handleは基本的に実行される
	page := Page{"Hello", 1}
	pageString := "resouces/index.tmpl.html"
	tmpl, err := template.ParseFiles(pageString)
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(w, page)
	if err != nil {
		panic(err)
	}
}