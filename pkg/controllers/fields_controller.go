package controllers

import (
	"encoding/json"
	"github.com/dimglaros/baladoroi-api/pkg/models"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
)

func CreateField(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	var err error
	var f *models.Field

	err = json.NewDecoder(r.Body).Decode(&f)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = db.Create(&f).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func GetField(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	var err error
	var f models.Field

	id := mux.Vars(r)["id"]
	if id == "" {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	err = db.First(&f, "id = ?", id).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(&f)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func UpdateField(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	var err error
	var f models.Field

	id := mux.Vars(r)["id"]
	if id == "" {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	err = db.First(&f, "id = ?", id).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&f)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = db.Updates(&f).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
