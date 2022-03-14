package domain

import (
	"fmt"
	"image"
	"time"

	"github.com/go-playground/validator/v10"
)

type Todo struct {
	ID         int         `json:"id"`
	UserID     int         `json:"userId" validate:"required"`
	Title      string      `json:"title" validate:"required,gte=1,lt=50"`
	Content    string      `json:"content" validate:"max=2000"`
	Image      image.Image `json:"image"`
	ImagePath  string      `json:"imagePath"`
	IsFinished bool        `json:"isFinished"`
	CreatedAt  time.Time   `json:"createdAt"`
}

type Todos []Todo

func traslateTodosField(field string) (value string) {
	switch field {
	case "UserID":
		value = "ユーザーID"
	case "Title":
		value = "タイトル"
	case "Content":
		value = "メモ"
	}
	return
}

func (todo *Todo) TodoValidate() (err error) {
	validate := validator.New()
	err = validate.Struct(todo)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			value := traslateTodosField(err.Field())
			switch err.ActualTag() {
			case "required":
				return fmt.Errorf("%sは必須です。", value)
			case "gte":
				return fmt.Errorf("%sは%s文字以上が必須です。", value, err.Param())
			case "lt":
				return fmt.Errorf("%sは%s文字未満の入力になります。", value, err.Param())
			case "max":
				return fmt.Errorf("%sは%s文字を超えて入力はできません。", value, err.Param())
			}
		}
	}
	return err
}
