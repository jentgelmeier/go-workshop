package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"frontendmasters.com/go/museum/data"
)

func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// api/exhibitions?id=34
	id := r.URL.Query()["id"]
	if id != nil {
		finalID, err := strconv.Atoi(id[0])
		if err == nil && finalID < len(data.GetAll()) {
			json.NewEncoder(w).Encode(data.GetAll()[finalID])
		} else {
			http.Error(w, "Invalid Exhibition", http.StatusBadRequest)
		}
	} else {
		json.NewEncoder(w).Encode(data.GetAll())
	}
}