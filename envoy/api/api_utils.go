package api

import (
	"encoding/json"
	"net/http"
)

type TokenMessage struct {
	Token string `json:"token"`
}

type CountMessage struct {
	Count int64 `json:"count"`
}

func RespondWithJSON(w http.ResponseWriter, status int, object interface{}) {
	w.Header().Set("Content-Type", "application/json")

	body, err := json.Marshal(object)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		if status != http.StatusOK {
			w.WriteHeader(status)
		}
		w.Write(body)
	}
}
