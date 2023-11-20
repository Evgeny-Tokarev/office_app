package util

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func GetUniqueFileName(filePath string) string {
	base := filepath.Base(filePath)
	ext := filepath.Ext(base)
	name := strings.TrimSuffix(base, ext)
	dir := filepath.Dir(filePath)

	for i := 0; ; i++ {
		modifiedName := name
		if i > 0 {
			modifiedName = name + "(" + strconv.Itoa(i) + ")"
		}
		newFileName := filepath.Join(dir, modifiedName+ext)

		if _, err := os.Stat(newFileName); os.IsNotExist(err) {
			return newFileName
		}
	}
}
func SendTranscribedError(w http.ResponseWriter, errText string, status int) {

	responseBody, err := json.Marshal(errors.New(errText))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(status)
	_, err = w.Write(responseBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func ConvertToRegularString(ss sql.NullString) string {
	if ss.Valid {
		return ss.String
	}
	return ""
}
