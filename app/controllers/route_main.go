package controllers

import (
	"fmt"
	"net/http"
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
		fmt.Println(post)
	}
}
func handleSignup(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, nil, "layout", "public_navbar", "signup")
}
func handleLogin(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, nil, "layout", "public_navbar", "login")
}
