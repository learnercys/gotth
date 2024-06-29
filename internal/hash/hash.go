package hash

type PasswordHash interface {
	ComparePasswordAndHash(password, hash string) (match bool, err error)
	GenerateFromPassword(password string) (encodedHash string, err error)
}
