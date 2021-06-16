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
	controllers.StartMainServer()
}
