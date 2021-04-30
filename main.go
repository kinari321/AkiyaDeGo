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

	// u := &models.User{}
	// u.Name = "test2"
	// u.Email = "test2@example.com"
	// u.PassWord = "testtest"
	// fmt.Println(u)
	// u.CreateUser()
	// u, _ := models.GetUser()
	// fmt.Println(u)

	// u.Name = "test2"
	// u.Email = "test2@example.com"
	// u.UpdateUser()
	// u, _ = models.GetUser(1)
	// fmt.Println(u)

	// u.DeleteUser()
	// user, _ := models.GetUser(4)

	// user, _ := models.GetUser(1)
	// p := &models.Post{}
	// p.Title = "3 Post"
	// p.Description = "3 Description"
	// p.UserID = user.ID
	// fmt.Println(p)
	// p.CreatePost()

	// posts, _ := models.GetPosts()
	// for _, v := range posts {
	// 	fmt.Println(v)
	// }

	// p, _ := models.GetPost(3)
	// p.Title = "Update Title"
	// p.Description = "Update Description"
	// p.UpdatePost()
	// p, _ := models.GetPost(3)
	// p.DeletePost()

	controllers.StartMainServer()
}
