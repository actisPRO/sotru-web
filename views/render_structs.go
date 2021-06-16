package views

import (
	"github.com/google/logger"
	"sotru-web/cache"
	"sotru-web/models"
	"sotru-web/utils"
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
		modInfo, err := cache.GetUserInfo(input[i].Moderator, 600)
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

type RenderBlacklistEntry struct {
	ID          string
	DiscordUser string
	DiscordID   string
	Xbox        string
	Date        string
	Moderator   string
	Reason      string
	Additional  string
	IsLink      bool
	Website     string
}

func PrepareBlacklistEntries(input []models.BlacklistEntry) []RenderBlacklistEntry {
	var result []RenderBlacklistEntry
	for i := 0; i < len(input); i++ {
		logger.Info("Caching for ID ", i)
		user := input[i].DiscordID.String
		userInfo, err := cache.GetUserInfo(user, 3600*4)
		if err == nil {
			user = userInfo.Username
		} else {
			user = input[i].DiscordName.String
		}

		mod := input[i].ModeratorID
		modInfo, err := cache.GetUserInfo(mod, 600)
		if err == nil {
			mod = modInfo.Username
		}

		isLink := utils.IsLink(input[i].Additional.String)
		website := ""
		if isLink {
			website = utils.GetWebsiteName(input[i].Additional.String)
		}

		rbe := RenderBlacklistEntry{
			ID:          input[i].ID,
			DiscordUser: user,
			DiscordID:   input[i].DiscordID.String,
			Xbox:        input[i].XboxTag.String,
			Date:        input[i].BanDate.Format("02.01.2006"),
			Moderator:   mod,
			Reason:      input[i].Reason.String,
			Additional:  input[i].Additional.String,
			IsLink:      isLink,
			Website:     website,
		}

		result = append(result, rbe)
	}

	return result
}
