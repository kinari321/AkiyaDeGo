package models

import (
	"bytes"
	"encoding/base64"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	_ "image/png"
	"log"
	"os"
	"time"
)

type Post struct {
	ID          int
	ImagePath   string
	Title       string
	Type        string
	Prefecture  string
	Description string
	UserID      int
	CreatedAt   time.Time
}

func (p *Post) CreatePost() (err error) {
	cmd := `INSERT INTO posts (
		imagepath,
		title,
		type,
		prefecture,
		description,
		user_id,
		created_at) VALUES (?, ?, ?, ?, ?, ?, ?)`
	_, err = Db.Exec(cmd,
		p.ImagePath,
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
	cmd := `SELECT id, imagepath, title, type, prefecture, description, user_id, created_at FROM posts
	WHERE id = ?`
	post = Post{}
	err = Db.QueryRow(cmd, id).Scan(
		&post.ID,
		&post.ImagePath,
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
	cmd := `SELECT id, imagepath, title, type, prefecture, description, user_id, created_at FROM posts`
	rows, err := Db.Query(cmd)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var post Post
		err = rows.Scan(
			&post.ID,
			&post.ImagePath,
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
		{
			file, _ := os.Open(post.ImagePath)
			defer file.Close()
			decodeImage, _, _ := image.Decode(file)
			m := resize.Resize(400, 0, decodeImage, resize.Lanczos3)
			file, _ = os.Open(post.ImagePath)
			defer file.Close()
			buffer := new(bytes.Buffer)
			if err := jpeg.Encode(buffer, m, nil); err != nil {
				log.Fatalln("Unable to encode image.")
			}
			str := base64.StdEncoding.EncodeToString(buffer.Bytes())
			post.ImagePath = str
		}
		posts = append(posts, post)
	}
	rows.Close()
	return posts, err
}

func (u *User) GetPostsByUser() (posts []Post, err error) {
	cmd := `SELECT id, imagepath, title, type, prefecture, description, user_id, created_at FROM posts
	WHERE user_id = ?`
	rows, err := Db.Query(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var post Post
		err = rows.Scan(
			&post.ID,
			&post.ImagePath,
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
		{
			file, _ := os.Open(post.ImagePath)
			defer file.Close()
			decodeImage, _, _ := image.Decode(file)
			m := resize.Resize(400, 0, decodeImage, resize.Lanczos3)
			file, _ = os.Open(post.ImagePath)
			defer file.Close()
			buffer := new(bytes.Buffer)
			if err := jpeg.Encode(buffer, m, nil); err != nil {
				log.Fatalln("Unable to encode image.")
			}
			str := base64.StdEncoding.EncodeToString(buffer.Bytes())
			post.ImagePath = str
		}
		posts = append(posts, post)
	}
	rows.Close()
	return posts, err
}

// Updateは微妙！！！
func (p *Post) UpdatePost() (err error) {
	cmd := `UPDATE posts SET imagepath = ?, title = ?, type = ?, prefecture = ?,
		description = ?, user_id = ? WHERE id = ?`
	_, err = Db.Exec(cmd, p.ImagePath, p.Title, p.Type, p.Prefecture, p.Description, p.UserID, p.ID)
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
