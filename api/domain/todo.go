package domain

import (
	"image"
	"time"
)

type Todo struct {
	ID         int         `json:"id"`
	UserID     int         `json:"userId"`
	Title      string      `json:"title"`
	Content    string      `json:"content"`
	Image      image.Image `json:"image"`
	ImagePath  string      `json:"imagePath"`
	IsFinished bool        `json:"isFinished"`
	CreatedAt  time.Time   `json:"createdAt"`
}

type Todos []Todo
