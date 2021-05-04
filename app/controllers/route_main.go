package controllers

import (
	"AkiyaDeGo/app/models"
	"log"
	"net/http"
)

func handleTop(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		generateHTML(w, nil, "layout", "public_navbar", "top")
	} else {
		http.Redirect(w, r, "/index/", 302)
	}
}

func handleMain(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, nil, "layout", "public_navbar", "main")
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/top/", 302)
	} else {
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		posts, _ := user.GetPostsByUser()
		user.Posts = posts
		generateHTML(w, user, "layout", "private_navbar", "index")
	}
}

func postNew(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login/", 302)
	} else {
		generateHTML(w, nil, "layout", "private_navbar", "post_new")
	}
}

func postSave(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login/", 302)
	} else {
		err = r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}

		p := &models.Post{}
		p.Title = r.PostFormValue("title")
		p.Type = r.PostFormValue("type")
		p.Prefecture = r.PostFormValue("prefecture")
		p.Description = r.PostFormValue("description")
		p.UserID = user.ID
		if err := p.CreatePost(); err != nil {
			log.Println(err)
		}

		http.Redirect(w, r, "/index/", 302)
	}
}
