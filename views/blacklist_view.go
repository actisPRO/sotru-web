package views

import (
	"html/template"
	"net/http"
	"sotru-web/models"
)

func ExecuteBlacklist(entries []models.BlacklistEntry, w http.ResponseWriter, r *http.Request) error {
	tmpl, err := template.ParseFiles("templates/layout.gohtml", "templates/blacklist.gohtml", "templates/navbar.gohtml")
	if err != nil {
		return err
	}

	renderEntries := PrepareBlacklistEntries(entries)

	err = tmpl.ExecuteTemplate(w, "layout", Layout{
		Title:   "Чёрный список",
		Page:    "blacklist",
		Access:  0,
		Content: renderEntries,
	})
	if err != nil {
		return err
	}

	return nil
}
