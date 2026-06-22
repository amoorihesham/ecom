package bcryptx

import "golang.org/x/crypto/bcrypt"

type HashService interface {
	Hash(password string) (string, error)
	Compare(raw string, hash string) bool
}

type service struct {
	bcryptCost int
}

func New(cfg *Config) HashService {
	return &service{bcryptCost: cfg.Cost}
}

func (s *service) Hash(password string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(password), s.bcryptCost)
	if err != nil {
		return "", err
	}
	return string(b), nil

}
func (s *service) Compare(raw string, hash string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(raw)); err != nil {
		return false
	}
	return true
}
