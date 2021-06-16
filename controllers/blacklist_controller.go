package controllers

import (
	"github.com/google/logger"
	"net/http"
	"sotru-web/views"
)

func BlacklistController(w http.ResponseWriter, r *http.Request) {
	err := views.ExecuteBlacklist(w, r)
	if err != nil {
		logger.Error(err.Error())
		http.Error(w, err.Error(), 500)
		return
	}
}
