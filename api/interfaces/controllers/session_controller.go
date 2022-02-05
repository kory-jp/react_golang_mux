package controllers

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/kory-jp/react_golang_mux/api/domain"
	"github.com/kory-jp/react_golang_mux/api/interfaces/database"
	usecase "github.com/kory-jp/react_golang_mux/api/usecase/session"

	"github.com/gorilla/sessions"
)

type SessionController struct {
	Interactor usecase.SessionInteractor
}

type SessionValidError struct {
	Detail string
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

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

func (controller *SessionController) Login(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")
	bytesUser, err := io.ReadAll(r.Body)
	if err != nil {
		log.SetFlags(log.Llongfile)
		log.Println(err)
	}
	userType := new(domain.User)
	if err := json.Unmarshal(bytesUser, userType); err != nil {
		log.SetFlags(log.Llongfile)
		log.Println(err)
		return
	}
	user, err := controller.Interactor.Login(*userType)
	if err != nil {
		log.SetFlags(log.Llongfile)
		log.Println(err)
		validErr := &SessionValidError{Detail: err.Error()}
		e, _ := json.Marshal(validErr)
		fmt.Fprintln(w, string(e))
	} else {
		token, _ := MakeRandomStr(10)
		session.Values["token"] = token
		cookie := http.Cookie{
			Name:     "cookie-name",
			Value:    token,
			HttpOnly: true,
		}
		jsonUser, err := json.Marshal(user)
		if err != nil {
			log.SetFlags(log.Llongfile)
			log.Println(err)
		}
		http.SetCookie(w, &cookie)
		fmt.Fprintln(w, string(jsonUser))
	}
}

func (controller *SessionController) Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")
	session.Values["token"] = nil
	cookie := http.Cookie{
		Name:     "cookie-name",
		Value:    "",
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
}

func MakeRandomStr(digit uint32) (string, error) {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	b := make([]byte, digit)
	if _, err := rand.Read(b); err != nil {
		log.SetFlags(log.Llongfile)
		log.Println(err)
	}

	var result string
	for _, v := range b {
		result += string(letters[int(v)%len(letters)])
	}
	return result, nil
}
