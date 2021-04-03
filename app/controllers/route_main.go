package controllers

import (
	"github.com/playree/goingtpl"
	"html/template"
	"log"
	"net/http"
)

func handleTop(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, nil, "layout", "top")
}

var err error

/*
func handleTop(w http.ResponseWriter, r *http.Request) {
	goingtpl.SetBaseDir("./app/views/templates")
	tpl := template.Must(goingtpl.ParseFile("top.html"))
	if err != nil {
		log.Fatalln(err)
	}
	tpl.Execute(w, nil)
}
*/

func handleMain(w http.ResponseWriter, r *http.Request) {
	goingtpl.SetBaseDir("./app/views/templates")
	tpl := template.Must(goingtpl.ParseFile("index.html"))
	if err != nil {
		log.Fatalln(err)
	}
	tpl.Execute(w, nil)
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	goingtpl.SetBaseDir("./app/views/templates")
	tpl := template.Must(goingtpl.ParseFile("post.html"))
	if err != nil {
		log.Fatalln(err)
	}
	tpl.Execute(w, nil)
}
