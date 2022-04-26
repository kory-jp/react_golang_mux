package domain

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
)

type TaskCard struct {
	ID         int         `json:"id"`
	UserID     int         `json:"userId" validate:"required"`
	TodoID     json.Number `json:"todoId" validate:"required"`
	Title      string      `json:"title" validate:"required,gte=1,lt=50"`
	Purpose    string      `json:"purpose" validate:"lt=2000"`
	Content    string      `json:"content" validate:"lt=2000"`
	Memo       string      `json:"memo" validate:"lt=2000"`
	IsFinished bool        `json:"isFinished"`
	CreatedAt  time.Time   `json:"createdAt"`
}

type TaskCards []TaskCard

func translateTaskCardsField(field string) (value string) {
	switch field {
	case "UserID":
		value = "ユーザーID"
	case "TodoID":
		value = "TodoID"
	case "Title":
		value = "タイトル"
	case "Purpose":
		value = "目的"
	case "Content":
		value = "作業内容"
	case "Memo":
		value = "メモ"
	}
	return
}

func (taskCard *TaskCard) TaskCardValidate() (err error) {
	validate := validator.New()
	err = validate.Struct(taskCard)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			value := translateTaskCardsField(err.Field())
			switch err.ActualTag() {
			case "required":
				return fmt.Errorf("%sは必須です。", value)
			case "gte":
				return fmt.Errorf("%sは%s文字以上が必須です。", value, err.Param())
			case "lt":
				return fmt.Errorf("%sは%s文字以内の入力になります。", value, err.Param())
			}
		}
	}
	return err
}
