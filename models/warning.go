package models

import (
	"database/sql"
	"time"
)

type Warning struct {
	ID         string
	UserID     string
	Moderator  string
	Reason     string
	Date       time.Time
	LogMessage string
}

// Creates Warning and saves it to the database
func CreateWarning(id string, userID string, moderator string, reason string, date time.Time,
	logMessage string) (Warning, error) {
	_, err := db.Exec("INSERT INTO warnings(id, user, moderator, reason, date, logmessage) VALUES (?, ?, ?, ?, ?, ?)",
		id, userID, moderator, reason, date.Format("2006-01-02 15:04:05"), logMessage)
	if err != nil {
		return Warning{}, err
	}

	return Warning{
		ID:         id,
		UserID:     userID,
		Moderator:  moderator,
		Reason:     reason,
		Date:       date,
		LogMessage: logMessage,
	}, nil
}

// Gets Warning with the specified ID
func GetWarning(id string) (Warning, error) {
	result := Warning{}
	err := db.QueryRow("SELECT * FROM warnings WHERE id = ?", id).Scan(&result.ID, &result.UserID, &result.Moderator,
		&result.Reason, &result.Date, &result.LogMessage)
	if err != nil {
		return Warning{}, err
	}

	return result, nil
}

// Gets Warnings with the specified user ID
func GetUserWarnings(userID string) ([]Warning, error) {
	var result []Warning
	rows, err := db.Query("SELECT * FROM warnings WHERE user = ? ORDER BY date DESC", userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return result, nil
		}
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		nextWarn := Warning{}
		err = rows.Scan(&nextWarn.ID, &nextWarn.UserID, &nextWarn.Moderator,
			&nextWarn.Reason, &nextWarn.Date, &nextWarn.LogMessage)
		if err != nil {
			return nil, err
		}
		result = append(result, nextWarn)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Deletes Warning with the specified ID
func DeleteWarning(id string) error {
	_, err := db.Exec("DELETE FROM warnings WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}

// Setters to be implemented later
