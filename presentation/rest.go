package presentation

import (
	"net/http"

	"github.com/samuelsih/gosckt/business"
)

// TODO
func WriteRestOutput[T any](w http.ResponseWriter, r *http.Request, out T, cr *business.CommonResponse) {
	println(w, r, out, cr)
}


// TODO
func ReadRestInput[T any](w http.ResponseWriter, r *http.Request, in *T, cr *business.CommonRequest) {
	println(w, r, in, cr)
}
