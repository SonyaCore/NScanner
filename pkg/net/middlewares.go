package net

import (
	"Scanner/pkg/validate"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func ValidateMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var Scanner Network

		bodyBytes, _ := io.ReadAll(r.Body)

		json.Unmarshal(bodyBytes, &Scanner)

		r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		err := validate.ValRange(Scanner.Range)
		if err != nil {
			_ = json.NewEncoder(w).Encode(
				&Error{err.Error(), http.StatusFailedDependency},
			)
			return
		}

		err = validate.ValProtocols(Scanner.Protocol)
		if err != nil {
			_ = json.NewEncoder(w).Encode(
				&Error{err.Error(), http.StatusFailedDependency},
			)
			return
		}

		next(w, r)
	}
}
