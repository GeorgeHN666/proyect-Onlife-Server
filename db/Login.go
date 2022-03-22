package db

import (
	"github.com/GeorgeHN666/proyect-Onlife-Server/models"
	"golang.org/x/crypto/bcrypt"
)

func Login(email, password string) (models.User, bool) {

	user, exist, _ := UserExist(email)
	if !exist {
		return user, false
	}

	passwordEntered := []byte(password)

	passwordStored := []byte(user.Pass)

	err := bcrypt.CompareHashAndPassword(passwordStored, passwordEntered)
	if err != nil {
		return user, false
	}

	return user, true

}
