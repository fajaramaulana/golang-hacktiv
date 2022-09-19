package controllers

import (
	"encoding/json"
	"net/http"
	"sesi6/webserver/repositories"
)

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	users, err := repositories.GetUser()
	if err != nil {
		writerJsonResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"message": "Error when get user",
			"error":   err.Error(),
		})
		return
	}

	writerJsonResponse(w, http.StatusOK, map[string]interface{}{
		"message": "Success get user",
		"data":    users,
	})
}

func writerJsonResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
