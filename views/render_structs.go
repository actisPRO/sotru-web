package views

import (
	"sotru-web/cache"
	"sotru-web/models"
)

// Represents warning, but all the fields, required for rendering pages are correctly formatted strings.
// Use PrepareWarnings method
type RenderWarning struct {
	ID        string
	Date      string
	Moderator string
	Reason    string
}

// Converts all models.Warning in the slice to RenderWarning
// Formats date and gets moderator from cache
func PrepareWarnings(input []models.Warning) []RenderWarning {
	var result []RenderWarning

	for i := 0; i < len(input); i++ {
		mod := input[i].Moderator
		modInfo, err := cache.GetUserInfo(input[i].Moderator)
		if err == nil {
			mod = modInfo.Username
		}

		rw := RenderWarning{
			ID:        input[i].ID,
			Date:      input[i].Date.Format("15:04:05 02.01.2006"),
			Moderator: mod,
			Reason:    input[i].Reason,
		}
		result = append(result, rw)
	}

	return result
}
