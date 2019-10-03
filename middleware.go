package httpext

import (
	"encoding/json"
	"net/http"
)

type responseTemplate struct {
	Error   string      `json:"error,omitempty"`
	Payload interface{} `json:"payload,omitempty"`
}

type handler = func(r *http.Request) (*ResponseError, interface{})

func Wrap(handler handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err, result = handler(r)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(err.status)

		var response = responseTemplate{Error: err.Error(), Payload: result}
		var jsonResponse, merror = json.Marshal(response)
		if merror != nil {
			// Note: this should not happen
			http.Error(w, "failed to construct response", http.StatusInternalServerError)
		}

		_, _ = w.Write(jsonResponse)
	})
}
