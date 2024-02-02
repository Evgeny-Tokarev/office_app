package util

import (
	"bytes"
	"github.com/chai2010/webp"
	"github.com/disintegration/imaging"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

func SaveImage(r *http.Request) (string, int, error) {
	var closeError error
	file, header, err := r.FormFile("image")
	if err != nil {
		return "", http.StatusBadRequest, err
	}
	defer func(file multipart.File) {
		closeError = file.Close()
	}(file)

	img, err := imaging.Decode(file)
	if err != nil {
		return "", http.StatusUnprocessableEntity, err
	}
	var buffer bytes.Buffer

	err = webp.Encode(&buffer, img, nil)
	if err != nil {
		return "", http.StatusUnprocessableEntity, err
	}
	parts := strings.Split(header.Filename, ".")
	var filename string
	if len(parts) > 1 {
		filename = strings.Join(parts[:len(parts)-1], ".")
	} else {
		filename = parts[0]
	}
	webpFilePath := GetUniqueFileName("./images/" + filename + ".webp")
	//fmt.Println(filename, webpFilePath)
	webpFile, err := os.Create(webpFilePath)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}
	defer func(webpFile multipart.File) {
		closeError = file.Close()
	}(webpFile)

	_, err = webpFile.Write(buffer.Bytes())
	if err != nil {
		return "", http.StatusInternalServerError, err
	}
	return webpFilePath, http.StatusOK, closeError
}
