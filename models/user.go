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

func (user *User) SetRegistration(registration time.Time) error {
	_, err := db.Exec("UPDATE web_users SET registered_on = ? WHERE id = ?",
		registration.Format("2006-01-02 15:04:05"), user.id)
	if err != nil {
		return err
	}

	user.registeredOn = registration
	return nil
}

func (user *User) GetLastLogin() time.Time {
	return user.lastLogin
}

func (user *User) SetLastLogin(lastLogin time.Time) error {
	_, err := db.Exec("UPDATE web_users SET last_login = ? WHERE id = ?",
		lastLogin.Format("2006-01-02 15:04:05"), user.id)
	if err != nil {
		return err
	}

	user.registeredOn = lastLogin
	return nil
}

func (user *User) GetAvatar() string {
	return user.avatarURL
}

func (user *User) SetAvatar(avatar string) error {
	_, err := db.Exec("UPDATE web_users SET avatar_url = ? WHERE id = ?", avatar, user.id)
	if err != nil {
		return err
	}

	user.avatarURL = avatar
	return nil
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

func (user *User) SetAccessToken(accessToken string) error {
	_, err := db.Exec("UPDATE web_users SET access_token = ? WHERE id = ?", accessToken, user.id)
	if err != nil {
		return err
	}

	user.accessToken = accessToken
	return nil
}

func (user *User) GetRefreshToken() string {
	return user.refreshToken
}

func (user *User) SetRefreshToken(refreshToken string) error {
	_, err := db.Exec("UPDATE web_users SET refresh_token = ? WHERE id = ?", refreshToken, user.id)
	if err != nil {
		return err
	}

	user.refreshToken = refreshToken
	return nil
}

func (user *User) GetAccessExpiration() time.Time {
	return user.accessExpiration
}

func (user *User) SetAccessExpiration(accessExpiration time.Time) error {
	_, err := db.Exec("UPDATE web_users SET access_expiration = ? WHERE id = ?",
		accessExpiration.Format("2006-01-02 15:04:05"), user.id)
	if err != nil {
		return err
	}

	user.accessExpiration = accessExpiration
	return nil
}