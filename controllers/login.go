package controllers

import (
	"encoding/json"
	"github.com/google/logger"
	"net/http"
	"net/url"
)

func LoginController(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Only GET is allowed", 403)
		return
	}

	q := r.URL.Query()
	if q.Get("error") != "" {
		// might be access_denied if user cancels

		// also we should handle other errors (check error_description)
	}

	// if user is redirected back by Discord, query will have
	// code field. Otherwise we should redirect him to Discord.
	if q.Get("code") == "" {
		w.Header().Set("Content-Type", "")
		http.Redirect(w, r, config.DiscordOAuthURL, 303)
		return
	}

	// OAuth2 flow
	discordResponse, err := http.PostForm("https://discord.com/api/oauth2/token", url.Values{
		"client_id":     {config.DiscordClient},
		"client_secret": {"00"},
		"grant_type":    {"authorization_code"},
		"code":          {q.Get("code")},
		"redirect_uri":  {config.ServerAddress + "login"},
		"scope":         {"identify connections"},
	})
	if err != nil {
		logger.Error(err.Error())
		http.Error(w, err.Error(), 500)
		return
	}

	var discordJSON map[string]interface{}
	err = json.NewDecoder(discordResponse.Body).Decode(&discordJSON)
	if err != nil {
		logger.Error(err.Error())
		http.Error(w, err.Error(), 500)
		return
	}
	if discordJSON["error"] != nil {
		logger.Error(discordJSON["error"].(string))
		http.Error(w, discordJSON["error"].(string), 500)
		return
	}
}
