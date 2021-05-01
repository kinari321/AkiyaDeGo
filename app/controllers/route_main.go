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
	_, err := session(w, r)
	if err != nil {
		generateHTML(w, "Hello", "layout", "public_navbar", "top")
	} else {
		http.Redirect(w, r, "/mytop/", 302)
	}
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

func index(w http.ResponseWriter, r *http.Request) {
	// ログインしているかどうかの判定を取得する
	_, err := session(w, r)
	// errがある場合は/top/リダイレクトにさせる
	if err != nil {
		http.Redirect(w, r, "/top/", 302)
	} else {
		// セッションがある場合はindex表示させる。を
		// user, err := sess.GetUserBySession()
		// if err != nil {
		// 	log.Println(err)
		// }
		// todos, _ := user.GetTodosByUser()
		// user.Todos = todos
		generateHTML(w, nil, "layout", "private_navbar", "index")
	}
}
