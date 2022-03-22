package jwt

import (
	"time"

	"github.com/GeorgeHN666/proyect-Onlife-Server/models"
	jwt "github.com/dgrijalva/jwt-go"
)

func GenerateJWT(t models.User) (string, error) {

	pass := []byte("andresbeatrizcuzco")

	payload := jwt.MapClaims{
		"_id":   t.ID.Hex(),
		"name":  t.Name,
		"last":  t.Last,
		"email": t.Email,
		"bDay":  t.Bday,
		"bio":   t.Bio,
		"loc":   t.Loc,
		"web":   t.Web,
		"exp":   time.Now().Add(time.Hour * 1460).Unix(),
	}

	setting := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	tokenStr, err := setting.SignedString(pass)

	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil

}
