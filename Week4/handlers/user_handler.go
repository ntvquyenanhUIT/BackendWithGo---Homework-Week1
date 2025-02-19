package handlers

import (
	"net/http"
	"os"
	"path"

	"github.com/jmoiron/sqlx"
)

type UserHandler struct {
	db *sqlx.DB
}

func NewUserHandler(db *sqlx.DB) *UserHandler {

	return &UserHandler{db: db}

}

func (h *UserHandler) UploadImage(w http.ResponseWriter, r *http.Request) {

	err := r.ParseMultipartForm(10 << 20) // parse the form-data and stored in mem (or disk if memmo not enough)
	if err != nil {
		http.Error(w, "Error parsing multipart form", http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("image") // retrieves the file from the parsed data in memory
	if err != nil {
		http.Error(w, "Error getting image file", http.StatusBadRequest)
		return
	}

	defer file.Close()
	uploadsDir := "./uploads"
	if _, err := os.Stat(uploadsDir); os.IsNotExist(err) {
		os.Mkdir(uploadsDir, os.ModePerm)
	}

	filePath := path.Join(uploadsDir, handler.Filename)
	dst, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "Failed to create file", http.StatusInternalServerError)
		return
	}

	defer dst.Close()
	username := r.FormValue("username")
	query := `UPDATE user SET image_path = ?
				WHERE username = ?`
	_, err = h.db.Exec(query, filePath, username)
	if err != nil {
		http.Error(w, "Failed to update user image", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)

}
