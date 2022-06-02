package controllers_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	mock_database "github.com/kory-jp/react_golang_mux/api/interfaces/mock"

	usecase "github.com/kory-jp/react_golang_mux/api/usecase/sessions"
	mock_usecase "github.com/kory-jp/react_golang_mux/api/usecase/sessions/mock"

	"github.com/golang/mock/gomock"
	"github.com/kory-jp/react_golang_mux/api/domain"
	controllers "github.com/kory-jp/react_golang_mux/api/interfaces/controllers/sessions"
)

var response *controllers.Response

func TestLogin(t *testing.T) {
	c := gomock.NewController(t)
	defer c.Finish()
	sqlhandler := mock_database.NewMockSqlHandler(c)
	ctrl := controllers.NewSessionController(sqlhandler)
	SessionRepository := mock_usecase.NewMockSessionRepository(c)
	inter := &usecase.SessionInteractor{}
	inter.SessionRepository = SessionRepository
	ctrl.Interactor = *inter

	cases := []struct {
		name          string
		args          domain.User
		requestBody   bool
		prepareMockFn func(m *mock_usecase.MockSessionRepository, args domain.User)
		response      controllers.Response
		haveCookie    bool
	}{
		{
			name: "login = success",
			args: domain.User{
				Email:    "test@exm.com",
				Password: "testPassword",
			},
			requestBody: true,
			prepareMockFn: func(m *mock_usecase.MockSessionRepository, args domain.User) {
				pUser := &domain.User{
					Email: args.Email,
					// --- 戻り値のuser.PasswordをEncryptメソッドで暗号化した文字列に変換 ---
					// --- 引数で渡たすargs.Password(暗号化前のuser.Passwordと同一文字列)と
					// 暗号化したuser.PasswordをCompareHashAndPasswordメソッドを用いて比較 ---
					Password: args.Encrypt(args.Password),
				}
				m.EXPECT().FindByEmail(args).Return(pUser, nil)
			},
			response: controllers.Response{
				Status:  200,
				Message: "ログインしました",
			},
			haveCookie: true,
		},
		{
			name: "when requestBody = nil, login = fail",
			args: domain.User{
				Email:    "test@exm.com",
				Password: "testPassword",
			},
			requestBody: false,
			prepareMockFn: func(m *mock_usecase.MockSessionRepository, args domain.User) {
				pUser := &domain.User{
					Email:    args.Email,
					Password: args.Encrypt(args.Password),
				}
				m.EXPECT().FindByEmail(args).Return(pUser, nil).AnyTimes()
			},
			response: controllers.Response{
				Status:  400,
				Message: "データ取得に失敗しました",
			},
			haveCookie: true,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			var req *http.Request
			tt.prepareMockFn(SessionRepository, tt.args)
			jsonArgs, _ := json.Marshal(tt.args)
			apiURL := "/api/login"
			if tt.requestBody {
				req = httptest.NewRequest("POST", apiURL, bytes.NewBuffer(jsonArgs))
			} else {
				req = httptest.NewRequest("POST", apiURL, nil)
			}
			w := httptest.NewRecorder()
			ctrl.Login(w, req)
			buf, _ := ioutil.ReadAll(w.Body)
			json.Unmarshal(buf, &response)
			assert.Equal(t, tt.response.Status, response.Status)
			assert.Equal(t, tt.response.Message, response.Message)
		})
	}
}

func TestAuthenticate(t *testing.T) {
	c := gomock.NewController(t)
	defer c.Finish()
	sqlhandler := mock_database.NewMockSqlHandler(c)
	ctrl := controllers.NewSessionController(sqlhandler)
	SessionRepository := mock_usecase.NewMockSessionRepository(c)
	inter := &usecase.SessionInteractor{}
	inter.SessionRepository = SessionRepository
	ctrl.Interactor = *inter

	cases := []struct {
		name          string
		userId        int
		isSession     bool
		prepareMockFn func(m *mock_usecase.MockSessionRepository, userId int)
		response      controllers.Response
		haveCookie    bool
	}{
		{
			name:      "authentication = success",
			userId:    1,
			isSession: true,
			prepareMockFn: func(m *mock_usecase.MockSessionRepository, userId int) {
				pUser := &domain.User{
					ID: userId,
				}
				m.EXPECT().FindById(userId).Return(pUser, nil)
			},
			response: controllers.Response{
				Status:  200,
				Message: "ログイン状態確認完了",
			},
			haveCookie: true,
		},
		{
			name:      "when userId = 0, authentication = fail",
			userId:    0,
			isSession: true,
			prepareMockFn: func(m *mock_usecase.MockSessionRepository, userId int) {
				pUser := &domain.User{
					ID: userId,
				}
				m.EXPECT().FindById(userId).Return(pUser, nil).AnyTimes()
			},
			response: controllers.Response{
				Status:  401,
				Message: "認証に失敗しました",
			},
			haveCookie: false,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepareMockFn(SessionRepository, tt.userId)
			apiURL := "/api/authenticate"
			req := httptest.NewRequest("GET", apiURL, nil)
			w := httptest.NewRecorder()
			session, err := controllers.Store.Get(req, "session")
			if err != nil {
				t.Error(err)
				return
			}
			session.Values["userId"] = tt.userId
			err = session.Save(req, w)
			if err != nil {
				t.Error(err)
				return
			}
			ctrl.Authenticate(w, req)
			buf, _ := ioutil.ReadAll(w.Body)
			json.Unmarshal(buf, &response)
			assert.Equal(t, tt.response.Status, response.Status)
			assert.Equal(t, tt.response.Message, response.Message)
		})
	}
}

func TestLogout(t *testing.T) {
	c := gomock.NewController(t)
	defer c.Finish()
	sqlhandler := mock_database.NewMockSqlHandler(c)
	ctrl := controllers.NewSessionController(sqlhandler)
	SessionRepository := mock_usecase.NewMockSessionRepository(c)
	inter := &usecase.SessionInteractor{}
	inter.SessionRepository = SessionRepository
	ctrl.Interactor = *inter

	cases := []struct {
		name          string
		userId        int
		prepareMockFn func(m *mock_usecase.MockSessionRepository, userId int)
		response      controllers.Response
	}{
		{
			name:   "when session.userId = userId, result = logout",
			userId: 1,
			prepareMockFn: func(m *mock_usecase.MockSessionRepository, userId int) {
				pUser := &domain.User{
					ID: userId,
				}
				m.EXPECT().FindById(userId).Return(pUser, nil).AnyTimes()
			},
			response: controllers.Response{
				Status:  200,
				Message: "ログアウトしました",
			},
		},
		{
			name:   "when session.userId = nil, logout = fail",
			userId: 0,
			prepareMockFn: func(m *mock_usecase.MockSessionRepository, userId int) {
				pUser := &domain.User{
					ID: userId,
				}
				m.EXPECT().FindById(userId).Return(pUser, nil).AnyTimes()
			},
			response: controllers.Response{
				Status:  400,
				Message: "データ取得に失敗しました",
			},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepareMockFn(SessionRepository, tt.userId)
			apiURL := "/api/logout"
			req := httptest.NewRequest("DELETE", apiURL, nil)
			w := httptest.NewRecorder()
			session, err := controllers.Store.Get(req, "session")
			if err != nil {
				t.Error(err)
				return
			}
			session.Values["userId"] = tt.userId
			err = session.Save(req, w)
			if err != nil {
				t.Error(err)
				return
			}
			ctrl.Logout(w, req)
			buf, _ := ioutil.ReadAll(w.Body)
			json.Unmarshal(buf, &response)
			assert.Equal(t, tt.response.Status, response.Status)
			assert.Equal(t, tt.response.Message, response.Message)
		})
	}
}
