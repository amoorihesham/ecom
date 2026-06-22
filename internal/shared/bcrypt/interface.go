package bcryptx

type HashService interface {
	Hash(password string) (string, error)
	Compare(raw string, hash string) bool
}
