package controllers

import (
	"encoding/json"
	"github.com/dimglaros/baladoroi-api/pkg/models"
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
