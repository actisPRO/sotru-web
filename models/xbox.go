package models

import (
	"database/sql"
	"time"
)

type Xbox struct {
	UserID   string
	Xbox     string
	LastUsed time.Time
}

// Creates Xbox and saves it to the database
func CreateXbox(userID string, xbox string, lastUsed time.Time) (Xbox, error) {
	xboxStruct := Xbox{
		UserID:   userID,
		Xbox:     xbox,
		LastUsed: lastUsed,
	}
	_, err := db.Exec("INSERT INTO web_xboxes(user_id, xbox, last_used) VALUES (?, ?, ?)", userID, xbox, lastUsed)
	if err != nil {
		return Xbox{}, err
	}

	return xboxStruct, nil
}

// Gets Xbox tags of the specified User
func GetXboxes(userID string) ([]Xbox, error) {
	var result []Xbox
	rows, err := db.Query("SELECT * FROM web_xboxes WHERE user_id = ?")
	if err != nil {
		if err == sql.ErrNoRows {
			return result, nil
		}
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		nextXbox := Xbox{}
		err = rows.Scan(&nextXbox.UserID, &nextXbox.Xbox, &nextXbox.LastUsed)
		if err != nil {
			return nil, err
		}
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (xbox *Xbox) SetLastUsed(lastUsed time.Time) error {
	_, err := db.Exec("UPDATE web_xboxes SET last_used = ? WHERE xbox = ? AND user_id = ?",
		lastUsed.Format("2006-01-02 15:04:05"), xbox.Xbox, xbox.UserID)
	if err != nil {
		return err
	}

	xbox.LastUsed = lastUsed
	return nil
}
