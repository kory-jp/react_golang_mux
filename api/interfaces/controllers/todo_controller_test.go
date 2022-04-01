package controllers_test

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http/httptest"
	"os"
	"strconv"
	"strings"
	"testing"

	mock_database "github.com/kory-jp/react_golang_mux/api/interfaces/mock"
	"github.com/stretchr/testify/assert"

	"github.com/golang/mock/gomock"
	"github.com/kory-jp/react_golang_mux/api/domain"
	"github.com/kory-jp/react_golang_mux/api/interfaces/controllers"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

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
	c := gomock.NewController(t)
	defer c.Finish()
	// --- api/interfaces/database/sqlhandlerのモック ---
	sqlhandler := mock_database.NewMockSqlHandler(c)
	ctrl := controllers.NewTodoController(sqlhandler)
	result := mock_database.NewMockResult(c)
	statement := `
		insert into
			todos(
				user_id,
				title,
				content,
				image_path,
				isFinished,
				created_at
			)
		value (?, ?, ?, ?, ?, ?)
	`

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

			// --- 画像データ ---
			if tt.isImage {
				fileWriter, err := writer.CreateFormFile("image", tt.args.ImagePath)
				if err != nil {
					t.Fatalf("Failed to create file writer. %s", err)
				}

				imgPath := "../../assets/test/img/" + tt.args.ImagePath
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
			titleWriter.Write([]byte(tt.args.Title))

			// --- 本文データ ---
			contentWriter, err := writer.CreateFormField("content")
			if err != nil {
				t.Fatalf("Failed to create file writer. %s", err)
			}
			contentWriter.Write([]byte(tt.args.Content))

			// --- フィールドの書き込みが終了後にClose ---
			writer.Close()

			req := httptest.NewRequest("POST", "/api/new", &buffer)
			req.Header.Add("Content-Type", writer.FormDataContentType())
			w := httptest.NewRecorder()

			// --- sessionにユーザーIDを保存処理 ---
			session, err := store.Get(req, "session")
			if err != nil {
				t.Error(err)
				return
			}
			session.Values["userId"] = tt.args.UserID
			err = session.Save(req, w)
			if err != nil {
				t.Error(err)
			}

			// --- mock ---
			tt.prepareMockFn(sqlhandler, result, statement, tt.args)

			// --- テスト実行 ---
			ctrl.Create(w, req)
			var tm TodoMessage
			buf, _ := ioutil.ReadAll(w.Body)
			if err = json.Unmarshal(buf, &tm); err != nil {
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
	c := gomock.NewController(t)
	defer c.Finish()
	sqlhandler := mock_database.NewMockSqlHandler(c)
	ctrl := controllers.NewTodoController(sqlhandler)
	row := mock_database.NewMockRow(c)
	var allTodosCount float64

	// 現在の投稿済みTodo総数取得
	statement1 := `
		select count(*) from
			todos
		where
			user_id = ?
	`

	statement2 := `
		select
			*
		from
			todos
		where
			user_id = ?
		order by
			id desc
		limit 5
		offset ?
	`
	var todo domain.Todo

	cases := []struct {
		name                     string
		userId                   int
		nowPage                  int
		prepareGetNumTodosMockFn func(m *mock_database.MockSqlHandler, r *mock_database.MockRow, statement string, userId int)
		prepareGetTodosMockFn    func(m *mock_database.MockSqlHandler, r *mock_database.MockRow, statement string, userId int, offset int, todo domain.Todo)
		message                  string
	}{
		{
			name:    "必須項目が入力された場合、データ取得に成功",
			userId:  1,
			nowPage: 1,
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
			name:    "userIdが0の場合、データ取得に失敗",
			userId:  0,
			nowPage: 1,
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
			name:    "現在ページ情報(nowPage)が0の場合、データ取得に失敗",
			userId:  1,
			nowPage: 0,
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

			session, err := store.Get(req, "session")
			if err != nil {
				t.Error(err)
				return
			}
			session.Values["userId"] = tt.userId
			err = session.Save(req, w)
			if err != nil {
				t.Error(err)
			}

			tt.prepareGetNumTodosMockFn(sqlhandler, row, statement1, tt.userId)
			tt.prepareGetTodosMockFn(sqlhandler, row, statement2, tt.userId, 0, todo)

			ctrl.Index(w, req)
			var rsp Response
			var mess TodoMessage
			buf, _ := ioutil.ReadAll(w.Body)
			if err = json.Unmarshal(buf, &rsp); err != nil {
				t.Error(err)
			}
			if err = json.Unmarshal(buf, &mess); err != nil {
				t.Error(err)
			}

			if rsp.Message != tt.message {
				assert.Equal(t, mess.Error, tt.message)
			}
		})
	}
}

func TestShow(t *testing.T) {
	c := gomock.NewController(t)
	defer c.Finish()
	sqlhandler := mock_database.NewMockSqlHandler(c)
	ctrl := controllers.NewTodoController(sqlhandler)
	row := mock_database.NewMockRow(c)

	stateShow := `
		select
			id,
			user_id,
			title,
			content,
			image_path,
			isFinished,
			created_at
		from
			todos
		where
			id = ?
		and
			user_id = ?
	`

	var todo domain.Todo

	cases := []struct {
		name          string
		todoId        int
		userId        int
		prepareMockFn func(m *mock_database.MockSqlHandler, r *mock_database.MockRow, statement string, id int, userId int, todo domain.Todo)
		message       string
	}{
		{
			name:   "必須項目が入力された場合、データ取得に成功",
			todoId: 1,
			userId: 1,
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
			name:   "todoIdがnilの場合、データ取得に失敗",
			todoId: 0,
			userId: 1,
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
			name:   "userIdがnilの場合、データ取得に失敗",
			todoId: 1,
			userId: 0,
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

			session, err := store.Get(req, "session")
			if err != nil {
				t.Error(err)
				return
			}
			session.Values["userId"] = tt.userId
			err = session.Save(req, w)
			if err != nil {
				t.Error(err)
			}

			tt.prepareMockFn(sqlhandler, row, stateShow, tt.todoId, tt.userId, todo)

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
	c := gomock.NewController(t)
	defer c.Finish()
	sqlhandler := mock_database.NewMockSqlHandler(c)
	ctrl := controllers.NewTodoController(sqlhandler)
	result := mock_database.NewMockResult(c)
	statement := `
		update
			todos
		set
			title = ?,
			content = ?,
			image_path = ?
		where
			id = ?
		and
			user_id = ?
	`

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

			// --- 画像データ ---
			if tt.isImage {
				fileWriter, err := writer.CreateFormFile("image", tt.args.ImagePath)
				if err != nil {
					t.Fatalf("Failed to create file writer. %s", err)
				}

				imgPath := "../../assets/test/img/" + tt.args.ImagePath
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
			titleWriter.Write([]byte(tt.args.Title))

			// --- 本文データ ---
			contentWriter, err := writer.CreateFormField("content")
			if err != nil {
				t.Fatalf("Failed to create file writer. %s", err)
			}
			contentWriter.Write([]byte(tt.args.Content))

			// --- フィールドの書き込みが終了後にClose ---
			writer.Close()
			// ---
			apiURL := "/api/todos/update/" + strconv.Itoa(tt.args.ID)
			req := httptest.NewRequest("POST", apiURL, &buffer)
			req.Header.Add("Content-Type", writer.FormDataContentType())
			w := httptest.NewRecorder()

			// --- sessionにユーザーIDを保存処理 ---
			session, err := store.Get(req, "session")
			if err != nil {
				t.Error(err)
				return
			}
			session.Values["userId"] = tt.loginUserId
			err = session.Save(req, w)
			if err != nil {
				t.Error(err)
			}

			// --- mock ---
			tt.prepareMockFn(sqlhandler, result, statement, tt.args)

			// --- テスト実行 ---
			ctrl.Update(w, req)
			var tm TodoMessage
			buf, _ := ioutil.ReadAll(w.Body)
			if err = json.Unmarshal(buf, &tm); err != nil {
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
