package controllers

import (
	"AkiyaDeGo/app/models"
	"AkiyaDeGo/config"
	"fmt"
	"html/template"
	"net/http"
	"regexp"
	"strconv"
)

func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}
	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
}

func session(w http.ResponseWriter, r *http.Request) (sess models.Session, err error) {
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		sess = models.Session{UUID: cookie.Value}
		if ok, _ := sess.CheckSession(); !ok {
			err = fmt.Errorf("Invalid session")
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
		// Atoi(q[2])は二番目のココの部分 → /(edit|update|delete)/
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

	http.HandleFunc("/top/", handleTop)
	http.HandleFunc("/", handleMain)

	http.HandleFunc("/signup/", handleSignup)
	http.HandleFunc("/login/", handleLogin)
	http.HandleFunc("/authenticate", handleAuthenticate)
	http.HandleFunc("/index/", handleIndex)
	http.HandleFunc("/logout/", handleLogout)

	http.HandleFunc("/post/new/", postNew)
	http.HandleFunc("/post/save/", postSave)
	http.HandleFunc("/post/edit/", parseURL(postEdit))
	http.HandleFunc("/post/update/", parseURL(postUpdate))

	http.HandleFunc("/imageUpload/", handleUpload)
	http.HandleFunc("/imageShow/", handleShow)
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
