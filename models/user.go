package models

import (
	"time"
)

// Represents a registered user
// Use setter methods for updating values, as they are stored in the database
type User struct {
	ID               string
	Username         string
	RegisteredOn     time.Time
	LastLogin        time.Time
	AvatarURL        string
	AccessToken      string
	RefreshToken     string
	AccessExpiration time.Time
}

// Creates User and saves it to the database
func CreateUser(id string, username string, registeredOn time.Time, lastLogin time.Time, avatarUrl string,
	accessToken string, refreshToken string, accessExpiration time.Time) (User, error) {
	user := User{
		ID:               id,
		Username:         username,
		RegisteredOn:     registeredOn,
		LastLogin:        lastLogin,
		AvatarURL:        avatarUrl,
		AccessToken:      accessToken,
		RefreshToken:     refreshToken,
		AccessExpiration: accessExpiration,
	}
	_, err := db.Exec(`INSERT INTO web_users(ID, Username, registered_on, last_login, avatar_url, access_token, 
                      refresh_token, access_expiration) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`, id, username,
		registeredOn.Format("2006-01-02 15:04:05"), lastLogin.Format("2006-01-02 15:04:05"),
		avatarUrl, accessToken, refreshToken, accessExpiration.Format("2006-01-02 15:04:05"))
	if err != nil {
		return User{}, err
	}

	return user, nil
}

// Gets User with the specified ID
func GetUser(id string) (User, error) {
	result := User{}
	err := db.QueryRow("SELECT * FROM web_users WHERE ID = ?", id).Scan(&result.ID, &result.Username,
		&result.RegisteredOn, &result.LastLogin, &result.AvatarURL, &result.AccessToken, &result.RefreshToken,
		&result.AccessExpiration)
	if err != nil {
		return User{}, err
	}

	return result, nil
}

// Removes User with the specified ID.
func DeleteUser(id string) error {
	_, err := db.Exec("DELETE FROM web_users WHERE ID = ?", id)
	if err != nil {
		return err
	}

	return nil
}

func (user *User) GetIPs() ([]IP, error) {
	result, err := GetIPs(user.ID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (user *User) AddIP(ip string, lastUsed time.Time) error {
	_, err := CreateIP(user.ID, ip, lastUsed)
	if err != nil {
		return err
	}

	return nil
}

func (user *User) GetXboxes() ([]Xbox, error) {
	result, err := GetXboxes(user.ID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (user *User) AddXbox(xbox string, lastUsed time.Time) error {
	_, err := CreateXbox(user.ID, xbox, lastUsed)
	if err != nil {
		return err
	}

	return nil
}

func (user *User) SetUsername(username string) error {
	_, err := db.Exec("UPDATE web_users SET Username = ? WHERE ID = ?", username, user.ID)
	if err != nil {
		return err
	}

	user.Username = username
	return nil
}

func (user *User) SetRegistration(registration time.Time) error {
	_, err := db.Exec("UPDATE web_users SET registered_on = ? WHERE ID = ?",
		registration.Format("2006-01-02 15:04:05"), user.ID)
	if err != nil {
		return err
	}

	user.RegisteredOn = registration
	return nil
}

func (user *User) SetLastLogin(lastLogin time.Time) error {
	_, err := db.Exec("UPDATE web_users SET last_login = ? WHERE ID = ?",
		lastLogin.Format("2006-01-02 15:04:05"), user.ID)
	if err != nil {
		return err
	}

	user.RegisteredOn = lastLogin
	return nil
}

func (user *User) SetAvatar(avatar string) error {
	_, err := db.Exec("UPDATE web_users SET avatar_url = ? WHERE ID = ?", avatar, user.ID)
	if err != nil {
		return err
	}

	user.AvatarURL = avatar
	return nil
}

func (user *User) SetAccessToken(accessToken string) error {
	_, err := db.Exec("UPDATE web_users SET access_token = ? WHERE ID = ?", accessToken, user.ID)
	if err != nil {
		return err
	}

	user.AccessToken = accessToken
	return nil
}

func (user *User) SetRefreshToken(refreshToken string) error {
	_, err := db.Exec("UPDATE web_users SET refresh_token = ? WHERE ID = ?", refreshToken, user.ID)
	if err != nil {
		return err
	}

	user.RefreshToken = refreshToken
	return nil
}

func (user *User) SetAccessExpiration(accessExpiration time.Time) error {
	_, err := db.Exec("UPDATE web_users SET access_expiration = ? WHERE ID = ?",
		accessExpiration.Format("2006-01-02 15:04:05"), user.ID)
	if err != nil {
		return err
	}

	user.AccessExpiration = accessExpiration
	return nil
}
