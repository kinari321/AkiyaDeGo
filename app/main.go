package main

import (
	_ "github.com/kinari321/AkiyaDeGo/app/config"
	"github.com/kinari321/AkiyaDeGo/app/pkg/controllers"
	_ "github.com/kinari321/AkiyaDeGo/app/pkg/models"
)

func main() {
	controllers.StartMainServer()
}
