package models

import (
	"log"
	"time"
)

type Post struct {
	ID          int
	Title       string
	Description string
	UserID      int
	CreatedAt   time.Time
}

func (p *Post) CreatePost() (err error) {
	cmd := `INSERT INTO posts (
		title,
		description,
		user_id,
		created_at) VALUES (?, ?, ?, ?)`
	// p.UserIDどうなる？
	_, err = Db.Exec(cmd,
		p.Title,
		p.Description,
		p.UserID,
		time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetPost(id int) (post Post, err error) {
	post = Post{}
	cmd := `SELECT id , title, description, user_id, created_at FROM posts
	WHERE id = ?`
	err = Db.QueryRow(cmd, id).Scan(
		&post.ID,
		&post.Title,
		&post.Description,
		&post.UserID,
		&post.CreatedAt,
	)
	return post, err
}

func GetPosts() (posts []Post, err error) {
	cmd := `SELECT id , title, description, user_id, created_at FROM posts`
	rows, err := Db.Query(cmd)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var post Post
		err = rows.Scan(
			&post.ID,
			&post.Title,
			&post.Description,
			&post.UserID,
			&post.CreatedAt,
		)
		if err != nil {
			log.Fatalln(err)
		}
		posts = append(posts, post)
	}
	rows.Close()
	return posts, err
}
