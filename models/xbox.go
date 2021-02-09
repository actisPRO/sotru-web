package models

import "time"

type Xbox struct {
	UserID   string
	Xbox     string
	LastUsed time.Time
}
