package hash

import (
	"gohub/pkg/logger"
	"golang.org/x/crypto/bcrypt"
)

// BcryptHash generate encrypt hash password
func BcryptHash(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	logger.LogIf(err)

	return string(bytes)
}

// BcryptCheck Compares a bcrypt hashed password
func BcryptCheck(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// BcryptIsHashed check string is hashed
func BcryptIsHashed(str string) bool {
	return len(str) == 60
}
