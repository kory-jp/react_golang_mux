package controllers_test

import (
	"fmt"
	"testing"

	controllers "github.com/kory-jp/react_golang_mux/api/interfaces/controllers/sessions"
	"github.com/stretchr/testify/assert"
)

type SessError struct {
	Error string
}

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

// func TestLogin(t *testing.T) {
// 	c := gomock.NewController(t)
// 	defer c.Finish()
// 	sqlhandler := mock_database.NewMockSqlHandler(c)
// 	ctrl := controllers.NewSessionController(sqlhandler)
// 	row := mock_database.NewMockRow(c)
// 	loginQuery := database.FindByEmailState
// 	var user domain.User

// 	cases := []struct {
// 		name          string
// 		args          domain.User
// 		prepareMockFn func(m *mock_database.MockSqlHandler, r *mock_database.MockRow, statement string, args domain.User)
// 		wantErr       string
// 	}{
// 		{
// 			name: "必須項目が入力された場合、ログイン成功",
// 			args: domain.User{
// 				Email:    "test@exm.com",
// 				Password: "testPassword",
// 			},
// 			prepareMockFn: func(m *mock_database.MockSqlHandler, r *mock_database.MockRow, statement string, args domain.User) {
// 				r.EXPECT().Next().Return(false)
// 				r.EXPECT().Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt).Return(nil)
// 				r.EXPECT().Close().Return(nil)
// 				m.EXPECT().Query(statement, args.Email).Return(r, nil)
// 			},
// 			wantErr: "",
// 		},
// 	}

// 	for _, tt := range cases {
// 		t.Run(tt.name, func(t *testing.T) {
// 			jsonArgs, _ := json.Marshal(tt.args)
// 			apiURL := "/api/login"
// 			req := httptest.NewRequest("POST", apiURL, bytes.NewBuffer(jsonArgs))
// 			w := httptest.NewRecorder()

// 			tt.prepareMockFn(sqlhandler, row, loginQuery, tt.args)

// 			ctrl.Login(w, req)
// 			var sErr *SessError
// 			buf, _ := ioutil.ReadAll(w.Body)
// 			json.Unmarshal(buf, &sErr)
// 			assert.Equal(t, tt.wantErr, sErr.Error)
// 		})
// 	}
// }
