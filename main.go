package main

import (
	"AkiyaDeGo/app/controllers"
	"AkiyaDeGo/config"
	"fmt"
	// "log"
)

func main() {
	fmt.Println(config.Config.Port)
	// log.Println("test")

	controllers.StartMainServer()
}
