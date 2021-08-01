package main

import (
	"github.com/kinari321/AkiyaDeGo/app/pkg/controllers"
	"github.com/kinari321/AkiyaDeGo/app/pkg/models"
	"github.com/kinari321/AkiyaDeGo/app/config"
	"fmt"
)

func main() {
	fmt.Println(config.Config.Port)
	fmt.Println(models.Db)
	controllers.StartMainServer()
}
