package controllers

import (
	"github.com/google/logger"
	"net/http"
	"sotru-web/models"
	"sotru-web/views"
)

func BlacklistController(w http.ResponseWriter, r *http.Request) {
	entries, err := models.GetAllBlacklistEntries()
	if err != nil {
		logger.Error(err.Error())
		http.Error(w, err.Error(), 500)
		return
	}

	err = views.ExecuteBlacklist(entries, w, r)
	if err != nil {
		logger.Error(err.Error())
		http.Error(w, err.Error(), 500)
		return
	}
}
