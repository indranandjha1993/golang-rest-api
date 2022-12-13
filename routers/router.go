package routers

import (
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func Routers() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	s := r.PathPrefix("/api").Subrouter()
	UserRouter(s)

	r.Use(loggingMiddleware)
	return r
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Info(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
