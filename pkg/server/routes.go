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

	// Field actions
	s.Router.HandleFunc("/field", func(w http.ResponseWriter, r *http.Request) {
		controllers.CreateField(w, r, s.DB)
	}).Methods("POST")

	s.Router.HandleFunc("/field/{id}", func(w http.ResponseWriter, r *http.Request) {
		controllers.GetField(w, r, s.DB)
	}).Methods("GET")

	s.Router.HandleFunc("/field/{id}", func(w http.ResponseWriter, r *http.Request) {
		controllers.UpdateField(w, r, s.DB)
	}).Methods("PUT")

	s.Router.HandleFunc("/field/{id}", func(w http.ResponseWriter, r *http.Request) {
		controllers.DeleteField(w, r, s.DB)
	}).Methods("DELETE")

	// Game actions
	s.Router.HandleFunc("/game", func(w http.ResponseWriter, r *http.Request) {
		controllers.CreateGame(w, r, s.DB)
	}).Methods("POST")

	s.Router.HandleFunc("/game/{id}", func(w http.ResponseWriter, r *http.Request) {
		controllers.GetGame(w, r, s.DB)
	}).Methods("GET")

	s.Router.HandleFunc("/game/{id}", func(w http.ResponseWriter, r *http.Request) {
		controllers.UpdateGame(w, r, s.DB)
	}).Methods("PUT")

	// Participation actions
	s.Router.HandleFunc("/participation", func(w http.ResponseWriter, r *http.Request) {
		controllers.CreateParticipation(w, r, s.DB)
	}).Methods("POST")

	s.Router.HandleFunc("/participation/{id}", func(w http.ResponseWriter, r *http.Request) {
		controllers.GetParticipation(w, r, s.DB)
	}).Methods("GET")

	s.Router.HandleFunc("/participation/{id}", func(w http.ResponseWriter, r *http.Request) {
		controllers.UpdateParticipation(w, r, s.DB)
	}).Methods("PUT")
}
