package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"regexp"
	"strconv"

	"github.com/kinari321/AkiyaDeGo/app/config"
	"github.com/kinari321/AkiyaDeGo/app/pkg/models"
)

func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("pkg/views/templates/%s.html", file))
	}
	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
}

func session(w http.ResponseWriter, r *http.Request) (sess models.Session, err error) {
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		sess = models.Session{UUID: cookie.Value}
		if ok, _ := sess.CheckSession(); !ok {
			return sess, err
		}
	}

	return sess, err
}

var validPath = regexp.MustCompile("^/post/(edit|update|delete)/([0-9]+)$")

func parseURL(fn func(http.ResponseWriter, *http.Request, int)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		q := validPath.FindStringSubmatch(r.URL.Path)
		if q == nil {
			http.NotFound(w, r)
			return
		}
		qi, err := strconv.Atoi(q[2])
		if err != nil {
			http.NotFound(w, r)
			return
		}

		fn(w, r, qi)
	}
}

func StartMainServer() error {
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))

	http.HandleFunc("/top", handleTop)
	http.HandleFunc("/", handleMain)

	http.HandleFunc("/signup", handleSignup)
	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/authenticate", handleAuthenticate)
	http.HandleFunc("/index", handleIndex)
	http.HandleFunc("/logout", handleLogout)

	http.HandleFunc("/post/new", postNew)
	http.HandleFunc("/post/save", postSave)
	http.HandleFunc("/post/edit/", parseURL(postEdit))
	http.HandleFunc("/post/update/", parseURL(postUpdate))
	http.HandleFunc("/post/delete/", parseURL(postDelete))

	return http.ListenAndServe(":8080", nil)
}
