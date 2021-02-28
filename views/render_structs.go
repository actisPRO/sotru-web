package views

import "sotru-web/models"

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
		rw := RenderWarning{
			ID:        input[i].ID,
			Date:      input[i].Date.Format("15:04:05 02.01.2006"),
			Moderator: input[i].Moderator,
			Reason:    input[i].Reason,
		}
		result = append(result, rw)
	}

	return result
}
