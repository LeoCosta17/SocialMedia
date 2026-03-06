package request

import (
	"encoding/json"
	"net/http"
)

func ReadJSON(w http.ResponseWriter, r *http.Request, data any) error {
	maxBytes := 1_048_578

	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))
	json.NewDecoder(r.Body).DisallowUnknownFields()

	return json.NewDecoder(r.Body).Decode(data)
}
