package cache

import (
	"time"
)

type UserInfo struct {
	DiscordID string
	Username  string
	AvatarURL string
	StoredAt  time.Time
}

var userCache = make(map[string]UserInfo)

func GetUserInfo(id string) (UserInfo, error) {
	res, ok := userCache[id]
	if !ok {
		err := loadUserToCache(id)
		if err != nil {
			return UserInfo{}, err
		}
	}

	// refresh cache, if the value was stored more then a minute ago
	if time.Now().Sub(res.StoredAt).Minutes() > 60 {
		_ = loadUserToCache(id)
		res = userCache[id]
	}

	return res, nil
}

func loadUserToCache(id string) error {
	user, err := bot.User(id)
	if err != nil {
		return err
	}

	uInfo := UserInfo{
		DiscordID: user.ID,
		Username:  user.String(),
		AvatarURL: user.AvatarURL(""),
		StoredAt:  time.Now(),
	}

	userCache[id] = uInfo
	return nil
}
