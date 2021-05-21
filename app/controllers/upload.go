package controllers

// show.htmlとpost.htmlで画像をアップロードするためのファイル

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
	"path/filepath"
)

func handleUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		generateHTML(w, nil, "layout", "public_navbar", "imageUpload")
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
		ext := filepath.Ext(uploadedFileName) //  アップロードされたファイル名の拡張子を取得
		if ext != ".jpeg" && ext != ".jpg" {
			log.Printf("ext Type:%v\n", ext)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		// f, err := os.Create("/usr/share/nginx/html/media/" + uploadedFileName)	// EC2用
		f, err := os.Create("/var/www/image/" + uploadedFileName) //ローカル用
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()

		io.Copy(f, file)
		http.Redirect(w, r, "/imageShow", 302)
	}
}

func handleShow(w http.ResponseWriter, r *http.Request) {
	// dir, err := os.Open("/usr/share/nginx/html/media/")	// EC2用
	dir, err := os.Open("/var/www/image/") // ローカル用
	defer dir.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	allImageNames, err := dir.Readdirnames(-1) // "/var/www/image/"のファイルの名前を配列に格納
	if err != nil {
		log.Fatalln("No files")
	}
	var decodeAllImages []image.Image // 配列decodeAllImagesを宣言
	// 全ての画像をデコード、リサイズしてdecodeAllImageseに格納
	for _, imageName := range allImageNames {
		// file, _ := os.Open("/usr/share/nginx/html/media/" + imageName)	// EC2用
		file, _ := os.Open("/var/www/image/" + imageName) //ローカル用
		defer file.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		decodeImage, _, err := image.Decode(file)
		m := resize.Resize(300, 0, decodeImage, resize.Lanczos3)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		decodeAllImages = append(decodeAllImages, m)
	}
	generateHTMLWithImage(w, nil, decodeAllImages, "layout", "public_navbar", "imageShow")
}

func generateHTMLWithImage(w http.ResponseWriter, data interface{}, decodeAllImages []image.Image, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}

	var encordImages []string
	for _, decodeImage := range decodeAllImages {
		buffer := new(bytes.Buffer)
		if err := jpeg.Encode(buffer, decodeImage, nil); err != nil {
			log.Fatalln("Unable to encode image.")
		}
		str := base64.StdEncoding.EncodeToString(buffer.Bytes())
		encordImages = append(encordImages, str)
	}

	image := map[string]interface{}{"Images": encordImages}
	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", image)
}
