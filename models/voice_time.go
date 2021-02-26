package models

import (
	"fmt"
	"time"
)

type VoiceTime struct {
	UserID  string
	Seconds uint64
}

func GetVoiceTime(id string) (VoiceTime, error) {
	result := VoiceTime{}
	err := db.QueryRow("SELECT * FROM voice_times WHERE user_id = ?", id).Scan(&result.UserID, &result.Seconds)
	if err != nil {
		return VoiceTime{}, err
	}

	return result, nil
}

func (voiceTime *VoiceTime) Time() time.Duration {
	str := fmt.Sprintf("%ds", voiceTime.Seconds)
	result, _ := time.ParseDuration(str)
	return result
}
