package main

import (
	"AkiyaDeGo/config"
	"github.com/playree/goingtpl"
    "fmt"
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
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleTop(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(goingtpl.ParseFile("top.html"))
	tpl.Execute(w, nil)
}

func handleMain(w http.ResponseWriter, r *http.Request) {
    tpl := template.Must(goingtpl.ParseFile("index.html"))
    tpl.Execute(w, nil)
}
