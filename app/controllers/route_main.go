package controllers

import (
	"AkiyaDeGo/app/models"
	// "fmt"
	"log"
	"net/http"
	// "os"
)

// type Post struct {
// 	Title      string
// 	Category   string
// 	Prefecture string
// 	Opinion    string
// }

func handleTop(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		generateHTML(w, "HELLO", "layout", "public_navbar", "top")
	} else {
		http.Redirect(w, r, "/index/", 302)
	}
}

func handleMain(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, nil, "layout", "public_navbar", "index")
}

// func handlePost(w http.ResponseWriter, r *http.Request) {
// 	generateHTML(w, nil, "layout", "public_navbar", "post")
// 	if r.Method == "POST" {
// 		post := Post{
// 			Title:      r.PostFormValue("タイトル"),
// 			Category:   r.PostFormValue("種類"),
// 			Prefecture: r.PostFormValue("都道府県"),
// 			Opinion:    r.PostFormValue("freeans"),
// 		}
// 		txt, err := os.OpenFile("a.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
// 		if err != nil {
// 			log.Fatalln(err)
// 		}
// 		fmt.Fprintln(txt, "タイトル：	"+post.Title)
// 		fmt.Fprintln(txt, "種類：　　　"+post.Category)
// 		fmt.Fprintln(txt, "都道府県：	"+post.Prefecture)
// 		fmt.Fprintln(txt, "body-start")
// 		fmt.Fprintln(txt, post.Opinion)
// 		fmt.Fprintln(txt, "body-end")
// 		fmt.Fprintln(txt, "")
// 	}
// }

// func handleIndex(w http.ResponseWriter, r *http.Request) {
// 	_, err := session(w, r)
// 	if err != nil {
// 		http.Redirect(w, r, "/top/", 302)
// 	} else {
// 		generateHTML(w, nil, "layout", "private_navbar", "index")
// 	}
// }

// func handleMytop(w http.ResponseWriter, r *http.Request) {
// 	_, err := session(w, r)
// 	if err != nil {
// 		http.Redirect(w, r, "/top/", 302)
// 	} else {
// 		generateHTML(w, nil, "layout", "private_navbar", "mytop")
// 	}
// }

func handleIndex(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/", 302)
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
		p.Description = r.PostFormValue("description")
		p.UserID = user.ID // 現段階では種類と都道府県はなし
		if err := p.CreatePost(); err != nil {
			log.Println(err)
		}

		http.Redirect(w, r, "/index/", 302)
	}
}
