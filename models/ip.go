package models

import "time"

type IP struct {
	UserID   string
	IP       string
	LastUsed time.Time
}
