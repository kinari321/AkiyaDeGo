package controllers

import (
	"AkiyaDeGo/config"
	"net/http"
)

func StartMainServer() error {
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
