package net

import (
	"Scanner/pkg/util"
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

		if Scanner.Range == "" {
			Scanner.Range = DefaultRange
		}

		startPort, endPort, err := util.SplitPorts(Scanner.Range)
		if err != nil {
			_ = json.NewEncoder(w).Encode(
				&Error{err.Error(), http.StatusFailedDependency},
			)
			return
		}

		// Validate start Port
		err = validate.ValRange(startPort)
		if err != nil {
			returnError(w, err, http.StatusFailedDependency)
			return
		}

		// Validate end Port
		err = validate.ValRange(endPort)
		if err != nil {
			returnError(w, err, http.StatusFailedDependency)
			return
		}

		// Validate Protcol
		err = validate.ValProtocols(Scanner.Protocol)
		if err != nil {
			returnError(w, err, http.StatusFailedDependency)
			return
		}

		next(w, r)
	}
}

func returnError(writer http.ResponseWriter, err error, statusCode int) (message interface{}) {
	return json.NewEncoder(writer).Encode(
		&Error{err.Error(), statusCode},
	)
}
