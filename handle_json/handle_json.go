package handle_json

import (
	"encoding/json"
	"log"
	"net/http"
)

func RespondWithError(w http.ResponseWriter, code int, err error) {
	if code > 499 {
		log.Printf("Error with 5XX code: %s", err)
	}

	RespondWithJSON(w, code, map[string]string{"error": err.Error()})
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)

	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError) // error code, should be 500
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code) // 200 is default
	w.Write(data)
	
}