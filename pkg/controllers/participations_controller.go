package controllers

import (
	"encoding/json"
	"github.com/dimglaros/baladoroi-api/pkg/models"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
)

func CreateParticipation(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	var err error
	var p models.Participation

	err = json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if p.UserID.String() == "" || p.GameID.String() == "" {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	err = db.Debug().Create(&p).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func GetParticipation(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	var err error
	var p models.Participation

	id := mux.Vars(r)["id"]
	if id == "" {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	err = db.Debug().First(&p, "id = ?", id).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func UpdateParticipation(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	var err error
	var p models.Participation

	id := mux.Vars(r)["id"]
	if id == "" {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	err = db.First(&p, "id = ?", id).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = db.Updates(&p).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
