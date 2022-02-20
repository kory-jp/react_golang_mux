package domain

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
)

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name" validate:"required,gte=2,lt=20"`
	Email     string    `json:"email" validate:"required,email"`
	Password  string    `json:"password" validate:"required,gte=5,lt=20"`
	CreatedAt time.Time `json:"createdAt"`
}

type Users []User

func (u User) Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}

func traslateField(field string) (value string) {
	switch field {
	case "Name":
		value = "名前"
	case "Email":
		value = "メールアドレス"
	case "Password":
		value = "パスワード"
	}
	return
}

func UserValidate(user *User) (err error) {
	validate := validator.New()
	err = validate.Struct(user)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			value := traslateField(err.Field())
			switch err.ActualTag() {
			case "required":
				return fmt.Errorf("%sは必須です。", value)
			case "gte":
				return fmt.Errorf("%sは%s文字以上が必須です。", value, err.Param())
			case "lt":
				return fmt.Errorf("%sは%s文字以内の入力になります。", value, err.Param())
			case "email":
				return errors.New("メールアドレスのフォーマットに誤りがあります")
			}
		}
	}
	return err
}
