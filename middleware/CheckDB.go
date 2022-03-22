package middleware

import (
	"net/http"

	"github.com/GeorgeHN666/proyect-Onlife-Server/db"
)

func CheckDB(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if db.DBConnectionCheck() == 1 {
			http.Error(w, "The Connection With Database Is Lost", 500)
			return
		}

		next.ServeHTTP(w, r)
	}

}
