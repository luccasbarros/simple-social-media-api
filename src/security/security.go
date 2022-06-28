package security

import "golang.org/x/crypto/bcrypt"

// Hash insert hash on strings
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

//VerifyPassword compares password and hash
func VerifyPassword(password string, hashPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(password), []byte(hashPassword))
}
