package models

import "time"

// Reminder describes reminder structure
type Reminder struct {
	ID        int           `json:"id"`
	Title     string        `json:"title"`
	Message   string        `json:"message"`
	Duration  time.Duration `json:"duration"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}
