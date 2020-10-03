package server

import (
	"github.com/dimglaros/baladoroi-api/pkg/controllers"
	"net/http"
)

func (s *Server) InitializeRoutes() {
	// User actions
	s.Router.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		controllers.CreateUser(w, r, s.DB)
	}).Methods("POST")

	s.Router.HandleFunc("/user/{id}", func(w http.ResponseWriter, r *http.Request) {
		controllers.GetUser(w, r, s.DB)
	}).Methods("GET")

	s.Router.HandleFunc("/user/{id}", func(w http.ResponseWriter, r *http.Request) {
		controllers.UpdateUser(w, r, s.DB)
	}).Methods("PUT")
}
