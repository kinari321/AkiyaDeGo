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

}
