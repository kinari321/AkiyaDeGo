package controllers

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/kinari321/AkiyaDeGo/app/errors"
	"github.com/kinari321/AkiyaDeGo/app/pkg/models"
)

func handleTop(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		posts, err := models.GetPosts()
		if err != nil {
			log.Printf("get posts failed: %+v\n", errors.StackTrace(err))
		}
		generateHTML(w, posts, "layout", "public_navbar", "top")
	} else {
		http.Redirect(w, r, "/index", 302)
	}
}

func handleMain(w http.ResponseWriter, r *http.Request) {
	var files []string
	files = append(files, "pkg/views/templates/layout-main.html")
	files = append(files, "pkg/views/templates/public_navbar.html")
	files = append(files, "pkg/views/templates/main.html")
	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout-main", nil)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/top", 302)
	} else {
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Printf("get user by session failed: %+v\n", errors.StackTrace(err))
		}
		posts, err := user.GetPostsByUser()
		if err != nil {
			log.Printf("get posts by user failed: %+v\n", errors.StackTrace(err))
		}
		user.Posts = posts
		generateHTML(w, user, "layout", "private_navbar", "index")
	}
}

func postNew(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		generateHTML(w, nil, "layout", "private_navbar", "post_new")
	}
}

func postSave(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		err = r.ParseForm()
		if err != nil {
			log.Printf("parse form failed: %+v\n", errors.StackTrace(err))
		}
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Printf("get user by session failed: %+v\n", errors.StackTrace(err))
			http.Redirect(w, r, "/post/new", 302)
		}
		err = r.ParseMultipartForm(32 << 20)
		if err != nil {
			log.Printf("parse form failed: %+v\n", errors.StackTrace(err))
			http.Redirect(w, r, "/post/new", 302)
		} else {

			file, fileHeader, err := r.FormFile("image")
			if err != nil {
				log.Printf("form file failed: %+v\n", errors.StackTrace(err))
				http.Redirect(w, r, "/post/new", 302)
			} else {
				defer file.Close()

				uploadedFileName := fileHeader.Filename
				log.Printf("uploadedFileName: %v\n", uploadedFileName)
				path := "/var/www/image/" + uploadedFileName // ローカル用
				// path := "/usr/share/nginx/html/media/" + uploadedFileName // EC2用
				f, err := os.Create(path)
				if err != nil {
					log.Printf("create file failed: %+v\n", errors.StackTrace(err))
					http.Redirect(w, r, "/post/new", 302)
				}
				defer f.Close()
				io.Copy(f, file)

				p := &models.Post{}
				p.ImagePath = path
				p.Title = r.PostFormValue("title")
				p.Category = r.PostFormValue("category")
				p.Prefecture = r.PostFormValue("prefecture")
				p.Description = r.PostFormValue("description")
				p.UserID = user.ID
				if err := p.CreatePost(); err != nil {
					log.Printf("create post failed: %+v\n", errors.StackTrace(err))
					http.Redirect(w, r, "/post/new", 302)
				}
				http.Redirect(w, r, "/index", 302)
			}
		}
	}
}

func postEdit(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		_, err := sess.GetUserBySession()
		if err != nil {
			log.Printf("get user by session failed: %+v\n", errors.StackTrace(err))
		}
		p, err := models.GetPost(id)
		if err != nil {
			log.Printf("get post failed: %+v\n", errors.StackTrace(err))
		}
		generateHTML(w, p, "layout", "private_navbar", "post_edit")
	}
}

func postUpdate(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		err = r.ParseForm()
		if err != nil {
			log.Printf("parse form failed: %+v\n", errors.StackTrace(err))
		}
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Printf("get user by session failed: %+v\n", errors.StackTrace(err))
		}
		p, err := models.GetPost(id)
		title := r.PostFormValue("title")
		description := r.PostFormValue("description")
		post := &models.Post{
			ID:          id,
			ImagePath:   p.ImagePath,
			Title:       title,
			Category:    p.Category,
			Prefecture:  p.Prefecture,
			Description: description,
			UserID:      user.ID}
		if err := post.UpdatePost(); err != nil {
			log.Printf("update post failed: %+v\n", errors.StackTrace(err))
		}
		http.Redirect(w, r, "/index", 302)
	}
}

func postDelete(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		_, err := sess.GetUserBySession()
		if err != nil {
			log.Printf("get user by session failed: %+v\n", errors.StackTrace(err))
		}
		t, err := models.GetPost(id)
		if err != nil {
			log.Printf("get post failed: %+v\n", errors.StackTrace(err))
		}
		if err := t.DeletePost(); err != nil {
			log.Printf("delete post failed: %+v\n", errors.StackTrace(err))
		}
		http.Redirect(w, r, "/index", 302)
	}
}
