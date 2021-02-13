package models

import (
	"database/sql"
	"time"
)

type IP struct {
	UserID   string
	IP       string
	LastUsed time.Time
}

// Creates IP and saves it to the database
func CreateIP(userID string, ip string, lastUsed time.Time) (IP, error) {
	ipStruct := IP{
		UserID:   userID,
		IP:       ip,
		LastUsed: lastUsed,
	}
	_, err := db.Exec("INSERT INTO web_ips(user_id, ip, last_used) VALUES (?, ?, ?)", userID, ip, lastUsed)
	if err != nil {
		return IP{}, err
	}

	return ipStruct, nil
}

// Gets IP addresses of the specified User
func GetIPs(userID string) ([]IP, error) {
	var result []IP
	rows, err := db.Query("SELECT * FROM web_ips WHERE user_id = ?")
	if err != nil {
		if err == sql.ErrNoRows {
			return result, nil
		}
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		nextIP := IP{}
		err = rows.Scan(&nextIP.UserID, &nextIP.IP, &nextIP.LastUsed)
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

func (ip *IP) SetLastUsed(lastUsed time.Time) error {
	_, err := db.Exec("UPDATE web_ips SET last_used = ? WHERE ip = ? AND user_id = ?",
		lastUsed.Format("2006-01-02 15:04:05"), ip.IP, ip.UserID)
	if err != nil {
		return err
	}

	ip.LastUsed = lastUsed
	return nil
}
