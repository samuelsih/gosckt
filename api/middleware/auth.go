package middleware

import (
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/samuelsih/gosckt/api/business"
	"github.com/samuelsih/gosckt/api/transport"
)

func Auth(rdb *redis.Client) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			out := business.CommonResponse{}
		
			cookie, err := r.Cookie("remember_token")
			if err != nil || err == http.ErrNoCookie {
				out.SetError(http.StatusUnauthorized, "you must login first")
				transport.WriteOutput(w, out, out)
				return
			}

			_, err = rdb.Get(r.Context(), cookie.Value).Result()
			if err != nil {
				out.SetError(http.StatusUnauthorized, "unknown cookie")
				transport.WriteOutput(w, out, out)
				return
			} 

			next.ServeHTTP(w, r)
		})
	}
}