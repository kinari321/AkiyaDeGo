package controllers

import (
	"fmt"
	"io"
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
		fmt.Fprintln(io.Writer(txt), post)
	}
}
func handleSignup(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, nil, "layout", "public_navbar", "signup")
}
func handleLogin(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, nil, "layout", "public_navbar", "login")
}
