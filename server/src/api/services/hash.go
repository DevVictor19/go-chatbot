package services

import "golang.org/x/crypto/bcrypt"

func Hash(value string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func Compare(hashed, value string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(value))
	return err == nil
}
