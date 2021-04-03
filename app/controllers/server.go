package controllers

import (
	"AkiyaDeGo/config"
	"net/http"
)

func StartMainServer() error {
	http.HandleFunc("/top/", handleTop)
	http.HandleFunc("/", handleMain)
	http.HandleFunc("/post/", handlePost)
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
