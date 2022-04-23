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
	Content    string      `json:"content" validate:"lt=2000"`
	Image      image.Image `json:"image"`
	ImagePath  string      `json:"imagePath"`
	IsFinished bool        `json:"isFinished"`
	Importance int         `json:"importance" validate:"required,min=1,max=3"`
	Urgency    int         `json:"urgency"  validate:"required,min=1,max=3"`
	CreatedAt  time.Time   `json:"createdAt"`
	Tags       Tags        `json:"tags"`
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
	case "Importance":
		value = "重要度"
	case "Urgency":
		value = "緊急度"
	}
	return
}

// 文字数バリデーションには"lt","get"を使用
// 数値範囲バリデーションには"min","max"を使用
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
			case "min":
				return fmt.Errorf("%sに異常な値が入力されました。", value)
			case "max":
				return fmt.Errorf("%sに異常な値が入力されました", value)
			}
		}
	}
	return err
}
