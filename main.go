package main

import (
	"AkiyaDeGo/app/controllers"
	"AkiyaDeGo/app/models"
	"AkiyaDeGo/config"
	"fmt"
	// "log"
)

func main() {
	fmt.Println(config.Config.Port)
	fmt.Println(models.Db)

	controllers.StartMainServer()

	// user, _ := models.GetUserByEmail("test@example.com")
	// fmt.Println(user)
	// u, _ := models.GetUser(1)
	// fmt.Println(u)

	// u.Email = "test@example.com"
	// u.UpdateUser()
	// u, _ = models.GetUser(1)
	// user := u.models{}
	// user.Email = "test@example.com"
	// user.UpdateUser()
}
