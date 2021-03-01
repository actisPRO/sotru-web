package cache

import (
	"github.com/bwmarrin/discordgo"
)

var (
	bot   *discordgo.Session
	guild string
)

func UseDiscord(guildID string, session *discordgo.Session) {
	guild = guildID
	bot = session
}
