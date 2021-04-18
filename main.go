package main

import (
	"AkiyaDeGo/app/controllers"
	"AkiyaDeGo/config"
	"fmt"
)

func main() {
	fmt.Println(config.Config.Port)

	controllers.StartMainServer()
}
