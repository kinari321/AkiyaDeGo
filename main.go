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

	http.HandleFunc("/", handleTest)
	// http.HandleFunc("/post", handleTest2)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleTest(w http.ResponseWriter, r *http.Request) {
	// parent.htmlをパース
	tpl := template.Must(goingtpl.ParseFile("layout.html"))
	tpl.Execute(w, nil)
}

/*
func handleTest2(w http.ResponseWriter, r *http.Request) {
    // parent.htmlをパース
    tpl := template.Must(goingtpl.ParseFile("index2.html"))
    tpl.Execute(w, nil)
}
*/
