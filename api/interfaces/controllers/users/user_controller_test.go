package controllers_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/kory-jp/react_golang_mux/api/domain"

	"github.com/kory-jp/react_golang_mux/api/interfaces/database"

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
	result := mock_database.NewMockResult(c)
	row := mock_database.NewMockRow(c)
	var user domain.User
	// var pUser *domain.User
	// --- databaseMockの引数へ渡すSQLクエリ ---
	createUserQuery := database.CreateUserState
	findUserQuery := database.FindUserState

	// --- テストケース ---
	cases := []struct {
		name               string
		args               domain.User
		userId             int
		requestBody        bool
		prepareStoreMockFn func(m *mock_database.MockSqlHandler, r *mock_database.MockResult, statement string, user *domain.User)
		prepareFindMockFn  func(m *mock_database.MockSqlHandler, r *mock_database.MockRow, statement string, userId int, user domain.User)
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
			prepareStoreMockFn: func(m *mock_database.MockSqlHandler, r *mock_database.MockResult, statement string, user *domain.User) {
				r.EXPECT().LastInsertId().Return(int64(0), nil)
				m.EXPECT().Execute(statement, gomock.Any()).Return(r, nil)
			},
			prepareFindMockFn: func(m *mock_database.MockSqlHandler, r *mock_database.MockRow, statement string, userId int, user domain.User) {
				r.EXPECT().Next().Return(false)
				r.EXPECT().Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt).Return(nil)
				r.EXPECT().Close().Return(nil)
				m.EXPECT().Query(statement, user.ID).Return(r, nil)
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
			prepareStoreMockFn: func(m *mock_database.MockSqlHandler, r *mock_database.MockResult, statement string, user *domain.User) {
				r.EXPECT().LastInsertId().Return(int64(0), nil).AnyTimes()
				m.EXPECT().Execute(statement, gomock.Any()).Return(r, nil).AnyTimes()
			},
			prepareFindMockFn: func(m *mock_database.MockSqlHandler, r *mock_database.MockRow, statement string, userId int, user domain.User) {
				r.EXPECT().Next().Return(false).AnyTimes()
				r.EXPECT().Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt).Return(nil).AnyTimes()
				r.EXPECT().Close().Return(nil).AnyTimes()
				m.EXPECT().Query(statement, user.ID).Return(r, nil).AnyTimes()
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
			pointerUser := &tt.args

			// --- mockインスタンス ---
			tt.prepareStoreMockFn(sqlhandler, result, createUserQuery, pointerUser)
			tt.prepareFindMockFn(sqlhandler, row, findUserQuery, tt.userId, user)

			// --- 通信実行 ---
			ctrl.Create(w, req)
			buf, _ := ioutil.ReadAll(w.Body)
			json.Unmarshal(buf, &response)
			assert.Equal(t, tt.response.Status, response.Status)
			assert.Equal(t, tt.response.Message, response.Message)
		})
	}
}
