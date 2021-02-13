package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/google/logger"
	"net/http"
	"net/url"
	"sotru-web/models"
	"sotru-web/utils"
	"time"
)

func LoginController(w http.ResponseWriter, r *http.Request) {
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
	if session.Values["userid"] != nil {
		fmt.Fprint(w, "You are authorized")
		return
	}

	q := r.URL.Query()
	if q.Get("error") != "" {
		// might be access_denied if user cancels
		logger.Error(q.Get("error"))
		http.Error(w, q.Get("error"), 500)
		return
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
		"client_secret": {config.DiscordSecret},
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

	// get OAuth2 tokens from JSON
	accessToken := discordJSON["access_token"].(string)
	refreshToken := discordJSON["refresh_token"].(string)
	expiresIn := discordJSON["expires_in"].(float64)
	accessExpiration := time.Now().Add(time.Second * time.Duration(expiresIn))

	// get information using Discord API
	discordSession, err := discordgo.New("Bearer " + accessToken)
	if err != nil {
		logger.Error(err.Error())
		http.Error(w, err.Error(), 500)
		return
	}
	discordUser, err := discordSession.User("@me")
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

	// ip will be an empty string, if not found
	ip, _ := utils.GetIP(r)

	user, err := models.GetUser(discordUser.ID)
	if err != nil {
		// there is no user with the specified ID => register user.
		if err == sql.ErrNoRows {
			user, err = models.CreateUser(discordUser.ID, discordUser.String(), time.Now(), time.Now(),
				discordUser.AvatarURL(""), accessToken, refreshToken, accessExpiration)
		}

		// second check in case we create user
		if err != nil {
			logger.Error(err.Error())
			http.Error(w, err.Error(), 500)
			return
		}
	}

	_ = user.SetLastLogin(time.Now())

	// if user's used this IP, we should find it and update last usage
	if ip != "" {
		userIPs, _ := user.GetIPs()
		foundIP := false
		for i := 0; i < len(userIPs); i++ {
			if userIPs[i].IP == ip {
				userIPs[i].LastUsed = time.Now()
				foundIP = true
				return
			}
		}
		// if we haven't found anything, we should add this IP
		if !foundIP {
			_ = user.AddIP(ip, time.Now())
		}
	}

	if xbox != "" {
		// the same logic for Xbox gametags
		userXboxes, _ := user.GetXboxes()
		foundXbox := false
		for i := 0; i < len(userXboxes); i++ {
			if userXboxes[i].Xbox == xbox {
				userXboxes[i].LastUsed = time.Now()
				foundXbox = true
				return
			}
		}
		// if we haven't found anything, we should add this Xbox
		if !foundXbox {
			_ = user.AddXbox(ip, time.Now())
		}
	}

	session.Values["userid"] = user.ID
	_ = session.Save(r, w)

	fmt.Fprintf(w, "Your IP: %s. Your Xbox: %s.\nYour User entity:\n%+v", ip, xbox, user)
	return
}
