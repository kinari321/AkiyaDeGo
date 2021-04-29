package main

import (
	"AkiyaDeGo/app/controllers"
	"AkiyaDeGo/app/models"
	// "AkiyaDeGo/config"
	"fmt"
)

func main() {
	// fmt.Println(config.Config.Port)
	// fmt.Println(models.Db)

	// u := &models.User{}
	// u.Name = "test4"
	// u.Email = "test4@example.com"
	// u.PassWord = "testtest"
	// fmt.Println(u)
	// u.CreateUser()
	// u, _ := models.GetUser(2)
	// fmt.Println(u)

	// u.Name = "test2"
	// u.Email = "test2@example.com"
	// u.UpdateUser()
	// u, _ = models.GetUser(1)
	// fmt.Println(u)

	// u.DeleteUser()
	// user, _ := models.GetUser(4)
	// p := &models.Post{}
	// p.Title = "4 Post"
	// p.Description = "4 Description"
	// p.UserID = user.ID // ここはどうする？
	// fmt.Println(p)
	// p.CreatePost()

	posts, _ := models.GetPosts()
	for _, v := range posts {
		fmt.Println(v)
	}
	controllers.StartMainServer()
}
