package server

import (
	"encoding/json"
	"net/http"
)

// WriteError writes the error message msg as a json to w.
func WriteError(w http.ResponseWriter, code int, msg string) {
	WriteJSON(w, code, map[string]string{"error": msg})
}

// WriteJSON writes the given the payload as a json to w, setting the code as Header.
func WriteJSON(w http.ResponseWriter, code int, payload interface{}) {
	rsp, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(rsp)
}
