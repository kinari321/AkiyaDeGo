package controllers

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/nfnt/resize"
	"html/template"
	"image"
	"image/jpeg"
	"io"
	"log"
	"net/http"
	"os"
)

func handleUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		generateHTML(w, nil, "layout", "public_navbar", "upload")
	} else if r.Method == "POST" {
		err := r.ParseMultipartForm(32 << 20) // maxMemory
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		file, fileHeader, err := r.FormFile("upload")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer file.Close()

		uploadedFileName := fileHeader.Filename
		f, err := os.Create("source/img/" + uploadedFileName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()

		io.Copy(f, file)
		http.Redirect(w, r, "/show/", 302)
	}
}

func handleShow(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("source/img/")
	defer file.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	img, _, err := image.Decode(file)
	m := resize.Resize(300, 0, img, resize.Lanczos3)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	generateHTMLWithImage(w, nil, &m, "layout", "public_navbar", "show")
}

func generateHTMLWithImage(w http.ResponseWriter, data interface{}, m *image.Image, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}

	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, *m, nil); err != nil {
		log.Fatalln("Unable to encode image.")
	}
	str := base64.StdEncoding.EncodeToString(buffer.Bytes())
	image := map[string]interface{}{"Image": str}
	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", image)
}
