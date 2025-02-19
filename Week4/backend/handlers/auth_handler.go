package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"engineerpro_ex_week4/models"

	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	db *sqlx.DB
}

func NewAuthHanlder(db *sqlx.DB) *AuthHandler {

	return &AuthHandler{db: db}

}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {

	// Getting data from the request
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Error parsing JSON", http.StatusBadRequest)
		return
	}

	// Check if username exist
	var existingUser models.User
	query := `SELECT username FROM user where username = ?`
	err := h.db.Get(&existingUser, query, user.Username)
	if err == nil {
		http.Error(w, "Username Already Exist", http.StatusConflict)
		return
	} else if err != sql.ErrNoRows {
		http.Error(w, "Error checking username", http.StatusInternalServerError)
		return
	}

	// Insert new user
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error comparing password", http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword)
	query = `INSERT INTO user(username, password) VALUES (?, ?)`
	_, err = h.db.Exec(query, user.Username, user.Password)
	if err != nil {
		http.Error(w, "Error adding new user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Error parsing JSON", http.StatusInternalServerError)
		return
	}

	query := `SELECT username, password FROM user WHERE username = ?`
	var existingUser models.User
	err := h.db.Get(&existingUser, query, user.Username)
	if err == sql.ErrNoRows {
		http.Error(w, "username not exist", http.StatusUnauthorized)
		return
	} else if err != nil {
		fmt.Println(err)
		http.Error(w, "Error when fetching user", http.StatusInternalServerError)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password))
	if err != nil {
		http.Error(w, "Incorrect Password", http.StatusUnauthorized)
		return
	}

	// Trả data là username về cho client
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"username": existingUser.Username})

}
