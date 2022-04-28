package controllers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
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

func TestMakeRandomStr(t *testing.T) {
	cases := []struct {
		name        string
		argumentInt uint32
		tokenLength int
	}{
		{
			name:        "引数10を渡すと、10文字のtokenが生成",
			argumentInt: 10,
			tokenLength: 10,
		},
		{
			name:        "引数20を渡すと、20文字のtokenが生成",
			argumentInt: 20,
			tokenLength: 20,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			token, _ := controllers.MakeRandomStr(tt.argumentInt)
			reg := fmt.Sprintf("[0-9a-zA-Z]{%d}", tt.tokenLength)
			assert.Len(t, token, tt.tokenLength)
			assert.Regexp(t, reg, token)
		})
	}
}

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
		name            string
		args            domain.User
		withCredentials bool
		requestBody     bool
		prepareMockFn   func(m *mock_usecase.MockSessionRepository, args domain.User)
		response        controllers.Response
		haveCookie      bool
	}{
		{
			name: "必須項目が入力された場合、ログイン成功",
			args: domain.User{
				Email:    "test@exm.com",
				Password: "testPassword",
			},
			withCredentials: true,
			requestBody:     true,
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
				Message: "ログインに成功しました",
			},
			haveCookie: true,
		},
		{
			name: "リクエストにcookieの認証情報が含まれない場合(withCredentials: false)、ログイン失敗",
			args: domain.User{
				Email:    "test@exm.com",
				Password: "testPassword",
			},
			withCredentials: false,
			requestBody:     true,
			prepareMockFn: func(m *mock_usecase.MockSessionRepository, args domain.User) {
				pUser := &domain.User{
					Email:    args.Email,
					Password: args.Encrypt(args.Password),
				}
				m.EXPECT().FindByEmail(args).Return(pUser, nil).AnyTimes()
			},
			response: controllers.Response{
				Status:  401,
				Message: "認証に失敗しました",
			},
			haveCookie: true,
		},
		{
			name: "リクエストボディがnilの場合、ログイン失敗",
			args: domain.User{
				Email:    "test@exm.com",
				Password: "testPassword",
			},
			withCredentials: true,
			requestBody:     false,
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
			if tt.withCredentials {
				req.AddCookie(&http.Cookie{
					Name: "cookie-name",
				})
			}
			ctrl.Login(w, req)
			buf, _ := ioutil.ReadAll(w.Body)
			json.Unmarshal(buf, &response)
			assert.Equal(t, tt.response.Status, response.Status)
			assert.Equal(t, tt.response.Message, response.Message)
			// --- ログインされている場合,tokenを保持ているcookie-nameが設定されている ---
			for _, v := range w.Header()["Set-Cookie"] {
				if strings.Contains(v, "cookie-name") {
					assert.Equal(t, tt.haveCookie, strings.Contains(v, "cookie-name"))
				}
			}
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
		name            string
		userId          int
		isSession       bool
		withCredentials bool
		prepareMockFn   func(m *mock_usecase.MockSessionRepository, userId int)
		response        controllers.Response
		haveCookie      bool
	}{
		{
			name:            "sessionとcookieにuserIdが保存されている場合、userIdに一致するUser情報を取得",
			userId:          1,
			isSession:       true,
			withCredentials: true,
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
			name:            "sessionのuserIdがnilの場合、User情報の取得失敗",
			userId:          0,
			isSession:       true,
			withCredentials: true,
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
		{
			name:            "sessionにtokenが設定されていない場合、User情報の取得失敗",
			userId:          1,
			isSession:       false,
			withCredentials: true,
			prepareMockFn: func(m *mock_usecase.MockSessionRepository, userId int) {
				pUser := &domain.User{
					ID: userId,
				}
				m.EXPECT().FindById(userId).Return(pUser, nil).AnyTimes()
			},
			response: controllers.Response{
				Status:  401,
				Message: "ログインしてください",
			},
			haveCookie: false,
		},
		{
			name:            "cookieにtokenが設定されていない場合、User情報の取得失敗",
			userId:          1,
			isSession:       true,
			withCredentials: false,
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
			token, _ := controllers.MakeRandomStr(10)
			session, err := controllers.Store.Get(req, "session")
			if err != nil {
				t.Error(err)
				return
			}
			session.Values["userId"] = tt.userId
			if tt.isSession {
				session.Values["token"] = token
			}
			err = session.Save(req, w)
			if err != nil {
				t.Error(err)
				return
			}
			if tt.withCredentials {
				req.AddCookie(&http.Cookie{
					Name:  "cookie-name",
					Value: token,
				})
			}
			ctrl.Authenticate(w, req)
			buf, _ := ioutil.ReadAll(w.Body)
			json.Unmarshal(buf, &response)
			assert.Equal(t, tt.response.Status, response.Status)
			assert.Equal(t, tt.response.Message, response.Message)
			// --- ログインされている場合,tokenを保持ているcookie-nameが設定されている ---
			for _, v := range w.Header()["Set-Cookie"] {
				if strings.Contains(v, "cookie-name") {
					assert.Equal(t, tt.haveCookie, strings.Contains(v, "cookie-name"))
				}
			}
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
		name            string
		userId          int
		withCredentials bool
		prepareMockFn   func(m *mock_usecase.MockSessionRepository, userId int)
		response        controllers.Response
	}{
		{
			name:            "sessionにuserIdが設定されている場合、ログアウト成功",
			userId:          1,
			withCredentials: true,
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
			name:            "リクエストにcookieの認証情報が含まれない場合(withCredentials: false)、ログアウト失敗",
			userId:          1,
			withCredentials: false,
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
		},
		{
			name:            "sessionにuserIdが設定されていない場合、ログアウト失敗",
			userId:          0,
			withCredentials: true,
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
			token, _ := controllers.MakeRandomStr(10)
			session, err := controllers.Store.Get(req, "session")
			if err != nil {
				t.Error(err)
				return
			}
			session.Values["userId"] = tt.userId
			session.Values["token"] = token
			err = session.Save(req, w)
			if err != nil {
				t.Error(err)
				return
			}
			if tt.withCredentials {
				req.AddCookie(&http.Cookie{
					Name: "cookie-name",
				})
			}
			ctrl.Logout(w, req)
			buf, _ := ioutil.ReadAll(w.Body)
			json.Unmarshal(buf, &response)
			assert.Equal(t, tt.response.Status, response.Status)
			assert.Equal(t, tt.response.Message, response.Message)
		})
	}
}
