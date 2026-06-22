package jwt

import (
	"ecom/internal/shared/httpx"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type service struct {
	secret string
	issuer string
	expiry time.Duration
}

func New(cfg *Config) JWTService {
	return &service{
		secret: cfg.Secret,
		issuer: cfg.Issuer,
		expiry: cfg.AccessExpiry,
	}
}

func (s *service) Generate(user httpx.AuthUser) (string, error) {

	claims := Claims{
		UserID: user.PublicID,
		Role:   user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    s.issuer,
			Subject:   user.PublicID,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.expiry)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(s.secret))
}

func (s *service) Parse(tokenStr string) (*httpx.AuthUser, error) {

	token, err := jwt.ParseWithClaims(
		tokenStr,
		&Claims{},
		func(token *jwt.Token) (any, error) {
			return []byte(s.secret), nil
		},
	)

	if err != nil || !token.Valid {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, err
	}

	return &httpx.AuthUser{
		PublicID: claims.UserID,
		Role:     claims.Role,
	}, nil
}
