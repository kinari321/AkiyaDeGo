package controllers

import (
	"AkiyaDeGo/app/models"
	"AkiyaDeGo/config"
	"fmt"
	"html/template"
	"net/http"
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

	http.HandleFunc("/upload/", handleUpload)
	http.HandleFunc("/show/", handleShow)
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
