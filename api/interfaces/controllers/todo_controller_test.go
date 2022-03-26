package controllers_test

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/kory-jp/react_golang_mux/api/domain"
	mock_database "github.com/kory-jp/react_golang_mux/api/interfaces/mock"

	"github.com/golang/mock/gomock"
	"github.com/kory-jp/react_golang_mux/api/interfaces/controllers"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

type TodoMessage struct {
	Message string
	Error   string
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
			responseMessage: "ユーザーIDは必須です。",
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
				if tm.Message != tt.responseMessage {
					t.Error("actual:", tm.Message, "want:", tt.responseMessage)
					return
				}
			}
			if tm.Error != "" {
				if tm.Error != tt.responseMessage {
					t.Error("actual:", tm.Error, "want:", tt.responseMessage)
					return
				}
			}
		})
	}
}
