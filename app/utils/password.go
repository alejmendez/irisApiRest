package utils

import "golang.org/x/crypto/bcrypt"

// Generate return a hashed password
func GenerateHash(raw string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(raw), 10)

	if err != nil {
		panic(err)
	}

	return string(hash)
}

// Verify compares a hashed password with plaintext password
func VerifyHash(hash string, raw string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(raw))
	return err == nil
}
