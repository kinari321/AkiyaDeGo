package models

import (
	"testing"
)

var post = &Post{
	Title:      "タイトル",
	Category:   "空き家",
	Prefecture: "和歌山県",
	UserID:     1,
}

func TestCreatePost(t *testing.T) {
	t.Skip("Skipping GetPost test")
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
