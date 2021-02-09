package controllers

import (
	"fmt"
	"net/http"
)

func LoginController(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Only GET is allowed", 403)
		return
	}

	q := r.URL.Query()
	if q.Get("error") != "" {

	}

	if q.Get("code") == "" {
		w.Header().Set("Content-Type", "")
		http.Redirect(w, r, config.DiscordOAuthURL, 303)
		return
	}

	fmt.Fprint(w, "Code is: ", q.Get("code"))
}
