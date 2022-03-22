package db

import "golang.org/x/crypto/bcrypt"

func HashingPassword(password string) (string, error) {

	cost := 6

	pass, err := bcrypt.GenerateFromPassword([]byte(password), cost)

	return string(pass), err

}
