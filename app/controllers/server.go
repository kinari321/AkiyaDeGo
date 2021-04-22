package controllers

import (
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

func StartMainServer() error {
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))

	http.HandleFunc("/top/", handleTop)
	http.HandleFunc("/", handleMain)
	http.HandleFunc("/post/", handlePost)
	http.HandleFunc("/signup/", handleSignup)
	http.HandleFunc("/login/", handleLogin)
	http.HandleFunc("/upload", handleUpload)
	http.HandleFunc("/upload/show", handleShow)
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
