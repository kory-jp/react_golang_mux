package controllers

import (
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/kory-jp/react_golang_mux/api/domain"
	"github.com/kory-jp/react_golang_mux/api/interfaces/database"
	usecase "github.com/kory-jp/react_golang_mux/api/usecase/session"
)

type SessionController struct {
	Interactor usecase.SessionInteractor
}

type SessionValidError struct {
	Error string
}

func (serr *SessionValidError) MakeErr(mess string) (errStr string) {
	err := errors.New(mess)
	todosErr := &SessionValidError{Error: err.Error()}
	e, _ := json.Marshal(todosErr)
	errStr = string(e)
	return
}

func NewSessionController(sqlHandler database.SqlHandler) *SessionController {
	return &SessionController{
		Interactor: usecase.SessionInteractor{
			SessionRepository: &database.SessionRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

var Store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

func (controller *SessionController) Login(w http.ResponseWriter, r *http.Request) {
	bytesUser, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		errStr := new(SessionValidError).MakeErr("データ取得に失敗しました")
		fmt.Fprintln(w, errStr)
		return
	}
	userType := new(domain.User)
	if err := json.Unmarshal(bytesUser, userType); err != nil {
		fmt.Println(err)
		log.Println(err)
		errStr := new(SessionValidError).MakeErr("データ取得に失敗しました")
		fmt.Fprintln(w, errStr)
		return
	}
	user, err := controller.Interactor.Login(*userType)
	if err != nil {
		errStr := new(SessionValidError).MakeErr(err.Error())
		fmt.Fprintln(w, errStr)
	} else {
		token, err := MakeRandomStr(10)
		if err != nil {
			fmt.Println(err)
			log.Println(err)
			errStr := new(SessionValidError).MakeErr("認証に失敗しました")
			fmt.Fprintln(w, errStr)
			return
		}
		session, _ := Store.New(r, "session")
		session.Values["token"] = token
		session.Values["userId"] = user.ID
		cookie := &http.Cookie{
			Name:     "cookie-name",
			Value:    token,
			HttpOnly: true,
		}
		jsonUser, err := json.Marshal(user)
		if err != nil {
			fmt.Println(err)
			log.Println(err)
			errStr := new(SessionValidError).MakeErr("データ取得に失敗しました")
			fmt.Fprintln(w, errStr)
			return
		}
		session.Save(r, w)
		http.SetCookie(w, cookie)
		fmt.Fprintln(w, string(jsonUser))
	}
}

func (controller *SessionController) Authenticate(w http.ResponseWriter, r *http.Request) {
	session, err := Store.Get(r, "session")
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		errStr := new(SessionValidError).MakeErr("データ取得に失敗しました")
		fmt.Fprintln(w, errStr)
		return
	}
	cookie, err := r.Cookie("cookie-name")
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		errStr := new(SessionValidError).MakeErr("データ取得に失敗しました")
		fmt.Fprintln(w, errStr)
		return
	}
	if session.Values["token"] == cookie.Value {
		if userId, ok := session.Values["userId"].(int); ok {
			user, err := controller.Interactor.IsLoggedin(userId)
			if err != nil {
				errStr := new(SessionValidError).MakeErr(err.Error())
				fmt.Fprintln(w, errStr)
				return
			}
			token, err := MakeRandomStr(10)
			if err != nil {
				fmt.Println(err)
				log.Println(err)
				errStr := new(SessionValidError).MakeErr("認証に失敗しました")
				fmt.Fprintln(w, errStr)
				return
			}
			session.Values["token"] = token
			session.Values["userId"] = user.ID
			cookie := &http.Cookie{
				Name:     "cookie-name",
				Value:    token,
				HttpOnly: true,
			}
			jsonUser, err := json.Marshal(user)
			if err != nil {
				fmt.Println(err)
				log.Println(err)
				errStr := new(SessionValidError).MakeErr("データ取得に失敗しました")
				fmt.Fprintln(w, errStr)
				return
			}
			session.Save(r, w)
			http.SetCookie(w, cookie)
			fmt.Fprintln(w, string(jsonUser))
		}
	} else {
		errStr := new(SessionValidError).MakeErr("ログインしてください")
		fmt.Fprintln(w, errStr)
	}
}

func (controller *SessionController) Logout(w http.ResponseWriter, r *http.Request) {
	session, err := Store.Get(r, "session")
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		errStr := new(SessionValidError).MakeErr("データ取得に失敗しました")
		fmt.Fprintln(w, errStr)
		return
	}
	session.Values["token"] = nil
	cookie := http.Cookie{
		Name:     "cookie-name",
		Value:    "",
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
}

func MakeRandomStr(digit uint32) (token string, err error) {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	b := make([]byte, digit)
	if _, err := rand.Read(b); err != nil {
		log.Println(err)
		return "", err
	}

	for _, v := range b {
		token += string(letters[int(v)%len(letters)])
	}
	return token, nil
}
