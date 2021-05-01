package main

import (
	// "AkiyaDeGo/app/controllers"
	"AkiyaDeGo/app/models"
	// "AkiyaDeGo/config"
	"fmt"
	"log"
)

func main() {
	// fmt.Println(config.Config.Port)
	// fmt.Println(models.Db)
	// ーーーーーーーーーーGetUserByEmailーーーーーーーーーー
	user, _ := models.GetUserByEmail("test2@example.com")
	fmt.Println(user)
	// ーーーーーーーーーーGetUserByEmailーーーーーーーーーー

	// ーーーーーーーーーーCreateSessionーーーーーーーーーー
	session, err := user.CreateSession()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(session)
	// ーーーーーーーーーーCreateSessionーーーーーーーーーー

	// ーーーーーーーーーーCreateSessionーーーーーーーーーー
	valid, _ := session.CheckSession()
	fmt.Println(valid)
	// ーーーーーーーーーーCreateSessionーーーーーーーーーー

	// controllers.StartMainServer()

}
