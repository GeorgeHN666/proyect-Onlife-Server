package middleware

import (
	"net/http"

	"github.com/GeorgeHN666/proyect-Onlife-Server/routers"
)

func ValidToken(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.ProcessToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "There was a problem in the Token", http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	}

}
