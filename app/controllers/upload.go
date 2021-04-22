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
	generateHTML(w, nil, "layout", "public_navbar", "upload")
	// if r.Method != "POST" {
	// 	http.Error(w, "Allowed POST method only", http.StatusMethodNotAllowed)
	// 	return
	// }

	if r.Method == "POST" {
		err := r.ParseMultipartForm(32 << 20) // maxMemory
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		file, _, err := r.FormFile("upload")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer file.Close()

		f, err := os.Create("./test.jpg")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()

		io.Copy(f, file)
		http.Redirect(w, r, "/show", 302)
	}
}

func handleShow(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("./test.jpg")
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
	// templates := template.Must(template.ParseFiles(files...))

	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, *m, nil); err != nil {
		log.Fatalln("Unable to encode image.")
	}
	str := base64.StdEncoding.EncodeToString(buffer.Bytes())
	image := map[string]interface{}{"Image": str}
	// renderTemplate(w, tmpl, data)
	// generateHTML(w, data, "layout", "public_navbar", "show")
	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", image)
}
