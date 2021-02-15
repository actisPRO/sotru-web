package controllers

import (
	"github.com/google/logger"
	"net/http"
	"sotru-web/models"
	"sotru-web/views"
)

func IndexController(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "sotru-web")
	if err != nil {
		logger.Error(err.Error())
		http.Error(w, err.Error(), 500)
		return
	}

	user := models.User{}

	if session.Values["userid"] != nil {
		user, err = models.GetUser(session.Values["userid"].(string))
		if err != nil {
			logger.Error(err.Error())
			http.Error(w, err.Error(), 500)
			return
		}
	}

	err = views.ExecuteIndex(user, w)
	if err != nil {
		logger.Error(err.Error())
		http.Error(w, err.Error(), 500)
		return
	}
}
