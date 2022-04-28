package domain

import (
	"time"
)

type Tag struct {
	ID        int       `json:"id"`
	Value     string    `json:"value"`
	Label     string    `json:"label"`
	CreatedAt time.Time `json:"createdAt"`
}

type Tags []Tag
