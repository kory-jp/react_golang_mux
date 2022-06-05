package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	controllers "github.com/kory-jp/react_golang_mux/api/interfaces/controllers/sessions"
	"github.com/kory-jp/react_golang_mux/api/interfaces/database"
	tags "github.com/kory-jp/react_golang_mux/api/interfaces/database/tags"
	usecase "github.com/kory-jp/react_golang_mux/api/usecase/tags"

	"github.com/kory-jp/react_golang_mux/api/domain"
)

type TagController struct {
	Interactor usecase.TagInteractor
}

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Tags    domain.Tags `json:"tags"`
}

func (res *Response) SetResp(status int, mess string, tags domain.Tags) (resStr string) {
	response := &Response{status, mess, tags}
	r, _ := json.Marshal(response)
	resStr = string(r)
	return
}

func NewTagController(sqlHandler database.SqlHandler) *TagController {
	return &TagController{
		Interactor: usecase.TagInteractor{
			TagRepository: &tags.TagRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func GetUserId(r *http.Request) (userId int, err error) {
	session, err := controllers.Store.Get(r, "session")
	if err != nil {
		log.Println(err)
		fmt.Println(err)
		return 0, err
	}
	userId = session.Values["userId"].(int)
	return userId, nil
}

func (controller *TagController) Index(w http.ResponseWriter, r *http.Request) {
	userId, err := GetUserId(r)
	if err != nil || userId == 0 {
		fmt.Println(err)
		log.Println(err)
		resStr := new(Response).SetResp(401, "ログインをしてください", nil)
		fmt.Fprintln(w, resStr)
		return
	}

	tags, err := controller.Interactor.Tags()
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		resStr := new(Response).SetResp(400, err.Error(), nil)
		fmt.Fprintln(w, resStr)
		return
	}

	resStr := new(Response).SetResp(200, "タグ一覧取得", tags)
	fmt.Fprintln(w, resStr)
}
