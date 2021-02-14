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
	if session.Values["userid"] == nil {
		// user is not authorized
		w.Header().Set("Content-Type", "")
		http.Redirect(w, r, "/login", 303)
		return
	}

	user, err := models.GetUser(session.Values["userid"].(string))
	if err != nil {
		logger.Error(err.Error())
		http.Error(w, err.Error(), 500)
		return
	}
	err = views.ExecuteIndex(user, w)
	if err != nil {
		logger.Error(err.Error())
		http.Error(w, err.Error(), 500)
		return
	}
}
