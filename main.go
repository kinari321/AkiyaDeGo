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

	// p, _ := models.GetPost(5)
	// fmt.Println(p)
	controllers.StartMainServer()
	// p := &models.Post{}
	// p.ImagePath = "/var/www/test/akiya2.jepg"
	// p.Title = "タイトル2アップデート"
	// p.Type = "空き家"
	// p.Prefecture = "東京都"
	// p.Description = "詳細説明2"
	// p.UserID = 1
	// fmt.Println(p)
	// p.CreatePost()
	// u := &models.User{}
	// u.Name = "test"
	// u.Email = "test@example.com"
	// u.PassWord = "test"
	// u.CreateUser()
	// user, _ := models.GetUser(1)
	// posts, _ := user.GetPostsByUser()
	// for _, v := range posts {
	// 	fmt.Println(v)
	// }
	// user, _ := models.GetUserByEmail("test@example.com")
	// fmt.Println(user)
	// session, err := user.CreateSession()
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(session)
	// valid, _ := session.CheckSession()
	// fmt.Println(valid)
	// u, _ := models.GetUser(1)
	// fmt.Println(u)

	// u.Email = "test@example.com"
	// u.UpdateUser()
	// u, _ = models.GetUser(1)
	// user := u.models{}
	// user.Email = "test@example.com"
	// user.UpdateUser()
}
