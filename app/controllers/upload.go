package controllers

import (
	"bytes"
	"encoding/base64"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"io"
	"log"
	"net/http"
	"os"
)

func handleUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Allowed POST method only", http.StatusMethodNotAllowed)
		return
	}

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

	f, err := os.Create("/tmp/test.jpg")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer f.Close()

	io.Copy(f, file)
	generateHTML(w, nil, "layout", "public_navbar", "upload")
	http.Redirect(w, r, "/show", http.StatusFound)
}

func handleShow(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("/tmp/test.jpg")
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

	writeImageWithTemplate(w, "show", &m)
}

func writeImageWithTemplate(w http.ResponseWriter, tmpl string, m *image.Image) {
	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, *m, nil); err != nil {
		log.Fatalln("Unable to encode image.")
	}

	str := base64.StdEncoding.EncodeToString(buffer.Bytes())
	data := map[string]interface{}{"Title": tmpl, "Image": str}
	// renderTemplate(w, tmpl, data)
	generateHTML(w, data, "layout", "public_navbar", "show")
}
