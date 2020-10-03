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
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = db.Create(&user).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func UpdateUser(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	var err error
	var u models.User

	id := mux.Vars(r)["id"]
	if id == "" {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	err = db.First(&u, "id = ?", id).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = db.Updates(&u).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
