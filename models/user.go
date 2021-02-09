package models

import "time"

type User struct {
	id               string
	username         string
	registeredOn     time.Time
	lastLogin        time.Time
	avatarURL        string
	ips              []string
	xboxNames        []string
	accessToken      string
	refreshToken     string
	accessExpiration time.Time
}

func (user *User) GetID() string {
	return user.id
}

func (user *User) GetUsername() string {
	return user.username
}

func (user *User) SetUsername(username string) error {
	_, err := db.Exec("UPDATE web_users SET username = ? WHERE id = ?", username, user.id)
	if err != nil {
		return err
	}

	user.username = username
	return nil
}

func (user *User) GetRegistration() time.Time {
	return user.registeredOn
}

func (user *User) GetLastLogin() time.Time {
	return user.lastLogin
}

func (user *User) GetAvatar() string {
	return user.avatarURL
}

func (user *User) GetIPs() []string {
	return user.ips
}

func (user *User) GetXboxes() []string {
	return user.xboxNames
}

func (user *User) GetAccessToken() string {
	return user.accessToken
}

func (user *User) GetRefreshToken() string {
	return user.refreshToken
}

func (user *User) GetAccessExpiration() time.Time {
	return user.accessExpiration
}
