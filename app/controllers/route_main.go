package controllers

import (
	"html/template"
	"log"
	"net/http"

	"github.com/playree/goingtpl"
)

var err error

func handleTop(w http.ResponseWriter, r *http.Request) {
	goingtpl.SetBaseDir("./app/views/templates")
	tpl := template.Must(goingtpl.ParseFile("top.html"))
	if err != nil {
		log.Fatalln(err)
	}
	tpl.Execute(w, nil)
}

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
