package converter

import (
	"encoding/json"
	"net/http"
)

func ConvertToJSON(w http.ResponseWriter,data any) error {
	w.Header().Set("Content-Type","application/json")
	
	return json.NewEncoder(w).Encode(data)
}

func DecodeFromJSON(r *http.Request,result any) error{
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	return decoder.Decode(result)
}