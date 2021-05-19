package password

import (
	"github.com/dfzhou6/goblog/pkg/logger"
	"golang.org/x/crypto/bcrypt"
)

func Hash(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	logger.LogError(err)
	return string(bytes)
}

func CheckHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(hash))
	logger.LogError(err)
	return err == nil
}

func IsHashed(str string) bool {
	return len(str) == 60
}
