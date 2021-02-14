package models

import (
	"database/sql"
	"github.com/bwmarrin/discordgo"
	"sotru-web/utils"
)

var (
	config utils.Config
	db     *sql.DB = nil
	bot    *discordgo.Session
	guild  string
)

// Globally sets config for the controllers package
func UseConfig(conf utils.Config) {
	config = conf
}

// Globally sets the database connection for the models package
func UseDB(DBConnection *sql.DB) {
	db = DBConnection
}

// Globally sets the discord session for the models package
func UseDiscord(guildID string, session *discordgo.Session) {
	guild = guildID
	bot = session
}
