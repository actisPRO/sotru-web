package controllers

import (
	"github.com/google/logger"
	"net/http"
)

func LogoutController(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Only GET is allowed", 403)
		return
	}

	session, err := store.Get(r, "sotru-web")
	if err != nil {
		logger.Error(err.Error())
		http.Error(w, err.Error(), 500)
		return
	}
	if session.Values["userid"] == nil {
		// user is not authorized
		w.Header().Set("Content-Type", "")
		http.Redirect(w, r, "/", 303)
		return
	}

	session.Values["userid"] = nil
	_ = session.Save(r, w)

	// logout was successful and we can redirect user to the index page
	w.Header().Set("Content-Type", "")
	http.Redirect(w, r, "/", 303)
	return
}
