package controllers

import (
	"AkiyaDeGo/config"
	"fmt"
	"html/template"
	// "log"
	"net/http"
)

func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	// log.Println("helloworld")
	// fmt.Println("helloworld")
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
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
