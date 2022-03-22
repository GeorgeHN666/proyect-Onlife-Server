package models

import (
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JWTClaims struct {
	Email string             `json:"email"`
	IDJWT primitive.ObjectID `bson:"_id"  json:"_id,omitempty"`
	jwt.StandardClaims
}
