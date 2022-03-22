package routers

import (
	"errors"
	"strings"

	"github.com/GeorgeHN666/proyect-Onlife-Server/db"
	"github.com/GeorgeHN666/proyect-Onlife-Server/models"
	"github.com/dgrijalva/jwt-go"
)

var Email string

var GlobalID string

func ProcessToken(t string) (*models.JWTClaims, bool, string, error) {

	password := []byte(" Here Goes Your Password ")

	claims := &models.JWTClaims{}

	splitToken := strings.Split(t, "Bearer")

	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("invalid token format")
	}

	t = strings.TrimSpace(splitToken[1])

	token, err := jwt.ParseWithClaims(t, claims, func(t *jwt.Token) (interface{}, error) {
		return password, nil
	})

	if err == nil {

		_, exist, _ := db.UserExist(claims.Email)
		if exist {
			Email = claims.Email
			GlobalID = claims.IDJWT.Hex()
		}

		return claims, exist, GlobalID, nil
	}

	if !token.Valid {
		return claims, false, string(""), errors.New("invalid Token")
	}

	return claims, false, string(""), err

}
