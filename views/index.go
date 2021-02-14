package views

import (
	"html/template"
	"net/http"
	"sotru-web/models"
)

type IndexContent struct {
	Username string
	Avatar   string
}

// Executes layout for the index page. If user is not authorized, use empty struct
func ExecuteIndex(user models.User, w http.ResponseWriter) error {
	access := 0

	// check if user is empty
	if user.ID != "" {
		access = user.GetAccess()
	}

	content := IndexContent{
		Username: user.Username,
		Avatar:   user.AvatarURL,
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
