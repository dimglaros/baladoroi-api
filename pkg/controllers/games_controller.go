package controllers

import (
	"encoding/json"
	"github.com/dimglaros/baladoroi-api/pkg/models"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"net/http"
)

func CreateGame(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	var err error
	var g models.Game

	err = json.NewDecoder(r.Body).Decode(&g)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = db.Debug().Create(&g).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = db.Debug().Model(&models.User{}).Where("id = ?", g.HostID).First(&g.Host).Error
	if err != nil {
		db.Unscoped().Delete(&g)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func GetGame(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	var err error
	var g models.Game

	id := mux.Vars(r)["id"]
	if id == "" {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	err = db.Debug().Preload(clause.Associations).First(&g, "id = ?", id).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(&g)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func UpdateGame(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	var err error
	var g models.Game

	id := mux.Vars(r)["id"]
	if id == "" {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	err = db.First(&g, "id = ?", id).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&g)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = db.Updates(&g).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
