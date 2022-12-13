package routers

import (
	"github.com/gorilla/mux"
	"site-api/Handlers"
)

func UserRouter(r *mux.Router) *mux.Router {
	s := r.PathPrefix("/users").Subrouter()
	s.HandleFunc("", Handlers.GetAllUsers).Methods("GET")
	s.HandleFunc("/{id}", Handlers.GetUserById).Methods("GET")
	s.HandleFunc("", Handlers.CreateUser).Methods("POST")
	s.HandleFunc("", Handlers.UpdateUser).Methods("PUT")
	s.HandleFunc("/{id}", Handlers.DeleteUser).Methods("DELETE")
	return s
}
