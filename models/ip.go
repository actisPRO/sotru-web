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
	rows, err := db.Query("SELECT * FROM web_ips WHERE user_id = ?", userID)
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
		result = append(result, nextIP)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Gets IP addresses of the blacklisted users
func GetBlacklistedIPs() ([]IP, error) {
	var result []IP
	rows, err := db.Query(
		`SELECT web_ips.ip
			   FROM blacklist
    			    JOIN web_users ON blacklist.discord_id = web_users.id
    			    JOIN web_ips   ON web_users.id = web_ips.user_id;`)
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
		result = append(result, nextIP)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Checks if the specified IP is blacklisted
func IsBlacklisted(ip string) (bool, error) {
	ips, err := GetBlacklistedIPs()
	if err != nil {
		return false, err
	}

	for i := 0; i < len(ips); i++ {
		if ip == ips[i].IP {
			return true, nil
		}
	}

	return false, nil
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
