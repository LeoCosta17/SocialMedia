package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			log.Fatal(err)
		}
	}
}

func WriteJSONError(w http.ResponseWriter, statusCode int, error string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	data := map[string]any{
		"error": error,
	}

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatal(err)
	}
}
