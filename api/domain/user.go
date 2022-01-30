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
	UUID      string    `json:"uuid"`
	Name      string    `json:"name" validate:"required,gte=2,lt=20"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
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
				return errors.New(fmt.Sprintf("%sは必須です。", value))
			case "gte":
				return errors.New(fmt.Sprintf("%sは%s文字以上が必須です。", value, err.Param()))
			}
		}
	}
	return err
}
