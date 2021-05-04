package models

import (
	"log"
	"time"
)

type Post struct {
	ID          int
	Title       string
	Type        string
	Prefecture  string
	Description string
	UserID      int
	CreatedAt   time.Time
}

func (p *Post) CreatePost() (err error) {
	cmd := `INSERT INTO posts (
		title,
		type,
		prefecture,
		description,
		user_id,
		created_at) VALUES (?, ?, ?, ?, ?, ?)`
	_, err = Db.Exec(cmd,
		p.Title,
		p.Type,
		p.Prefecture,
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
	cmd := `SELECT id, title, type, prefecture, description, user_id, created_at FROM posts
	WHERE id = ?`
	err = Db.QueryRow(cmd, id).Scan(
		&post.ID,
		&post.Title,
		&post.Type,
		&post.Prefecture,
		&post.Description,
		&post.UserID,
		&post.CreatedAt,
	)
	return post, err
}

func GetPosts() (posts []Post, err error) {
	cmd := `SELECT id, title, type, prefecture, description, user_id, created_at FROM posts`
	rows, err := Db.Query(cmd)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var post Post
		err = rows.Scan(
			&post.ID,
			&post.Title,
			&post.Type,
			&post.Prefecture,
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

func (u *User) GetPostsByUser() (posts []Post, err error) {
	cmd := `SELECT id, title, type, prefecture, description, user_id, created_at FROM posts
	WHERE user_id = ?`
	rows, err := Db.Query(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var post Post
		err = rows.Scan(
			&post.ID,
			&post.Title,
			&post.Type,
			&post.Prefecture,
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

func (p *Post) UpdatePost() (err error) {
	cmd := `UPDATE posts SET title = ?, type = ?, prefecture = ?, description = ?, WHERE id = ?`
	_, err = Db.Exec(cmd, p.Title, p.Type, p.Prefecture, p.Description, p.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (p *Post) DeletePost() (err error) {
	cmd := `DELETE FROM posts WHERE id = ?`
	_, err = Db.Exec(cmd, p.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
