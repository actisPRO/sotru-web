package views

import (
	"html/template"
	"net/http"
	"sotru-web/models"
	"sotru-web/utils"
)

type IndexContent struct {
	Username  string
	Avatar    string
	JoinedAt  string
	Xbox      string
	VoiceTime string
	Warnings  []models.Warning
}

// Executes layout for the index page. If user is not authorized, use empty struct
func ExecuteIndex(user models.User, w http.ResponseWriter, r *http.Request) error {
	access := 0

	// check if user is empty
	if user.ID != "" {
		access = user.GetAccess()
	} else {
		w.Header().Set("Content-Type", "")
		http.Redirect(w, r, "/login", http.StatusFound)
		return nil
	}

	voiceTime := ""
	vtd, err := user.GetVoiceTime()
	if err == nil {
		voiceTime = utils.FormatDuration(vtd)
	}

	content := IndexContent{
		Username:  user.Username,
		Avatar:    user.AvatarURL,
		VoiceTime: voiceTime,
	}

	tmpl, err := template.ParseFiles("templates/layout.gohtml", "templates/index.gohtml")
	if err != nil {
		return err
	}

	err = tmpl.ExecuteTemplate(w, "layout", Layout{
		Title:   "Главная страница",
		Access:  access,
		Content: content,
	})
	if err != nil {
		return err
	}

	return nil
}
