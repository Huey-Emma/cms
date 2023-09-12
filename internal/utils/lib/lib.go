package lib

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
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

func ReadJSON(w http.ResponseWriter, r *http.Request, v any) error {
	r.Body = http.MaxBytesReader(w, r.Body, int64(1_000_000))

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(v)
	if errors.Is(err, io.EOF) {
		return fmt.Errorf("request body has no content")
	}

	if errors.Is(err, io.ErrUnexpectedEOF) {
		return fmt.Errorf("malformed json")
	}

	var syntaxError *json.SyntaxError
	var unmarshalTypeError *json.UnmarshalTypeError

	if errors.As(err, &syntaxError) {
		return fmt.Errorf("malformed at position %d", syntaxError.Offset)
	}

	if errors.As(err, &unmarshalTypeError) {
		if unmarshalTypeError.Field != "" {
			return fmt.Errorf("invalid field type at %s", unmarshalTypeError.Field)
		}

		return fmt.Errorf("invalid type at position %d", unmarshalTypeError.Offset)
	}

	if err != nil {
		return err
	}

	return nil
}
