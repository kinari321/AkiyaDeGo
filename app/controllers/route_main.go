package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type Post struct {
	Title      string
	Category   string
	Prefecture string
	Opinion    string
}

func handleTop(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, nil, "layout", "public_navbar", "top")
}
func handleMain(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, nil, "layout", "public_navbar", "index")
}
func handlePost(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, nil, "layout", "public_navbar", "post")
	if r.Method == "POST" {
		post := Post{
			Title:      r.PostFormValue("タイトル"),
			Category:   r.PostFormValue("種類"),
			Prefecture: r.PostFormValue("都道府県"),
			Opinion:    r.PostFormValue("freeans"),
		}
		txt, err := os.OpenFile("a.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Fprintln(txt, "タイトル：	"+post.Title)
		fmt.Fprintln(txt, "種類：　　　"+post.Category)
		fmt.Fprintln(txt, "都道府県：	"+post.Prefecture)
		fmt.Fprintln(txt, "body-start")
		fmt.Fprintln(txt, post.Opinion)
		fmt.Fprintln(txt, "body-end")
		fmt.Fprintln(txt, "")
	}
}
