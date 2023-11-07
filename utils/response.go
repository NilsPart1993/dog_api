package utils

import (
	"encoding/json"
	"net/http"
)

func ParseJSON(r *http.Request, v interface{}) error {
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(v); err != nil {
		return err
	}
	return nil
}

func RespondWithJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
	}
}

func RespondWithError(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	errorResponse := struct {
		Error string `json:"error"`
	}{
		Error: message,
	}
	if err := json.NewEncoder(w).Encode(errorResponse); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
