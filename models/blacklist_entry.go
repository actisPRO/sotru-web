package models

import (
	"time"
)

type BlacklistEntry struct {
	ID          string
	DiscordID   string
	DiscordName string
	XboxTag     string
	BanDate     time.Time
	ModeratorID string
	Reason      string
	Additional  string
}

// Creates new BlacklistEntry and saves it to the database
func CreateBlacklistEntry(id string, discordID string, discordName string, xbox string, banDate time.Time,
	moderatorID string, reason string, additional string) (BlacklistEntry, error) {
	_, err :=
		db.Exec("INSERT INTO blacklist(id, discord_id, discord_username, xbox, ban_date, moderator_id, reason, additional) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
			id, discordID, discordName, xbox, banDate.Format("2006-01-02"), moderatorID, reason, additional)
	if err != nil {
		return BlacklistEntry{}, err
	}

	return BlacklistEntry{
		ID:          id,
		DiscordID:   discordID,
		DiscordName: discordName,
		XboxTag:     xbox,
		BanDate:     banDate,
		ModeratorID: moderatorID,
		Reason:      reason,
		Additional:  additional,
	}, nil
}

// Gets BlacklistEntry with the specified ID
func GetBlacklistEntry(id string) (BlacklistEntry, error) {
	result := BlacklistEntry{}
	err := db.QueryRow("SELECT * FROM blacklist WHERE id = ?", id).Scan(&result.ID, &result.DiscordID,
		&result.DiscordName, &result.XboxTag, &result.BanDate, &result.ModeratorID, &result.Reason, &result.Additional)
	if err != nil {
		return BlacklistEntry{}, err
	}

	return result, nil
}

// Deletes BlacklistEntry with the specified ID
func DeleteBlacklistEntry(id string) error {
	_, err := db.Exec("DELETE FROM blacklist WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}

// Checks if Discord ID - Xbox pair is blacklisted.
func IsBlacklisted(discordID string, xbox string) bool {
	result := BlacklistEntry{}
	err := db.QueryRow("SELECT * FROM blacklist WHERE discord_id = ? AND xbox = ?", discordID, xbox).
		Scan(&result.ID, &result.DiscordID,
			&result.DiscordName, &result.XboxTag, &result.BanDate, &result.ModeratorID, &result.Reason, &result.Additional)
	if err != nil || result.ID == "" {
		return false
	}

	return true
}

func (entry *BlacklistEntry) SetDiscordID(value string) error {
	_, err := db.Exec("UPDATE blacklist SET discord_id = ? WHERE id = ?", value, entry.ID)
	if err != nil {
		return err
	}

	entry.DiscordID = value
	return nil
}

func (entry *BlacklistEntry) SetDiscordName(value string) error {
	_, err := db.Exec("UPDATE blacklist SET discord_username = ? WHERE id = ?", value, entry.ID)
	if err != nil {
		return err
	}

	entry.DiscordName = value
	return nil
}

func (entry *BlacklistEntry) SetXboxTag(value string) error {
	_, err := db.Exec("UPDATE blacklist SET xbox = ? WHERE id = ?", value, entry.ID)
	if err != nil {
		return err
	}

	entry.XboxTag = value
	return nil
}

func (entry *BlacklistEntry) SetBanDate(value time.Time) error {
	_, err := db.Exec("UPDATE blacklist SET ban_date = ? WHERE id = ?", value.Format("2006-01-02"), entry.ID)
	if err != nil {
		return err
	}

	entry.BanDate = value
	return nil
}

func (entry *BlacklistEntry) SetModeratorID(value string) error {
	_, err := db.Exec("UPDATE blacklist SET moderator_id = ? WHERE id = ?", value, entry.ID)
	if err != nil {
		return err
	}

	entry.ModeratorID = value
	return nil
}

func (entry *BlacklistEntry) SetReason(value string) error {
	_, err := db.Exec("UPDATE blacklist SET reason = ? WHERE id = ?", value, entry.ID)
	if err != nil {
		return err
	}

	entry.Reason = value
	return nil
}

func (entry *BlacklistEntry) SetAdditional(value string) error {
	_, err := db.Exec("UPDATE blacklist SET additional = ? WHERE id = ?", value, entry.ID)
	if err != nil {
		return err
	}

	entry.Additional = value
	return nil
}
