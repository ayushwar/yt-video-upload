package models

import "time"

type Video struct {
	ID             int
	Title          string
	Description    string
	YouTubeVideoID string
	CreatedAt      time.Time
}