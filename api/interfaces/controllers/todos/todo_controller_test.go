package controllers_test

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/sessions"
	"github.com/kory-jp/react_golang_mux/api/domain"
	controllers "github.com/kory-jp/react_golang_mux/api/interfaces/controllers/todos"
	"github.com/kory-jp/react_golang_mux/api/interfaces/database"
	mock_database "github.com/kory-jp/react_golang_mux/api/interfaces/mock"
	"github.com/stretchr/testify/assert"
)

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
var todo domain.Todo
var allTodosCount float64

type TodoMessage struct {
	Message string
	Error   string
}

type Response struct {
	Todos   domain.Todos
	SumPage int
	Message string
}

func TestCreate(t *testing.T) {
	// --- 各種mockをインスタンス ---
	sqlhandler, ctrl, result, _ := setMock(t)
	// --- api/interfaces/databases/todo_repository
	createTodoQuery := database.CreateTodoState

	cases := []struct {
		name            string
		args            domain.Todo
		isImage         bool
		prepareMockFn   func(m *mock_database.MockSqlHandler, r *mock_database.MockResult, statement string, todo domain.Todo)
		responseMessage string
	}{
		{
			name: "必須項目が入力された場合(画像有り)、データ保存成功",
			args: domain.Todo{
				UserID:    1,
				Title:     "test title",
				Content:   "test content",
				ImagePath: "test.png",
			},
			isImage: true,
			prepareMockFn: func(m *mock_database.MockSqlHandler, r *mock_database.MockResult, statement string, args domain.Todo) {
				r.EXPECT().RowsAffected().Return(int64(0), nil).AnyTimes()
				m.EXPECT().
					// --- 保存される画像名は自動生成される文字列となるためテストデータにおいて同名の画像名を引数として渡すことができないため
					// gomock.Any()を利用してどのような引数でも受け取れる手法を用いている(created_atも同様) ---
					Execute(statement, args.UserID, args.Title, args.Content, gomock.Any(), false, gomock.Any()).
					Return(r, nil)
			},
			responseMessage: "保存しました",
		},
		{
			name: "必須項目が入力された場合(画像無し),データ保存成功",
			args: domain.Todo{
				UserID:  1,
				Title:   "test title",
				Content: "test content",
			},
			isImage: false,
			prepareMockFn: func(m *mock_database.MockSqlHandler, r *mock_database.MockResult, statement string, args domain.Todo) {
				r.EXPECT().RowsAffected().Return(int64(0), nil).AnyTimes()
				m.EXPECT().
					Execute(statement, args.UserID, args.Title, args.Content, gomock.Any(), false, gomock.Any()).
					Return(r, nil)
			},
			responseMessage: "保存しました",
		},
		{
			name: "UserIDが0の場合,データ保存失敗",
			args: domain.Todo{
				UserID:  0,
				Title:   "test title",
				Content: "test content",
			},
			isImage: false,
			prepareMockFn: func(m *mock_database.MockSqlHandler, r *mock_database.MockResult, statement string, args domain.Todo) {
				r.EXPECT().RowsAffected().Return(int64(0), nil).AnyTimes()
				m.EXPECT().
					Execute(statement, args.UserID, args.Title, args.Content, gomock.Any(), false, gomock.Any()).
					Return(r, nil).
					AnyTimes()
			},
			responseMessage: "ログインしてください",
		},
		{
			name: "タイトルが0文字の場合,データ保存失敗",
			args: domain.Todo{
				UserID:  1,
				Title:   "",
				Content: "test content",
			},
			isImage: false,
			prepareMockFn: func(m *mock_database.MockSqlHandler, r *mock_database.MockResult, statement string, args domain.Todo) {
				r.EXPECT().RowsAffected().Return(int64(0), nil).AnyTimes()
				m.EXPECT().
					Execute(statement, args.UserID, args.Title, args.Content, gomock.Any(), false, gomock.Any()).
					Return(r, nil).
					AnyTimes()
			},
			responseMessage: "タイトルは必須です。",
		},
		{
			name: "タイトルが50文字以上の場合,データ保存失敗",
			args: domain.Todo{
				UserID:  1,
				Title:   strings.Repeat("t", 51),
				Content: "test content",
			},
			isImage: false,
			prepareMockFn: func(m *mock_database.MockSqlHandler, r *mock_database.MockResult, statement string, args domain.Todo) {
				r.EXPECT().RowsAffected().Return(int64(0), nil).AnyTimes()
				m.EXPECT().
					Execute(statement, args.UserID, args.Title, args.Content, gomock.Any(), false, gomock.Any()).
					Return(r, nil).
					AnyTimes()
			},
			responseMessage: "タイトルは50文字未満の入力になります。",
		},
		{
			name: "メモが2001文字以上の場合,データ保存失敗",
			args: domain.Todo{
				UserID:  1,
				Title:   "test title",
				Content: strings.Repeat("t", 2001),
			},
			isImage: false,
			prepareMockFn: func(m *mock_database.MockSqlHandler, r *mock_database.MockResult, statement string, args domain.Todo) {
				r.EXPECT().RowsAffected().Return(int64(0), nil).AnyTimes()
				m.EXPECT().
					Execute(statement, args.UserID, args.Title, args.Content, gomock.Any(), false, gomock.Any()).
					Return(r, nil).
					AnyTimes()
			},
			responseMessage: "メモは2000文字を超えて入力はできません。",
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			var buffer bytes.Buffer
			writer := multipart.NewWriter(&buffer)
			// -- 各種フィールドに値を設定 ---
			setField(t, writer, tt.isImage, tt.args.ImagePath, tt.args.Title, tt.args.Content)
			// --- フィールドの書き込みが終了後にClose ---
			writer.Close()

			req := httptest.NewRequest("POST", "/api/new", &buffer)
			req.Header.Add("Content-Type", writer.FormDataContentType())
			w := httptest.NewRecorder()

			// --- sessionにユーザーIDを保存処理 ---
			SetSessionUserId(t, w, req, tt.args.UserID)

			// --- mock ---
			tt.prepareMockFn(sqlhandler, result, createTodoQuery, tt.args)

			// --- テスト実行 ---
			ctrl.Create(w, req)
			var tm TodoMessage
			buf, _ := ioutil.ReadAll(w.Body)
			if err := json.Unmarshal(buf, &tm); err != nil {
				t.Error(err)
			}

			if w.Code != 200 {
				t.Error(w.Code)
			}

			if tm.Message != "" {
				assert.Equal(t, tm.Message, tt.responseMessage)
			}
			if tm.Error != "" {
				assert.Equal(t, tm.Error, tt.responseMessage)
			}
		})
	}
}

func TestIndex(t *testing.T) {
	sqlhandler, ctrl, _, row := setMock(t)
	sumTodoItemsQuery := database.SumTodoItemsState
	getTodosQuery := database.GetTodosState

	cases := []struct {
		name                     string
		loginUserId              int
		nowPage                  int
		prepareGetNumTodosMockFn func(m *mock_database.MockSqlHandler, r *mock_database.MockRow, statement string, userId int)
		prepareGetTodosMockFn    func(m *mock_database.MockSqlHandler, r *mock_database.MockRow, statement string, userId int, offset int, todo domain.Todo)
		message                  string
	}{
		{
			name:        "必須項目が入力された場合、データ取得に成功",
			loginUserId: 1,
			nowPage:     1,
			prepareGetNumTodosMockFn: func(m *mock_database.MockSqlHandler, r *mock_database.MockRow, statement string, userId int) {
				r.EXPECT().Next().Return(false).AnyTimes()
				r.EXPECT().Scan(&allTodosCount).Return(nil).AnyTimes()
				r.EXPECT().Err().Return(nil).AnyTimes()
				r.EXPECT().Close().Return(nil).AnyTimes()
				m.EXPECT().Query(statement, userId).Return(r, nil).AnyTimes()
			},
			prepareGetTodosMockFn: func(m *mock_database.MockSqlHandler, r *mock_database.MockRow, statement string, userId int, offset int, todo domain.Todo) {
				r.EXPECT().Scan(&todo.ID, &todo.UserID, &todo.Title, &todo.Content, &todo.ImagePath, &todo.IsFinished, &todo.CreatedAt).Return(nil).AnyTimes()
				r.EXPECT().Next().Return(false).AnyTimes()
				r.EXPECT().Err().Return(nil).AnyTimes()
				r.EXPECT().Close().Return(nil).AnyTimes()
				m.EXPECT().Query(statement, userId, offset).Return(r, nil).AnyTimes()
			},
			message: "",
		},
		{
			name:        "userIdが0の場合、データ取得に失敗",
			loginUserId: 0,
			nowPage:     1,
			prepareGetNumTodosMockFn: func(m *mock_database.MockSqlHandler, r *mock_database.MockRow, statement string, userId int) {
				r.EXPECT().Next().Return(false).AnyTimes()
				r.EXPECT().Scan(&allTodosCount).Return(nil).AnyTimes()
				r.EXPECT().Err().Return(nil).AnyTimes()
				r.EXPECT().Err().Return(nil).AnyTimes()
				r.EXPECT().Close().Return(nil).AnyTimes()
				m.EXPECT().Query(statement, userId).Return(r, nil).AnyTimes()
			},
			prepareGetTodosMockFn: func(m *mock_database.MockSqlHandler, r *mock_database.MockRow, statement string, userId int, offset int, todo domain.Todo) {
				r.EXPECT().Scan(&todo.ID, &todo.UserID, &todo.Title, &todo.Content, &todo.ImagePath, &todo.IsFinished, &todo.CreatedAt).Return(nil).AnyTimes()
				r.EXPECT().Next().Return(false).AnyTimes()
				r.EXPECT().Err().Return(nil).AnyTimes()
				r.EXPECT().Close().Return(nil).AnyTimes()
				m.EXPECT().Query(statement, userId, offset).Return(r, nil).AnyTimes()
			},
			message: "ログインしてください",
		},
		{
			name:        "現在ページ情報(nowPage)が0の場合、データ取得に失敗",
			loginUserId: 1,
			nowPage:     0,
			prepareGetNumTodosMockFn: func(m *mock_database.MockSqlHandler, r *mock_database.MockRow, statement string, userId int) {
				r.EXPECT().Next().Return(false).AnyTimes()
				r.EXPECT().Scan(&allTodosCount).Return(nil).AnyTimes()
				r.EXPECT().Err().Return(nil).AnyTimes()
				r.EXPECT().Close().Return(nil).AnyTimes()
				m.EXPECT().Query(statement, userId).Return(r, nil).AnyTimes()
			},
			prepareGetTodosMockFn: func(m *mock_database.MockSqlHandler, r *mock_database.MockRow, statement string, userId int, offset int, todo domain.Todo) {
				r.EXPECT().Scan(&todo.ID, &todo.UserID, &todo.Title, &todo.Content, &todo.ImagePath, &todo.IsFinished, &todo.CreatedAt).Return(nil).AnyTimes()
				r.EXPECT().Next().Return(false).AnyTimes()
				r.EXPECT().Err().Return(nil).AnyTimes()
				r.EXPECT().Close().Return(nil).AnyTimes()
				m.EXPECT().Query(statement, userId, offset).Return(r, nil).AnyTimes()
			},
			message: "データ取得に失敗しました",
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			var buffer bytes.Buffer
			writer := multipart.NewWriter(&buffer)
			writer.Close()
			apiURL := "/api/todos?page=" + strconv.Itoa(tt.nowPage)
			req := httptest.NewRequest("GET", apiURL, &buffer)
			req.Header.Add("Content-Type", writer.FormDataContentType())
			w := httptest.NewRecorder()
			SetSessionUserId(t, w, req, tt.loginUserId)

			tt.prepareGetNumTodosMockFn(sqlhandler, row, sumTodoItemsQuery, tt.loginUserId)
			tt.prepareGetTodosMockFn(sqlhandler, row, getTodosQuery, tt.loginUserId, 0, todo)

			ctrl.Index(w, req)
			var rsp Response
			var mess TodoMessage
			buf, _ := ioutil.ReadAll(w.Body)
			if err := json.Unmarshal(buf, &rsp); err != nil {
				t.Error(err)
			}
			if err := json.Unmarshal(buf, &mess); err != nil {
				t.Error(err)
			}

			if rsp.Message != tt.message {
				assert.Equal(t, mess.Error, tt.message)
			}
		})
	}
}

func TestShow(t *testing.T) {
	sqlhandler, ctrl, _, row := setMock(t)
	showTodoQuery := database.ShowTodoState

	cases := []struct {
		name          string
		todoId        int
		loginUserId   int
		prepareMockFn func(m *mock_database.MockSqlHandler, r *mock_database.MockRow, statement string, id int, userId int, todo domain.Todo)
		message       string
	}{
		{
			name:        "必須項目が入力された場合、データ取得に成功",
			todoId:      1,
			loginUserId: 1,
			prepareMockFn: func(m *mock_database.MockSqlHandler, r *mock_database.MockRow, statement string, id int, userId int, todo domain.Todo) {
				r.EXPECT().Scan(&todo.ID, &todo.UserID, &todo.Title, &todo.Content, &todo.ImagePath, &todo.IsFinished, &todo.CreatedAt).Return(nil).AnyTimes()
				r.EXPECT().Next().Return(false).AnyTimes()
				r.EXPECT().Err().Return(nil).AnyTimes()
				r.EXPECT().Close().Return(nil).AnyTimes()
				m.EXPECT().Query(statement, id, userId).Return(r, nil).AnyTimes()
			},
			message: "",
		},
		{
			name:        "todoIdがnilの場合、データ取得に失敗",
			todoId:      0,
			loginUserId: 1,
			prepareMockFn: func(m *mock_database.MockSqlHandler, r *mock_database.MockRow, statement string, id int, userId int, todo domain.Todo) {
				r.EXPECT().Scan(&todo.ID, &todo.UserID, &todo.Title, &todo.Content, &todo.ImagePath, &todo.IsFinished, &todo.CreatedAt).Return(nil).AnyTimes()
				r.EXPECT().Next().Return(false).AnyTimes()
				r.EXPECT().Err().Return(nil).AnyTimes()
				r.EXPECT().Close().Return(nil).AnyTimes()
				m.EXPECT().Query(statement, id, userId).Return(r, nil).AnyTimes()
			},
			message: "データ取得に失敗しました",
		},
		{
			name:        "userIdがnilの場合、データ取得に失敗",
			todoId:      1,
			loginUserId: 0,
			prepareMockFn: func(m *mock_database.MockSqlHandler, r *mock_database.MockRow, statement string, id int, userId int, todo domain.Todo) {
				r.EXPECT().Scan(&todo.ID, &todo.UserID, &todo.Title, &todo.Content, &todo.ImagePath, &todo.IsFinished, &todo.CreatedAt).Return(nil).AnyTimes()
				r.EXPECT().Next().Return(false).AnyTimes()
				r.EXPECT().Err().Return(nil).AnyTimes()
				r.EXPECT().Close().Return(nil).AnyTimes()
				m.EXPECT().Query(statement, id, userId).Return(r, nil).AnyTimes()
			},
			message: "ログインをしてください",
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			var buffer bytes.Buffer
			writer := multipart.NewWriter(&buffer)
			writer.Close()
			apiURL := "/api/todos/" + strconv.Itoa(tt.todoId)
			req := httptest.NewRequest("GET", apiURL, &buffer)
			req.Header.Add("Content-Type", writer.FormDataContentType())
			w := httptest.NewRecorder()
			SetSessionUserId(t, w, req, tt.loginUserId)

			tt.prepareMockFn(sqlhandler, row, showTodoQuery, tt.todoId, tt.loginUserId, todo)

			ctrl.Show(w, req)
			var mess *TodoMessage
			buf, _ := ioutil.ReadAll(w.Body)
			json.Unmarshal(buf, &mess)

			if mess.Error != "" {
				assert.Equal(t, mess.Error, tt.message)
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	sqlhandler, ctrl, result, _ := setMock(t)
	upateTodoQuery := database.UpdateTodoState

	cases := []struct {
		name            string
		args            domain.Todo
		loginUserId     int
		isImage         bool
		prepareMockFn   func(m *mock_database.MockSqlHandler, r *mock_database.MockResult, statement string, todo domain.Todo)
		responseMessage string
	}{
		{
			name: "必須項目が入力された場合(画像有り)、データ保存成功",
			args: domain.Todo{
				ID:        1,
				UserID:    1,
				Title:     "test title",
				Content:   "test content",
				ImagePath: "test.png",
			},
			loginUserId: 1,
			isImage:     true,
			prepareMockFn: func(m *mock_database.MockSqlHandler, r *mock_database.MockResult, statement string, args domain.Todo) {
				r.EXPECT().RowsAffected().Return(int64(0), nil).AnyTimes()
				m.EXPECT().
					Execute(statement, args.Title, args.Content, gomock.Any(), args.ID, args.UserID).
					Return(r, nil)
			},
			responseMessage: "更新しました",
		},
		{
			name: "必須項目が入力された場合(画像無し),データ保存成功",
			args: domain.Todo{
				ID:      1,
				UserID:  1,
				Title:   "test title",
				Content: "test content",
			},
			loginUserId: 1,
			isImage:     false,
			prepareMockFn: func(m *mock_database.MockSqlHandler, r *mock_database.MockResult, statement string, args domain.Todo) {
				r.EXPECT().RowsAffected().Return(int64(0), nil).AnyTimes()
				m.EXPECT().
					Execute(statement, args.Title, args.Content, gomock.Any(), args.ID, args.UserID).
					Return(r, nil)
			},
			responseMessage: "更新しました",
		},
		{
			name: "ログイン状態が確認できない場合,データ保存失敗",
			args: domain.Todo{
				ID:      1,
				UserID:  1,
				Title:   "test title",
				Content: "test content",
			},
			loginUserId: 0,
			isImage:     false,
			prepareMockFn: func(m *mock_database.MockSqlHandler, r *mock_database.MockResult, statement string, args domain.Todo) {
				r.EXPECT().RowsAffected().Return(int64(0), nil).AnyTimes()
				m.EXPECT().
					Execute(statement, args.Title, args.Content, gomock.Any(), args.ID, args.UserID).
					Return(r, nil).
					AnyTimes()
			},
			responseMessage: "ログインしてください",
		},
		{
			name: "タイトルが0文字の場合,データ保存失敗",
			args: domain.Todo{
				ID:      1,
				UserID:  1,
				Title:   "",
				Content: "test content",
			},
			loginUserId: 1,
			isImage:     false,
			prepareMockFn: func(m *mock_database.MockSqlHandler, r *mock_database.MockResult, statement string, args domain.Todo) {
				r.EXPECT().RowsAffected().Return(int64(0), nil).AnyTimes()
				m.EXPECT().
					Execute(statement, args.Title, args.Content, gomock.Any(), args.ID, args.UserID).
					Return(r, nil).
					AnyTimes()
			},
			responseMessage: "タイトルは必須です。",
		},
		{
			name: "タイトルが50文字以上の場合,データ保存失敗",
			args: domain.Todo{
				ID:      1,
				UserID:  1,
				Title:   strings.Repeat("t", 51),
				Content: "test content",
			},
			loginUserId: 1,
			isImage:     false,
			prepareMockFn: func(m *mock_database.MockSqlHandler, r *mock_database.MockResult, statement string, args domain.Todo) {
				r.EXPECT().RowsAffected().Return(int64(0), nil).AnyTimes()
				m.EXPECT().
					Execute(statement, args.Title, args.Content, gomock.Any(), args.ID, args.UserID).
					Return(r, nil).
					AnyTimes()
			},
			responseMessage: "タイトルは50文字未満の入力になります。",
		},
		{
			name: "メモが2001文字以上の場合,データ保存失敗",
			args: domain.Todo{
				ID:      1,
				UserID:  1,
				Title:   "test title",
				Content: strings.Repeat("t", 2001),
			},
			loginUserId: 1,
			isImage:     false,
			prepareMockFn: func(m *mock_database.MockSqlHandler, r *mock_database.MockResult, statement string, args domain.Todo) {
				r.EXPECT().RowsAffected().Return(int64(0), nil).AnyTimes()
				m.EXPECT().
					Execute(statement, args.Title, args.Content, gomock.Any(), args.ID, args.UserID).
					Return(r, nil).
					AnyTimes()
			},
			responseMessage: "メモは2000文字を超えて入力はできません。",
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			var buffer bytes.Buffer
			writer := multipart.NewWriter(&buffer)
			// -- 各種フィールドに値を設定 ---
			setField(t, writer, tt.isImage, tt.args.ImagePath, tt.args.Title, tt.args.Content)
			// --- フィールドの書き込みが終了後にClose ---
			writer.Close()
			// ---
			apiURL := "/api/todos/update/" + strconv.Itoa(tt.args.ID)
			req := httptest.NewRequest("POST", apiURL, &buffer)
			req.Header.Add("Content-Type", writer.FormDataContentType())
			w := httptest.NewRecorder()

			// --- sessionにユーザーIDを保存処理 ---
			SetSessionUserId(t, w, req, tt.loginUserId)

			// --- mock ---
			tt.prepareMockFn(sqlhandler, result, upateTodoQuery, tt.args)

			// --- テスト実行 ---
			ctrl.Update(w, req)
			var tm TodoMessage
			buf, _ := ioutil.ReadAll(w.Body)
			if err := json.Unmarshal(buf, &tm); err != nil {
				t.Error(err)
			}

			if w.Code != 200 {
				t.Error(w.Code)
			}

			if tm.Message != "" {
				assert.Equal(t, tm.Message, tt.responseMessage)
			}
			if tm.Error != "" {
				assert.Equal(t, tm.Error, tt.responseMessage)
			}
		})
	}
}

func TestIsFinished(t *testing.T) {
	sqlhandler, ctrl, result, row := setMock(t)
	changeBoolQuery := database.ChangeBoolState
	showTodoQuery := database.ShowTodoState

	cases := []struct {
		name                               string
		todoId                             int
		args                               domain.Todo
		loginUserId                        int
		prepareChangeBoolMockFn            func(m *mock_database.MockSqlHandler, r *mock_database.MockResult, statement string, todoId int, todo domain.Todo, userId int)
		prepareFindTodoByIdAndUserIdMockFn func(m *mock_database.MockSqlHandler, r *mock_database.MockRow, statement string, todoId int, userId int, todo domain.Todo)
		responseMessage                    string
	}{
		{
			name:   "必須項目が入力された場合、更新メッセージを取得",
			todoId: 1,
			args: domain.Todo{
				IsFinished: true,
			},
			loginUserId: 1,
			prepareChangeBoolMockFn: func(m *mock_database.MockSqlHandler, r *mock_database.MockResult, statement string, todoId int, todo domain.Todo, userId int) {
				r.EXPECT().RowsAffected().Return(int64(0), nil).AnyTimes()
				m.EXPECT().
					Execute(statement, todo.IsFinished, todoId, userId).
					Return(r, nil)
			},
			prepareFindTodoByIdAndUserIdMockFn: func(m *mock_database.MockSqlHandler, r *mock_database.MockRow, statement string, todoId int, userId int, todo domain.Todo) {
				r.EXPECT().Next().Return(false).AnyTimes()
				r.EXPECT().Scan(&todo.ID, &todo.UserID, &todo.Title, &todo.Content, &todo.ImagePath, &todo.IsFinished, &todo.CreatedAt).Return(nil).AnyTimes()
				r.EXPECT().Close().Return(nil).AnyTimes()
				m.EXPECT().Query(statement, todoId, userId).Return(r, nil)
			},
			responseMessage: "未完了の項目が追加されました",
		},
		{
			name:   "todoIdが0の場合、エラーメッセージを取得",
			todoId: 0,
			args: domain.Todo{
				IsFinished: true,
			},
			loginUserId: 1,
			prepareChangeBoolMockFn: func(m *mock_database.MockSqlHandler, r *mock_database.MockResult, statement string, todoId int, todo domain.Todo, userId int) {
				r.EXPECT().RowsAffected().Return(int64(0), nil).AnyTimes()
				m.EXPECT().
					Execute(statement, todo.IsFinished, todoId, userId).
					Return(r, nil).
					AnyTimes()
			},
			prepareFindTodoByIdAndUserIdMockFn: func(m *mock_database.MockSqlHandler, r *mock_database.MockRow, statement string, todoId int, userId int, todo domain.Todo) {
				r.EXPECT().Next().Return(false).AnyTimes()
				r.EXPECT().Scan(&todo.ID, &todo.UserID, &todo.Title, &todo.Content, &todo.ImagePath, &todo.IsFinished, &todo.CreatedAt).Return(nil).AnyTimes()
				r.EXPECT().Close().Return(nil).AnyTimes()
				m.EXPECT().Query(statement, todoId, userId).
					Return(r, nil).
					AnyTimes()
			},
			responseMessage: "データ取得に失敗しました",
		},
		{
			name:   "todoIdが0の場合、エラーメッセージを取得",
			todoId: 1,
			args: domain.Todo{
				IsFinished: true,
			},
			loginUserId: 0,
			prepareChangeBoolMockFn: func(m *mock_database.MockSqlHandler, r *mock_database.MockResult, statement string, todoId int, todo domain.Todo, userId int) {
				r.EXPECT().RowsAffected().Return(int64(0), nil).AnyTimes()
				m.EXPECT().
					Execute(statement, todo.IsFinished, todoId, userId).
					Return(r, nil).
					AnyTimes()
			},
			prepareFindTodoByIdAndUserIdMockFn: func(m *mock_database.MockSqlHandler, r *mock_database.MockRow, statement string, todoId int, userId int, todo domain.Todo) {
				r.EXPECT().Next().Return(false).AnyTimes()
				r.EXPECT().Scan(&todo.ID, &todo.UserID, &todo.Title, &todo.Content, &todo.ImagePath, &todo.IsFinished, &todo.CreatedAt).Return(nil).AnyTimes()
				r.EXPECT().Close().Return(nil).AnyTimes()
				m.EXPECT().Query(statement, todoId, userId).
					Return(r, nil).
					AnyTimes()
			},
			responseMessage: "ログインしてください",
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			jsonArgs, _ := json.Marshal(tt.args)
			apiURL := "/api/todos/isfinished/" + strconv.Itoa(tt.todoId)
			req := httptest.NewRequest("POST", apiURL, bytes.NewBuffer(jsonArgs))
			w := httptest.NewRecorder()

			// --- sessionにユーザーIDを保存処理 ---
			SetSessionUserId(t, w, req, tt.loginUserId)

			// --- mock ---
			tt.prepareChangeBoolMockFn(sqlhandler, result, changeBoolQuery, tt.todoId, tt.args, tt.loginUserId)
			tt.prepareFindTodoByIdAndUserIdMockFn(sqlhandler, row, showTodoQuery, tt.todoId, tt.loginUserId, todo)

			// --- テスト実行 ---
			ctrl.IsFinished(w, req)
			var tm TodoMessage
			buf, _ := ioutil.ReadAll(w.Body)
			if err := json.Unmarshal(buf, &tm); err != nil {
				t.Error(err)
			}

			if w.Code != 200 {
				t.Error(w.Code)
			}

			if tm.Message != "" {
				assert.Equal(t, tm.Message, tt.responseMessage)
			}
			if tm.Error != "" {
				assert.Equal(t, tm.Error, tt.responseMessage)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	sqlhandler, ctrl, result, _ := setMock(t)
	deleteTodoQuery := database.DeleteTodoState

	cases := []struct {
		name            string
		todoId          int
		loginUserId     int
		prepareMockFn   func(m *mock_database.MockSqlHandler, r *mock_database.MockResult, statement string, todoId int, loginUserId int)
		responseMessage string
	}{
		{
			name:        "必須項目が入力された場合、データ削除成功",
			todoId:      1,
			loginUserId: 1,
			prepareMockFn: func(m *mock_database.MockSqlHandler, r *mock_database.MockResult, statement string, todoId int, loginUserId int) {
				r.EXPECT().RowsAffected().Return(int64(0), nil).AnyTimes()
				m.EXPECT().
					Execute(statement, todoId, loginUserId).
					Return(r, nil)
			},
			responseMessage: "削除しました",
		},
		{
			name:        "todoIdがnilの場合、データ削除失敗",
			todoId:      0,
			loginUserId: 1,
			prepareMockFn: func(m *mock_database.MockSqlHandler, r *mock_database.MockResult, statement string, todoId int, loginUserId int) {
				r.EXPECT().RowsAffected().Return(int64(0), nil).AnyTimes()
				m.EXPECT().
					Execute(statement, todoId, loginUserId).
					Return(r, nil).
					AnyTimes()
			},
			responseMessage: "データ取得に失敗しました",
		},
		{
			name:        "loginUserIdがnilの場合、データ削除失敗",
			todoId:      1,
			loginUserId: 0,
			prepareMockFn: func(m *mock_database.MockSqlHandler, r *mock_database.MockResult, statement string, todoId int, loginUserId int) {
				r.EXPECT().RowsAffected().Return(int64(0), nil).AnyTimes()
				m.EXPECT().
					Execute(statement, todoId, loginUserId).
					Return(r, nil).
					AnyTimes()
			},
			responseMessage: "ログインしてください",
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			apiURL := "/api/todos/delete/" + strconv.Itoa(tt.todoId)
			req := httptest.NewRequest("DELETE", apiURL, nil)
			w := httptest.NewRecorder()

			// --- sessionにユーザーIDを保存処理 ---
			SetSessionUserId(t, w, req, tt.loginUserId)

			// --- mock ---
			tt.prepareMockFn(sqlhandler, result, deleteTodoQuery, tt.todoId, tt.loginUserId)

			// --- テスト実行 ---
			ctrl.Delete(w, req)
			var tm TodoMessage
			buf, _ := ioutil.ReadAll(w.Body)
			if err := json.Unmarshal(buf, &tm); err != nil {
				t.Error(err)
			}

			if w.Code != 200 {
				t.Error(w.Code)
			}

			if tm.Message != "" {
				assert.Equal(t, tm.Message, tt.responseMessage)
			}
			if tm.Error != "" {
				assert.Equal(t, tm.Error, tt.responseMessage)
			}
		})
	}
}

func TestDeleteIndex(t *testing.T) {
	sqlhandler, ctrl, result, row := setMock(t)
	deleteTodoQuery := database.DeleteTodoState
	getSumTodoItemsQuery := database.SumTodoItemsState
	getTodosQuery := database.GetTodosState

	cases := []struct {
		name                     string
		todoId                   int
		loginUserId              int
		page                     int
		prepareDeleteMockFn      func(m *mock_database.MockSqlHandler, r *mock_database.MockResult, statement string, todoId int, loginUserId int)
		prepareGetNumTodosMockFn func(m *mock_database.MockSqlHandler, r *mock_database.MockRow, statement string, userId int)
		prepareGetTodosMockFn    func(m *mock_database.MockSqlHandler, r *mock_database.MockRow, statement string, userId int, offset int, todo domain.Todo)
		responseMessage          string
	}{
		{
			name:        "必須項目が入力された場合、データ削除成功",
			todoId:      2,
			loginUserId: 1,
			page:        1,
			prepareDeleteMockFn: func(m *mock_database.MockSqlHandler, r *mock_database.MockResult, statement string, todoId int, loginUserId int) {
				r.EXPECT().RowsAffected().Return(int64(0), nil).AnyTimes()
				m.EXPECT().
					Execute(statement, todoId, loginUserId).
					Return(r, nil)
			},
			prepareGetNumTodosMockFn: func(m *mock_database.MockSqlHandler, r *mock_database.MockRow, statement string, userId int) {
				r.EXPECT().Next().Return(false).AnyTimes()
				r.EXPECT().Scan(&allTodosCount).Return(nil).AnyTimes()
				r.EXPECT().Err().Return(nil).AnyTimes()
				r.EXPECT().Close().Return(nil).AnyTimes()
				m.EXPECT().Query(statement, userId).Return(r, nil).AnyTimes()
			},
			prepareGetTodosMockFn: func(m *mock_database.MockSqlHandler, r *mock_database.MockRow, statement string, userId int, offset int, todo domain.Todo) {
				r.EXPECT().Scan(&todo.ID, &todo.UserID, &todo.Title, &todo.Content, &todo.ImagePath, &todo.IsFinished, &todo.CreatedAt).Return(nil).AnyTimes()
				r.EXPECT().Next().Return(false).AnyTimes()
				r.EXPECT().Err().Return(nil).AnyTimes()
				r.EXPECT().Close().Return(nil).AnyTimes()
				m.EXPECT().Query(statement, userId, offset).Return(r, nil).AnyTimes()
			},
			responseMessage: "削除しました",
		},
		{
			name:        "todoIdがnilの場合、データ削除失敗",
			todoId:      0,
			loginUserId: 1,
			page:        1,
			prepareDeleteMockFn: func(m *mock_database.MockSqlHandler, r *mock_database.MockResult, statement string, todoId int, loginUserId int) {
				r.EXPECT().RowsAffected().Return(int64(0), nil).AnyTimes()
				m.EXPECT().
					Execute(statement, todoId, loginUserId).
					Return(r, nil).
					AnyTimes()
			},
			prepareGetNumTodosMockFn: func(m *mock_database.MockSqlHandler, r *mock_database.MockRow, statement string, userId int) {
				r.EXPECT().Next().Return(false).AnyTimes()
				r.EXPECT().Scan(&allTodosCount).Return(nil).AnyTimes()
				r.EXPECT().Err().Return(nil).AnyTimes()
				r.EXPECT().Close().Return(nil).AnyTimes()
				m.EXPECT().Query(statement, userId).Return(r, nil).AnyTimes()
			},
			prepareGetTodosMockFn: func(m *mock_database.MockSqlHandler, r *mock_database.MockRow, statement string, userId int, offset int, todo domain.Todo) {
				r.EXPECT().Scan(&todo.ID, &todo.UserID, &todo.Title, &todo.Content, &todo.ImagePath, &todo.IsFinished, &todo.CreatedAt).Return(nil).AnyTimes()
				r.EXPECT().Next().Return(false).AnyTimes()
				r.EXPECT().Err().Return(nil).AnyTimes()
				r.EXPECT().Close().Return(nil).AnyTimes()
				m.EXPECT().Query(statement, userId, offset).Return(r, nil).AnyTimes()
			},
			responseMessage: "データ取得に失敗しました",
		},
		{
			name:        "loginUserIdがnilの場合、データ削除失敗",
			todoId:      1,
			loginUserId: 0,
			page:        1,
			prepareDeleteMockFn: func(m *mock_database.MockSqlHandler, r *mock_database.MockResult, statement string, todoId int, loginUserId int) {
				r.EXPECT().RowsAffected().Return(int64(0), nil).AnyTimes()
				m.EXPECT().
					Execute(statement, todoId, loginUserId).
					Return(r, nil).
					AnyTimes()
			},
			prepareGetNumTodosMockFn: func(m *mock_database.MockSqlHandler, r *mock_database.MockRow, statement string, userId int) {
				r.EXPECT().Next().Return(false).AnyTimes()
				r.EXPECT().Scan(&allTodosCount).Return(nil).AnyTimes()
				r.EXPECT().Err().Return(nil).AnyTimes()
				r.EXPECT().Close().Return(nil).AnyTimes()
				m.EXPECT().Query(statement, userId).Return(r, nil).AnyTimes()
			},
			prepareGetTodosMockFn: func(m *mock_database.MockSqlHandler, r *mock_database.MockRow, statement string, userId int, offset int, todo domain.Todo) {
				r.EXPECT().Scan(&todo.ID, &todo.UserID, &todo.Title, &todo.Content, &todo.ImagePath, &todo.IsFinished, &todo.CreatedAt).Return(nil).AnyTimes()
				r.EXPECT().Next().Return(false).AnyTimes()
				r.EXPECT().Err().Return(nil).AnyTimes()
				r.EXPECT().Close().Return(nil).AnyTimes()
				m.EXPECT().Query(statement, userId, offset).Return(r, nil).AnyTimes()
			},
			responseMessage: "ログインしてください",
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			apiURL := "/api/todos/deleteinindex/" + strconv.Itoa(tt.todoId) + "?page=" + strconv.Itoa(tt.page)
			req := httptest.NewRequest("DELETE", apiURL, nil)
			w := httptest.NewRecorder()

			// --- sessionにユーザーIDを保存処理 ---
			SetSessionUserId(t, w, req, tt.loginUserId)

			// --- mock ---
			tt.prepareDeleteMockFn(sqlhandler, result, deleteTodoQuery, tt.todoId, tt.loginUserId)
			tt.prepareGetNumTodosMockFn(sqlhandler, row, getSumTodoItemsQuery, tt.loginUserId)
			tt.prepareGetTodosMockFn(sqlhandler, row, getTodosQuery, tt.loginUserId, 0, todo)

			// --- テスト実行 ---
			ctrl.DeleteInIndex(w, req)
			var tm TodoMessage
			buf, _ := ioutil.ReadAll(w.Body)
			if err := json.Unmarshal(buf, &tm); err != nil {
				t.Error(err)
			}

			if w.Code != 200 {
				t.Error(w.Code)
			}

			if tm.Message != "" {
				assert.Equal(t, tm.Message, tt.responseMessage)
			}
			if tm.Error != "" {
				assert.Equal(t, tm.Error, tt.responseMessage)
			}
		})
	}
}

func setMock(t *testing.T) (sqlhandler *mock_database.MockSqlHandler, ctrl *controllers.TodoController, result *mock_database.MockResult, row *mock_database.MockRow) {
	c := gomock.NewController(t)
	defer c.Finish()
	// --- api/interfaces/database/sqlhandlerのモック ---
	sqlhandler = mock_database.NewMockSqlHandler(c)
	ctrl = controllers.NewTodoController(sqlhandler)
	result = mock_database.NewMockResult(c)
	row = mock_database.NewMockRow(c)
	return
}

func setField(t *testing.T, writer *multipart.Writer, isImage bool, imagePath string, title string, content string) {
	// --- 画像データ ---
	if isImage {
		// fileWriter, err := writer.CreateFormFile("image", tt.args.ImagePath)
		fileWriter, err := writer.CreateFormFile("image", imagePath)
		if err != nil {
			t.Fatalf("Failed to create file writer. %s", err)
		}
		imgPath := "../../../assets/test/img/" + imagePath
		readFile, err := os.Open(imgPath)
		if err != nil {
			t.Fatalf("Failde to create file writer. %s", err)
		}
		defer readFile.Close()
		io.Copy(fileWriter, readFile)
	}

	// --- タイトルデータ ---
	titleWriter, err := writer.CreateFormField("title")
	if err != nil {
		t.Fatalf("Failed to create file writer. %s", err)
	}
	titleWriter.Write([]byte(title))
	// --- 本文データ ---
	contentWriter, err := writer.CreateFormField("content")
	if err != nil {
		t.Fatalf("Failed to create file writer. %s", err)
	}
	contentWriter.Write([]byte(content))
}

// --- ユーザーのログイン情報をsessionに設定 ---
func SetSessionUserId(t *testing.T, w *httptest.ResponseRecorder, req *http.Request, userId int) {
	session, err := store.Get(req, "session")
	if err != nil {
		t.Error(err)
		return
	}
	session.Values["userId"] = userId
	err = session.Save(req, w)
	if err != nil {
		t.Error(err)
	}
}
