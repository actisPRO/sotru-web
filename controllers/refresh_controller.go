package controllers

import (
	"github.com/bwmarrin/discordgo"
	"github.com/google/logger"
	"net/http"
	"sotru-web/models"
	"time"
)

func RefreshController(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "sotru-web")
	if err != nil {
		logger.Error(err.Error())
		http.Error(w, err.Error(), 500)
		return
	}
	if session.Values["userid"] == nil {
		// user is not authorized => login
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

	discordSession, err := discordgo.New("Bearer " + user.AccessToken)
	if err != nil {
		logger.Error(err.Error())
		http.Error(w, err.Error(), 500)
		return
	}

	connections, err := discordSession.UserConnections()
	if err != nil {
		logger.Error(err.Error())
		http.Error(w, err.Error(), 500)
		return
	}

	// getting xbox tag from connections
	xbox := ""
	for i := 0; i < len(connections); i++ {
		if connections[i].Type == "xbox" {
			xbox = connections[i].Name
			break
		}
	}
	if xbox != "" {
		userXboxes, _ := user.GetXboxes()
		foundXbox := false
		for i := 0; i < len(userXboxes); i++ {
			if userXboxes[i].Xbox == xbox {
				_ = userXboxes[i].SetLastUsed(time.Now())
				foundXbox = true
				break
			}
		}
		if !foundXbox {
			_ = user.AddXbox(xbox, time.Now())
		}
	}

	// Success => redirect to the index page
	w.Header().Set("Content-Type", "")
	http.Redirect(w, r, "/", 303)
	return
}
