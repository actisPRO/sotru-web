package cache

import (
	"github.com/bwmarrin/discordgo"
	"github.com/garyburd/redigo/redis"
)

var (
	Connection redis.Conn
	bot        *discordgo.Session
	guild      string
)

func UseDiscord(guildID string, session *discordgo.Session) {
	guild = guildID
	bot = session
}
