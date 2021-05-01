package main

import (
	// "AkiyaDeGo/app/controllers"
	"AkiyaDeGo/app/models"
	// "AkiyaDeGo/config"
	"fmt"
	// "log"
)

func main() {
	// fmt.Println(config.Config.Port)
	// fmt.Println(models.Db)

	// u := &models.User{}
	// u.Name = "test2"
	// u.Email = "test2@example.com"
	// u.PassWord = "testtest"
	// fmt.Println(u)
	// u.CreateUser()

	// user, _ := models.GetUser(2)
	// p := &models.Post{}
	// p.Title = "First Title"
	// p.Description = "First Description"
	// p.UserID = user.ID
	// p.CreatePost()
	// p, _ := models.GetPost(1)
	// fmt.Println(p)
	// ーーーーーーーーーーGetPostsーーーーーーーーーー
	user, _ := models.GetUser(2)
	p := &models.Post{}
	p.Title = "Second Title"
	p.Description = "Second Description"
	p.UserID = user.ID
	p.CreatePost()
	posts, _ := models.GetPosts()
	for _, v := range posts {
		fmt.Println(v)
	}
	// ーーーーーーーーーーGetPostsーーーーーーーーーー
	// u, _ := models.GetUser(1)
	// fmt.Println(u)
	// u.Name = "test2"
	// u.Email = "test2@example.com"
	// u.UpdateUser()
	// u, _ = models.GetUser(1)
	// fmt.Println(u)
	// u.DeleteUser()
	// u, _ = models.GetUser(1)
	// fmt.Println(u)
	// user, _ := models.GetUserByEmail("test@example.com")
	// fmt.Println(user)

	// session, err := user.CreateSession()
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(session)
	// controllers.StartMainServer()
}
