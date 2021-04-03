package controllers

import (
	"AkiyaDeGo/config"
	"net/http"
)

func StartMainServer() error {
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))

	http.HandleFunc("/top/", handleTop)
	http.HandleFunc("/", handleMain)
	http.HandleFunc("/post/", handlePost)
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
