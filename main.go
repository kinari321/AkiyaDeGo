package main

import (
    "html/template"
    "log"
   	"net/http"

    "github.com/playree/goingtpl"
)

func main() {
    // テンプレートのディレクトリを設定
    goingtpl.SetBaseDir("./app/views/templates")

    http.HandleFunc("/", handleTest)
    http.HandleFunc("/post", handleTest2)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleTest(w http.ResponseWriter, r *http.Request) {
    // parent.htmlをパース
    tpl := template.Must(goingtpl.ParseFile("index.html"))
    tpl.Execute(w, nil)
}

func handleTest2(w http.ResponseWriter, r *http.Request) {
    // parent.htmlをパース
    tpl := template.Must(goingtpl.ParseFile("index2.html"))
    tpl.Execute(w, nil)
}