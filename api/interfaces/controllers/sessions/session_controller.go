package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/kory-jp/react_golang_mux/api/domain"
	"github.com/kory-jp/react_golang_mux/api/interfaces/database"
	sessionRepo "github.com/kory-jp/react_golang_mux/api/interfaces/database/sessions"
	usecase "github.com/kory-jp/react_golang_mux/api/usecase/sessions"
)

type SessionController struct {
	Interactor usecase.SessionInteractor
}

type Response struct {
	Status  int          `json:"status"`
	Message string       `json:"message"`
	User    *domain.User `json:"user"`
}

func (res *Response) SetResp(status int, mess string, user *domain.User) (resStr string) {
	response := &Response{status, mess, user}
	r, _ := json.Marshal(response)
	resStr = string(r)
	return
}

func NewSessionController(sqlHandler database.SqlHandler) *SessionController {
	return &SessionController{
		Interactor: usecase.SessionInteractor{
			SessionRepository: &sessionRepo.SessionRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

var Store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

func (controller *SessionController) Login(w http.ResponseWriter, r *http.Request) {
	if r.ContentLength == 0 {
		fmt.Println("NO DATA BODY")
		log.Println("NO DATA BODY")
		resStr := new(Response).SetResp(400, "データ取得に失敗しました", nil)
		fmt.Fprintln(w, resStr)
		return
	}
	bytesUser, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		resStr := new(Response).SetResp(400, "データ取得に失敗しました", nil)
		fmt.Fprintln(w, resStr)
		return
	}
	userType := new(domain.User)
	if err := json.Unmarshal(bytesUser, userType); err != nil {
		fmt.Println(err)
		log.Println(err)
		resStr := new(Response).SetResp(400, "データ取得に失敗しました", nil)
		fmt.Fprintln(w, resStr)
		return
	}
	user, err := controller.Interactor.Login(*userType)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		resStr := new(Response).SetResp(400, err.Error(), nil)
		fmt.Fprintln(w, resStr)
		return
	} else {
		session, _ := Store.New(r, "session")
		session.Values["userId"] = user.ID
		resStr := new(Response).SetResp(200, "ログインしました", user)
		session.Save(r, w)
		fmt.Fprintln(w, resStr)
	}
}

func (controller *SessionController) Authenticate(w http.ResponseWriter, r *http.Request) {
	session, err := Store.Get(r, "session")
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		resStr := new(Response).SetResp(400, "データ取得に失敗しました", nil)
		fmt.Fprintln(w, resStr)
		return
	}
	if userId, ok := session.Values["userId"].(int); ok {
		user, err := controller.Interactor.IsLoggedin(userId)
		if err != nil {
			resStr := new(Response).SetResp(401, err.Error(), nil)
			fmt.Fprintln(w, resStr)
			return
		}
		session.Values["userId"] = user.ID
		session.Save(r, w)
		resStr := new(Response).SetResp(200, "ログイン状態確認完了", user)
		fmt.Fprintln(w, resStr)
	} else {
		resStr := new(Response).SetResp(401, "認証に失敗しました", nil)
		fmt.Fprintln(w, resStr)
	}
}

func (controller *SessionController) Logout(w http.ResponseWriter, r *http.Request) {
	session, err := Store.Get(r, "session")
	if err != nil || session.Values["userId"] == 0 {
		fmt.Println(err)
		log.Println(err)
		resStr := new(Response).SetResp(400, "データ取得に失敗しました", nil)
		fmt.Fprintln(w, resStr)
		return
	}
	session.Values["userId"] = nil
	session.Save(r, w)
	resStr := new(Response).SetResp(200, "ログアウトしました", nil)
	fmt.Fprintln(w, resStr)
}
