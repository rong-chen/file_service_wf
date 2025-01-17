package utils

import "golang.org/x/crypto/bcrypt"

func GenerateFromPassword(password string) string {
	byteSize, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(byteSize)
}

func CompareHashAndPassword(hash, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
