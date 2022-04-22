package usecase

import (
	"errors"
	"fmt"
	"log"

	"github.com/kory-jp/react_golang_mux/api/domain"
)

type TagInteractor struct {
	TagRepository TagRepository
}

func (interactor *TagInteractor) Tags() (tags domain.Tags, err error) {
	tags, err = interactor.TagRepository.FindAll()
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		err = errors.New("データ取得に失敗しました")
		return nil, err
	}
	return tags, nil
}
