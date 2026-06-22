package jwt

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	UserID string `json:"sub"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}
