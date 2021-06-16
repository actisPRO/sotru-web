package cache

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

type UserInfo struct {
	DiscordID string
	Username  string
	AvatarURL string
	StoredAt  time.Time
}

// Gets UserInfo by ID from the cache. If there is no data in cache or cached data is older then maxAge, calls loadUserToCache
func GetUserInfo(id string, maxAge int) (UserInfo, error) {
	info := UserInfo{
		DiscordID: id,
	}
	username, err := redis.String(Connection.Do("HGET", id, "username"))
	avatar, err := redis.String(Connection.Do("HGET", id, "avatar"))
	timeStr, err := redis.String(Connection.Do("HGET", id, "stored"))
	if err != nil {
		err = loadUserToCache(id)
		if err != nil {
			return UserInfo{}, err
		}
	}

	info.Username = username
	info.AvatarURL = avatar

	info.StoredAt, err = time.Parse("2006-01-02 15:04:05 -0700", timeStr)
	if err != nil {
		err = loadUserToCache(id)
		if err != nil {
			return info, nil
		}

		return info, nil
	}

	// refresh cache if the value was stored for more then 'maxAge' seconds
	t := time.Now()
	age := t.Sub(info.StoredAt)

	if int(age.Seconds()) > maxAge {
		err = loadUserToCache(id)
		if err != nil {
			return info, nil
		}

		return info, nil
	}

	return info, nil
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

	_, err = Connection.Do("HMSET", id, "username", uInfo.Username, "avatar", uInfo.AvatarURL, "stored",
		uInfo.StoredAt.Format("2006-01-02 15:04:05 -0700"))
	if err != nil {
		return err
	}

	return nil
}
