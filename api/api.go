package api

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/samuelsih/gosckt/api/business/guest"
	"github.com/samuelsih/gosckt/api/middleware"
	"github.com/samuelsih/gosckt/config"
)

const (
	GET = http.MethodGet
	POST = http.MethodPost
)

var (
	r *chi.Mux
	psql *pgxpool.Pool
	rdb *redis.Client
)

func Serve() {
	psql = config.PsqlSetup()

	rdb = config.RedisSetup()

	guestDeps := guest.Guest{Psql: psql, Rdb: rdb}

	r = chi.NewRouter()
	r.Use(middleware.WrapMiddleware()...)

	addRoute("/register", POST, register(guestDeps))
	addRoute("/login", POST, login(guestDeps))
	addRouteWithMiddleware("/test", GET, test, middleware.Auth(rdb))
	addRouteWithMiddleware("/logout", POST, logout(), middleware.Auth(rdb))

	log.Println("Listening on port :8080")
	
	err := http.ListenAndServeTLS(":8080", "server.crt", "server.key", r)
	if err != nil {
		log.Panic(err)
	}
}

