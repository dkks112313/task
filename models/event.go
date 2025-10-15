package models

import "time"

type Metadata struct {
	Path string `json:"page"`
}

type Event struct {
	UserID    uint      `json:"user_id"`
	Action    string    `json:"action"`
	Metadata  Metadata  `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
}
