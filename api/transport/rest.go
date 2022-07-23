package transport

import (
	"encoding/json"
	"net/http"

	"github.com/samuelsih/gosckt/api/business"
)

func ReadInput[T any](r *http.Request, in *T) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	defer r.Body.Close()

	err := decoder.Decode(in)
	if err != nil {
		return err
	}

	return nil
}

func WriteOutput[T any](w http.ResponseWriter, out T, cr business.CommonResponse) {
	if cr.StatusCode == 0 {
		cr.StatusCode = 200
	}

	if cr.SetCookie != "" {
		http.SetCookie(w, &http.Cookie{
			Name:     "remember_token",
			Value:    cr.SetCookie,
			HttpOnly: true,
			MaxAge:   150,
			Secure:   true,
			Domain:   "localhost:8080",
		})
	}

	w.WriteHeader(cr.StatusCode)

	json.NewEncoder(w).Encode(out)
}
