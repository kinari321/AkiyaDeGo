package models

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"
	_ "image/png"
	"os"
	"time"

	"github.com/kinari321/AkiyaDeGo/app/errors"
	"github.com/nfnt/resize"
)

type Post struct {
	ID          int
	ImagePath   string
	Title       string
	Category    string
	Prefecture  string
	Description string
	UserID      int
	CreatedAt   time.Time
}

func (p *Post) CreatePost() (err error) {
	cmd := `INSERT INTO posts (
		imagepath,
		title,
		category,
		prefecture,
		description,
		user_id,
		created_at) VALUES (?, ?, ?, ?, ?, ?, ?)`
	_, err = Db.Exec(cmd,
		p.ImagePath,
		p.Title,
		p.Category,
		p.Prefecture,
		p.Description,
		p.UserID,
		time.Now())
	if err != nil {
		return errors.SetError(errors.ErrDataBase, fmt.Sprintf("create post failed: %s", err))
	}
	return err
}

func GetPost(id int) (post Post, err error) {
	cmd := `SELECT * FROM posts
	WHERE id = ?`
	post = Post{}
	err = Db.QueryRow(cmd, id).Scan(
		&post.ID,
		&post.ImagePath,
		&post.Title,
		&post.Category,
		&post.Prefecture,
		&post.Description,
		&post.UserID,
		&post.CreatedAt,
	)
	return post, err
}

func GetPosts() (posts []Post, err error) {
	cmd := `SELECT id, imagepath, title, category, prefecture, description, user_id, created_at FROM posts`
	rows, err := Db.Query(cmd)
	if err != nil {
		return nil, errors.SetError(errors.ErrDataBase, fmt.Sprintf("get posts failed: %s", err))
	}
	for rows.Next() {
		var post Post
		err = rows.Scan(
			&post.ID,
			&post.ImagePath,
			&post.Title,
			&post.Category,
			&post.Prefecture,
			&post.Description,
			&post.UserID,
			&post.CreatedAt,
		)
		if err != nil {
			return nil, errors.SetError(errors.ErrDataBase, fmt.Sprintf("get posts failed: %s", err))
		}
		{
			file, err := os.Open(post.ImagePath)
			defer file.Close()
			if err != nil {
				return nil, errors.SetError(errors.ErrPath, fmt.Sprintf("open file failed: %s", err))
			}
			decodeImage, _, err := image.Decode(file)
			if err != nil {
				return nil, errors.SetError(errors.ErrImage, fmt.Sprintf("decode image failed: %s", err))
			}
			m := resize.Resize(400, 0, decodeImage, resize.Lanczos3)
			file, err = os.Open(post.ImagePath)
			if err != nil {
				return nil, errors.SetError(errors.ErrPath, fmt.Sprintf("open file failed: %s", err))
			}
			defer file.Close()
			buffer := new(bytes.Buffer)
			if err := jpeg.Encode(buffer, m, nil); err != nil {
				return nil, errors.SetError(errors.ErrImage, fmt.Sprintf("encode image failed: %s", err))
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
	cmd := `SELECT id, imagepath, title, category, prefecture, description, user_id, created_at FROM posts
	WHERE user_id = ?`
	rows, err := Db.Query(cmd, u.ID)
	if err != nil {
		return nil, errors.SetError(errors.ErrDataBase, fmt.Sprintf("create uuid failed: %s", err))
	}
	for rows.Next() {
		var post Post
		err = rows.Scan(
			&post.ID,
			&post.ImagePath,
			&post.Title,
			&post.Category,
			&post.Prefecture,
			&post.Description,
			&post.UserID,
			&post.CreatedAt,
		)
		if err != nil {
			return nil, errors.SetError(errors.ErrDataBase, fmt.Sprintf("create uuid failed: %s", err))
		}
		{
			file, err := os.Open(post.ImagePath)
			if err != nil {
				return nil, errors.SetError(errors.ErrDataBase, fmt.Sprintf("create uuid failed: %s", err))
			}
			defer file.Close()
			decodeImage, _, err := image.Decode(file)
			if err != nil {
				return nil, errors.SetError(errors.ErrDataBase, fmt.Sprintf("create uuid failed: %s", err))
			}
			m := resize.Resize(400, 0, decodeImage, resize.Lanczos3)
			buffer := new(bytes.Buffer)
			if err := jpeg.Encode(buffer, m, nil); err != nil {
				return nil, errors.SetError(errors.ErrDataBase, fmt.Sprintf("create uuid failed: %s", err))
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
	cmd := `UPDATE posts SET imagepath = ?, title = ?, category = ?, prefecture = ?,
		description = ?, user_id = ? WHERE id = ?`
	_, err = Db.Exec(cmd, p.ImagePath, p.Title, p.Category, p.Prefecture, p.Description, p.UserID, p.ID)
	if err != nil {
		return errors.SetError(errors.ErrDataBase, fmt.Sprintf("update post failed: %s", err))
	}
	return err
}

func (p *Post) DeletePost() (err error) {
	cmd := `DELETE FROM posts WHERE id = ?`
	_, err = Db.Exec(cmd, p.ID)
	if err != nil {
		return errors.SetError(errors.ErrDataBase, fmt.Sprintf("delete post failed: %s", err))
	}
	return err
}
