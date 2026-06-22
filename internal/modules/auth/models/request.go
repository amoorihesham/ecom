package models

type RegisterRequest struct {
	Email        string `json:"email"`
	Fullname     string `json:"fullname"`
	PasswordHash string `json:"password"`
}
