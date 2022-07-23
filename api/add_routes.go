package api

import (
	"net/http"
)


func addRoute(path string, method string, handler http.HandlerFunc) {
	switch method {
	case http.MethodGet:
		r.Get(path, handler)
	
	case http.MethodPost:
		r.Post(path, handler)
	
	case http.MethodOptions:
		r.Options(path, handler)

	case http.MethodDelete:
		r.Delete(path, handler)

	case http.MethodPut:
		r.Put(path, handler)

	case http.MethodPatch:
		r.Patch(path, handler)

	default:
		r.Get(path, handler)
	}
}

func addRouteWithMiddleware(path string, method string, handler http.HandlerFunc, middlewares ...func(http.Handler) http.Handler) {
	switch method {
	case http.MethodGet:
		r.With(middlewares...).Get(path, handler)
	
	case http.MethodPost:
		r.With(middlewares...).Post(path, handler)
	
	case http.MethodOptions:
		r.With(middlewares...).Options(path, handler)

	case http.MethodDelete:
		r.With(middlewares...).Delete(path, handler)

	case http.MethodPut:
		r.With(middlewares...).Put(path, handler)

	case http.MethodPatch:
		r.With(middlewares...).Patch(path, handler)

	default:
		r.With(middlewares...).Get(path, handler)
	}
}