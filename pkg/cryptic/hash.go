package cryptic

import (
	"crypto/md5"

	"golang.org/x/crypto/bcrypt"
)

func TextHash(plainText string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(plainText), bcrypt.DefaultCost)
	return string(hashed), err
}

func CompareHash(plainText string, hashedText string) error {
	hashed := []byte(hashedText)
	return bcrypt.CompareHashAndPassword(hashed, []byte(plainText))
}

func MD5(data []byte) []byte {
	s := md5.Sum(data)
	return s[:]
}
