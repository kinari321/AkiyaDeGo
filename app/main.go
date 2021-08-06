package main

import (
	"fmt"
	"github.com/kinari321/AkiyaDeGo/app/config"
	"github.com/kinari321/AkiyaDeGo/app/pkg/controllers"
	"github.com/kinari321/AkiyaDeGo/app/pkg/models"
)

func main() {
	fmt.Println(config.Config.Port)
	controllers.StartMainServer()
}
