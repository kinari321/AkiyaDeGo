package models

import (
	// "fmt"
	// "github.com/DATA-DOG/go-sqlmock"
	// "regexp"
	"testing"
)

func TestCreatePost(t *testing.T) {
	t.Skip("Skipping init test")
}

func TestGetPost(t *testing.T) {
	// db, mock, err := sqlmock.New()
	// if err != nil {
	// 	t.Fatalf("failed to init db mock.")
	// }
	// defer db.Close()
	// id := 1
	// imagepath := "/var/www/image/akiya.jpeg"
	// title := "空き家投稿"
	// category := "空き家"
	// prefecture := "北海道"
	// description := "空き家の詳細説明"
	// userId := 1
	// createdAt := "2021-05-22 07:58:32 +0000 UTC"
	// mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "posts" WHERE (id = $1)`)).
	// 	WithArgs(id).
	// 	WillReturnRows(sqlmock.NewRows([]string{"id", "imagepath", "title", "category", "prefecture", "description", "user_id", "created_at"}).
	// 		AddRow(id, imagepath, title, category, prefecture, description, userId, createdAt))
	// post, err := GetPost(id)
	// if err != nil {
	// 	t.Fatalf("failed to get post: %s", err)
	// }
	// fmt.Printf("postは、%v\n", post)
	// if err := mock.ExpectationsWereMet(); err != nil {
	// 	t.Fatalf("failed to ExpectationWerMet(): %s", err)
	// }
	t.Skip("Skipping GetPost test")
}

func TestGetPosts(t *testing.T) {
	t.Skip("Skipping GetPosts test")
}
func TestGetPostsByUser(t *testing.T) {
	t.Skip("Skipping GetPostsByUser test")
}
func TestUpdatePost(t *testing.T) {
	t.Skip("Skipping UpdatePost test")
}
func TestDeletePost(t *testing.T) {
	t.Skip("Skipping DeletePost test")
}
