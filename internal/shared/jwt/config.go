package jwt

import "time"

type Config struct {
	Secret       string
	Issuer       string
	AccessExpiry time.Duration
}
