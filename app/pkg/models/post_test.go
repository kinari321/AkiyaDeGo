package models

import (
	// "fmt"
	// "github.com/DATA-DOG/go-sqlmock"
	// "regexp"
	"testing"
)

var post = &Post{
	Title:      "タイトル",
	Category:   "空き家",
	Prefecture: "和歌山県",
	UserID:     1,
}

func TestCreatePost(t *testing.T) {
	defer setup()
	if err := users[0].CreateUser(); err != nil {
		t.Errorf("Cannot create user. err: %v", err)
	}
	u, err := GetUserByEmail(users[0].Email)
	if err != nil {
		t.Errorf("User not created. err: %v", err)
	}
	post.UserID = u.ID
	err = post.CreatePost()
	if err != nil {
		t.Error(err, "Cannot create post")
	}
}

func TestGetPost(t *testing.T) {
	t.Skip("Skipping GetPost test")
}

func TestGetPosts(t *testing.T) {
}
func TestGetPostsByUser(t *testing.T) {
}
func TestUpdatePost(t *testing.T) {
}
func TestDeletePost(t *testing.T) {
}
