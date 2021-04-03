package main

import (
	"AkiyaDeGo/app/controllers"
	"AkiyaDeGo/config"
	"fmt"
	"github.com/playree/goingtpl"
	"html/template"
	"log"
	"net/http"
)

func main() {
	fmt.Println(config.Config.Port)
	log.Println("test")

	goingtpl.SetBaseDir("./app/views/templates")

	http.HandleFunc("/top/", handleTop)
	http.HandleFunc("/", handleMain)
	http.HandleFunc("/post/", handlePost)
	controllers.StartMainServer()
}

func handleTop(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(goingtpl.ParseFile("top.html"))
	tpl.Execute(w, nil)
}

func handleMain(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(goingtpl.ParseFile("index.html"))
	tpl.Execute(w, nil)
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(goingtpl.ParseFile("post.html"))
	tpl.Execute(w, nil)
}
