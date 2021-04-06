package controllers

import (
	"net/http"
)

func handleTop(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, nil, "layout", "top")
}
func handleMain(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, nil, "layout", "index")
}
func handlePost(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, nil, "layout", "post")
}
func handleSignup(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, nil, "layout", "signup")
}
