package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/samuelsih/gosckt/api/business/guest"
	"github.com/samuelsih/gosckt/api/transport"
)

func register(deps guest.Guest) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		in := guest.Guest_RegisterIn{}
	
		err := transport.ReadInput(r, &in)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
		}

		out := deps.Register(&in)

		transport.WriteOutput(w, out, out.CommonResponse)
	}
}

func login(deps guest.Guest) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		in := guest.Guest_LoginIn{}

		err := transport.ReadInput(r, &in)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
			return
		}

		out := deps.Login(&in)
		transport.WriteOutput(w, out, out.CommonResponse)
	}
}

func logout(rdb *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		
	}
}

func test(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("HAI SLUR"))
}