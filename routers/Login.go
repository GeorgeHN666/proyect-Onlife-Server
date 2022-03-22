package routers

import (
	"encoding/json"
	"net/http"

	"github.com/GeorgeHN666/proyect-Onlife-Server/db"
	"github.com/GeorgeHN666/proyect-Onlife-Server/jwt"
	"github.com/GeorgeHN666/proyect-Onlife-Server/models"
)

func LoginEndPoint(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "applization/json")

	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "There was an error reading the data", 400)
		return
	}

	user, good := db.Login(t.Email, t.Pass)
	if !good {
		http.Error(w, "Invalid User or Password, Try Again Later", 400)
		return
	}

	token, err := jwt.GenerateJWT(user)
	if err != nil {
		http.Error(w, "There was an error trying to generate the token", 400)
		return
	}

	response := models.LoginResponse{
		Token: token,
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(response)

}
