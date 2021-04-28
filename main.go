package main

import (
	"AkiyaDeGo/app/controllers"
	"AkiyaDeGo/app/models"
	"AkiyaDeGo/config"
	"fmt"
)

func main() {
	fmt.Println(config.Config.Port)
	fmt.Println(models.Db)

	// // u := &models.User{}
	// // u.Name = "test"
	// // u.Email = "test@example.com"
	// // u.PassWord = "testtest"
	// // fmt.Println(u)
	// u.CreateUser()
	u, _ := models.GetUser(1)
	fmt.Println(u)

	u.Name = "test2"
	u.Email = "test2@example.com"
	u.UpdateUser()
	u, _ = models.GetUser(1)
	controllers.StartMainServer()
}
