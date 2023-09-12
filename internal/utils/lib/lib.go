package lib

import (
	"encoding/json"
	"net/http"
)

type Map[T any] map[string]T

func WriteJSON(w http.ResponseWriter, code int, header http.Header, v any) {
	w.Header().Add("content-type", "application/json")

	for k, v := range header {
		w.Header()[k] = v
	}

	w.WriteHeader(code)

	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		panic(err)
	}
}

func WriteError(w http.ResponseWriter, code int, err error) {
	WriteJSON(w, code, make(http.Header), Map[string]{"detail": err.Error()})
}
