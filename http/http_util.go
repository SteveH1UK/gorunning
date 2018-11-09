package http

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/SteveH1UK/gorunning"
)

func errorWithJSON(w http.ResponseWriter, message string, code int, validationErrors []root.ValidationError) {
	errors := root.Error{}
	errors.HTTPCode = code
	errors.Message = message
	errors.ValidationErrors = validationErrors
	json, jsonParseErr := json.Marshal(errors)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	if jsonParseErr == nil {
		w.Write(json)
	} else {
		fmt.Fprintf(w, "{message: %q}", message)
	}
}

func responseWithJSON(w http.ResponseWriter, json []byte, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	w.Write(json)
}

func getFriendlyName(r *http.Request) string {
	params := r.URL.Query()
	friendlyName := params.Get(":friendly_name")
	return friendlyName
}

var callDecodeJSONFromBody = decodeJSONFromBody

func decodeJSONFromBody(body io.ReadCloser, athelete *root.NewAthelete) error {
	return json.NewDecoder(body).Decode(athelete)
}
