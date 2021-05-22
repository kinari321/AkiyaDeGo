package controllers

import (
	"AkiyaDeGo/app/models"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func handleTop(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		generateHTML(w, nil, "layout", "public_navbar", "top")
	} else {
		http.Redirect(w, r, "/index", 302)
	}
}

func handleMain(w http.ResponseWriter, r *http.Request) {
	allPosts, _ := models.GetPosts()
	generateHTML(w, allPosts, "layout", "public_navbar", "main")
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/top/", 302)
	} else {
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		posts, _ := user.GetPostsByUser()
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
			log.Println(err)
		}
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		err = r.ParseMultipartForm(32 << 20)
		if err != nil {
			log.Println(err)
		}

		file, fileHeader, err := r.FormFile("image")
		if err != nil {
			log.Println(err)
		}
		defer file.Close()

		uploadedFileName := fileHeader.Filename
		// uploadedFileName = CreateUUID() // ーーーーーーーーーーCreateUUID()をimportできなかったーーーーーーーーーー
		fmt.Println(uploadedFileName)
		path := "/var/www/image/" + uploadedFileName
		f, err := os.Create(path)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()
		io.Copy(f, file)

		p := &models.Post{}
		p.ImagePath = path
		p.Title = r.PostFormValue("title")
		p.Type = r.PostFormValue("type")
		p.Prefecture = r.PostFormValue("prefecture")
		p.Description = r.PostFormValue("description")
		p.UserID = user.ID
		if err := p.CreatePost(); err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/index", 302)
	}
}

func postEdit(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		_, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		p, err := models.GetPost(id)
		if err != nil {
			log.Println(err)
		}
		// 取得したpost idを渡す
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
			log.Println(err)
		}
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		p, err := models.GetPost(id)
		title := r.PostFormValue("title")
		description := r.PostFormValue("description")
		post := &models.Post{
			ID:          id,
			Title:       title,
			Type:        p.Type,
			Prefecture:  p.Prefecture,
			Description: description,
			UserID:      user.ID}
		if err := post.UpdatePost(); err != nil {
			log.Println(err)
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
			log.Println(err)
		}
		t, err := models.GetPost(id)
		if err != nil {
			log.Println(err)
		}
		if err := t.DeletePost(); err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/index", 302)
	}
}
