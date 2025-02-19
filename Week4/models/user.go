package models

type User struct {
	Username  string `json:"username" db:"username"`
	Password  string `json:"password" db:"password"`
	ImagePath string `json:"image_path" db:"image_path"`
}
