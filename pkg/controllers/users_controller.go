package controllers

import (
	"encoding/json"
	"github.com/dimglaros/baladoroi-api/pkg/models"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	var err error
	var user *models.User

	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), 400)
	}

	err = db.Create(&user).Error
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	var err error
	var u models.User

	id := mux.Vars(r)["id"]
	if id == "" {
		http.Error(w, "Bad request", 400)
	}

	err = db.First(&u, "id = ?", id).Error
	if err != nil {
		http.Error(w, err.Error(), 404)
	}

	err = json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), 400)
	}

	err = db.Updates(&u).Error
	if err != nil {
		http.Error(w, err.Error(), 400)
	}

	w.WriteHeader(http.StatusOK)
}
