package views

import (
	"html/template"
	"net/http"
)

func ExecuteBlacklist(w http.ResponseWriter, r *http.Request) error {
	tmpl, err := template.ParseFiles("templates/layout.gohtml", "templates/blacklist.gohtml", "templates/navbar.gohtml")
	if err != nil {
		return err
	}

	err = tmpl.ExecuteTemplate(w, "layout", Layout{
		Title: "Чёрный список",
		Page:  "blacklist",
	})
	if err != nil {
		return err
	}

	return nil
}
