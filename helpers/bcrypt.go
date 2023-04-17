package helpers

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(p string) (string, error) {
	salt := 8
	password := []byte(p)

	hash, err := bcrypt.GenerateFromPassword(password, salt)
	if err != nil {
		log.Fatal("error bcrypt hashing password")
		return "", err
	}

	return string(hash), nil
}

func CompareHash(h, p []byte) bool {
	err := bcrypt.CompareHashAndPassword(h, p)
	return err == nil
}
