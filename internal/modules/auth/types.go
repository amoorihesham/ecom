package auth

type AuthConfig struct {
	BcryptCost int
	JWTSecret  string
}
