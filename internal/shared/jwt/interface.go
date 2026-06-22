package jwt

import "ecom/internal/shared/httpx"

type JWTService interface {
	Generate(user httpx.AuthUser) (string, error)
	Parse(token string) (*httpx.AuthUser, error)
}
