package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/GeorgeHN666/proyect-Onlife-Server/db"
)

func ReadCommentsEndPoint(w http.ResponseWriter, r *http.Request) {

	postID := r.URL.Query().Get("id")

	if len(postID) < 1 {
		http.Error(w, "The Param id Must Be Sended", http.StatusBadRequest)
		return
	}

	page := r.URL.Query().Get("pg")

	if len(page) < 1 {
		http.Error(w, "The param pg must be sended", http.StatusBadRequest)
		return
	}

	num, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, "The param pg must be an integer", http.StatusBadRequest)
		return
	}

	pg := int64(num)

	response, good := db.ReadComments(postID, pg)
	if !good {
		http.Error(w, "Fail to fetch the comments", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(response)

}
