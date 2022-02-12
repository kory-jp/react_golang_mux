package domain

import (
	"image"
	"time"
)

type Todo struct {
	ID         int         `json:"id"`
	UserID     int         `json:"user_id"`
	Title      string      `json:"title"`
	Content    string      `json:"content"`
	Image      image.Image `json:"image"`
	ImagePath  string      `json:"image_path"`
	IsFinished bool        `json:"isFinished"`
	CreatedAt  time.Time   `json:"created_at"`
}

type Todos []Todo
