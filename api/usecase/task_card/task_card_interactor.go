package usecase

import (
	"errors"
	"fmt"
	"log"

	"github.com/kory-jp/react_golang_mux/api/domain"
)

type TaskCardInteractor struct {
	TaskCardRepository TaskCardRepository
}

type TaskCardMessage struct {
	Message string
}

func (interactor *TaskCardInteractor) Add(t domain.TaskCard) (mess *TaskCardMessage, err error) {
	if err = t.TaskCardValidate(); err == nil {
		err = interactor.TaskCardRepository.Store(t)
		if err != nil {
			fmt.Println(err)
			log.Println(err)
			err = errors.New("保存に失敗しました")
			return nil, err
		}
		mess = &TaskCardMessage{
			Message: "保存しました",
		}
		return mess, nil
	}
	return nil, err
}
