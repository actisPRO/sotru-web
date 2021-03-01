package views

import (
	"fmt"
	"html/template"
	"net/http"
	"sotru-web/models"
	"sotru-web/utils"
	"time"
)

type IndexContent struct {
	Username   string
	Avatar     string
	JoinedAt   string
	Xbox       string
	VoiceTime  string
	WarnsCount int
	Warnings   []RenderWarning
}

// Executes layout for the index page. If user is not authorized, use empty struct
func ExecuteIndex(user models.User, w http.ResponseWriter, r *http.Request) error {
	access := 0

	// check if user is empty
	if user.ID != "" {
		access = user.GetAccess()
	} else {
		// todo: normal auth page
		w.Header().Set("Content-Type", "")
		http.Redirect(w, r, "/login", http.StatusFound)
		return nil
	}

	// getting all required data
	voiceTime := ""
	vtd, err := user.GetVoiceTime()
	if err == nil {
		voiceTime = utils.FormatDuration(vtd)
	}
	xbox := ""
	xboxes, _ := user.GetXboxes()
	if len(xboxes) > 0 {
		xbox = xboxes[0].Xbox
	}
	joinedAtTime, err := user.GetGuildJoinDate()
	joinedAt := ""
	if err == nil {
		dif := int(time.Now().Sub(joinedAtTime).Milliseconds()) / 1000 / 3600 / 24
		days := utils.FormatUnit(dif, utils.Days)
		joinedAt = fmt.Sprintf("%s (%s)", utils.FormatDateTime(joinedAtTime), days)
	}
	warns, err := user.GetWarnings()
	if err != nil {
		warns = []models.Warning{}
	}

	// Preparing content and rendering
	content := IndexContent{
		Username:   user.Username,
		Avatar:     user.AvatarURL,
		JoinedAt:   joinedAt,
		Xbox:       xbox,
		VoiceTime:  voiceTime,
		WarnsCount: len(warns),
		Warnings:   PrepareWarnings(warns),
	}

	tmpl, err := template.ParseFiles("templates/layout.gohtml", "templates/index.gohtml", "templates/navbar.gohtml")
	if err != nil {
		return err
	}

	err = tmpl.ExecuteTemplate(w, "layout", Layout{
		Title:   "Главная страница",
		Page:    "index",
		Access:  access,
		Content: content,
	})
	if err != nil {
		return err
	}

	return nil
}
