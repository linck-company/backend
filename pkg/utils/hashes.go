package utils

import (
	"encoding/base64"

	"golang.org/x/crypto/bcrypt"
)

func Base64Decode(str string) string {
	bytes, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return str
	}
	return string(bytes)
}

func BcryptEncode(str string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(str), bcrypt.MinCost)
	return string(hash)
}

func BcryptCompare(str, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(str))
	return err == nil
}
