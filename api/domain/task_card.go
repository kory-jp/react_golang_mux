package domain

import (
	"time"
)

type TaskCard struct {
	ID         int       `json:"id"`
	TodoID     int       `json:"todoId" validate:"required"`
	Title      string    `json:"title" validate:"required,gte=1,lt=50"`
	Purpose    string    `json:"purpose" validate:"lt=2000"`
	Content    string    `json:"content" validate:"lt=2000"`
	Memo       string    `json:"memo" validate:"lt=2000"`
	IsFinished bool      `json:"isFinished"`
	CreatedAt  time.Time `json:"createdAt"`
}

type TaskCards []TaskCard
