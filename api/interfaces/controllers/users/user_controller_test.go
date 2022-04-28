package controllers_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	usecase "github.com/kory-jp/react_golang_mux/api/usecase/users"
	mock_usecase "github.com/kory-jp/react_golang_mux/api/usecase/users/mock"

	"github.com/stretchr/testify/assert"

	"github.com/kory-jp/react_golang_mux/api/domain"

	"github.com/golang/mock/gomock"
	controllers "github.com/kory-jp/react_golang_mux/api/interfaces/controllers/users"
	mock_database "github.com/kory-jp/react_golang_mux/api/interfaces/mock"
)

var response *controllers.Response

func TestCreate(t *testing.T) {
	// --- mockを新規インスタンス ---
	c := gomock.NewController(t)
	defer c.Finish()
	sqlhandler := mock_database.NewMockSqlHandler(c)
	ctrl := controllers.NewUserController(sqlhandler)
	UserRepository := mock_usecase.NewMockUserRepository(c)
	inter := &usecase.UserInteractor{}
	inter.UserRepository = UserRepository
	ctrl.Interactor = *inter

	// --- テストケース ---
	cases := []struct {
		name               string
		args               domain.User
		userId             int
		requestBody        bool
		prepareStoreMockFn func(m *mock_usecase.MockUserRepository, args domain.User)
		prepareFindMockFn  func(m *mock_usecase.MockUserRepository, id int)
		response           controllers.Response
	}{
		{
			name: "必須項目が入力された場合、データ保存成功",
			args: domain.User{
				Name:     "testUser",
				Email:    "test@exm.com",
				Password: "testPassword",
			},
			userId:      1,
			requestBody: true,
			prepareStoreMockFn: func(m *mock_usecase.MockUserRepository, args domain.User) {
				m.EXPECT().Store(gomock.Any()).Return(1, nil)
			},
			prepareFindMockFn: func(m *mock_usecase.MockUserRepository, id int) {
				m.EXPECT().FindById(id).Return(nil, nil)
			},
			response: controllers.Response{
				Status:  200,
				Message: "新規登録完了しました",
			},
		},
		{
			name: "リクエストボディがnilの場合、データ保存失敗",
			args: domain.User{
				Name:     "testUser",
				Email:    "test@exm.com",
				Password: "testPassword",
			},
			userId:      1,
			requestBody: false,
			prepareStoreMockFn: func(m *mock_usecase.MockUserRepository, args domain.User) {
				m.EXPECT().Store(gomock.Any()).Return(1, nil).AnyTimes()
			},
			prepareFindMockFn: func(m *mock_usecase.MockUserRepository, id int) {
				m.EXPECT().FindById(id).Return(nil, nil).AnyTimes()
			},
			response: controllers.Response{
				Status:  400,
				Message: "データ取得に失敗しました",
			},
		},
	}

	// --- テスト実行 ---
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			var req *http.Request
			jsonArgs, _ := json.Marshal(tt.args)
			apiURL := "/api/registration"
			if tt.requestBody {
				req = httptest.NewRequest("POST", apiURL, bytes.NewBuffer(jsonArgs))
			} else {
				req = httptest.NewRequest("POST", apiURL, nil)
			}
			w := httptest.NewRecorder()

			// --- mockインスタンス ---
			tt.prepareStoreMockFn(UserRepository, tt.args)
			tt.prepareFindMockFn(UserRepository, tt.userId)

			// --- 通信実行 ---
			ctrl.Create(w, req)
			buf, _ := ioutil.ReadAll(w.Body)
			json.Unmarshal(buf, &response)
			assert.Equal(t, tt.response.Status, response.Status)
			assert.Equal(t, tt.response.Message, response.Message)
		})
	}
}
