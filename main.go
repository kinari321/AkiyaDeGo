package main

import (
	// "AkiyaDeGo/app/controllers"
	"AkiyaDeGo/app/models"
	"AkiyaDeGo/config"
	"fmt"
	// "log"
)

func main() {
	fmt.Println(config.Config.Port)
	fmt.Println(models.Db)

	// ーーーーーーーーーーCreateUserーーーーーーーーーー
	// u := &models.User{}
	// u.Name = "test3"
	// u.Email = "test3@example.com"
	// u.PassWord = "testtest"
	// fmt.Println(u)
	// u.CreateUser()
	// ーーーーーーーーーーCreateUserーーーーーーーーーー

	// ーーーーーーーーーーCreatePostーーーーーーーーーー
	// user, _ := models.GetUser(3)
	// p := &models.Post{}
	// p.Title = "5 Title"
	// p.Description = "5 Description"
	// p.UserID = user.ID
	// p.CreatePost()
	// fmt.Println(p)
	// ーーーーーーーーーーCreatePostーーーーーーーーーー

	// ーーーーーーーーーーGetPostsーーーーーーーーーー
	// user, _ := models.GetUser(2)
	// p := &models.Post{}
	// p.Title = "Second Title"
	// p.Description = "Second Description"
	// p.UserID = user.ID
	// p.CreatePost()
	// posts, _ := models.GetPosts()
	// for _, v := range posts {
	// 	fmt.Println(v)
	// }
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

	// ーーーーーーーーーーGetPostsByUserーーーーーーーーーー
	// user2, _ := models.GetUser(2)
	// posts, _ := user2.GetPostsByUser()
	// for _, v := range posts {
	// 	fmt.Println(v)
	// }
	// ーーーーーーーーーーGetPostsByUserーーーーーーーーーー

	// ーーーーーーーーーーUpdatePostーーーーーーーーーー
	// p, _ := models.GetPost(1)
	// p.Title = "Update Title"
	// p.Description = "Update Description"
	// p.UpdatePost()
	// ーーーーーーーーーーUpdatePostーーーーーーーーーー

	// ーーーーーーーーーーDeletePostーーーーーーーーーー
	// p, _ := models.GetPost(3)
	// p.DeletePost()
	// ーーーーーーーーーーDeletePostーーーーーーーーーー

	// session, err := user.CreateSession()
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(session)
	// controllers.StartMainServer()
}
