package controllers_test

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/kory-jp/react_golang_mux/api/domain"
	mock_database "github.com/kory-jp/react_golang_mux/api/interfaces/mock"

	"github.com/golang/mock/gomock"
	"github.com/kory-jp/react_golang_mux/api/interfaces/controllers"

	"github.com/gorilla/sessions"
)

// var mux *http.ServeMux

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

func TestCreate(t *testing.T) {

	c := gomock.NewController(t)
	defer c.Finish()
	// --- api/interfaces/database/sqlhandlerのモック ---
	sqlhandler := mock_database.NewMockSqlHandler(c)
	ctrl := controllers.NewTodoController(sqlhandler)
	// mux = http.NewServeMux()
	// mux.HandleFunc("/api/new", ctrl.Create)

	// ---画像データ---
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)
	fileWriter, err := writer.CreateFormFile("image", "test.png")
	if err != nil {
		t.Fatalf("Failed to create file writer. %s", err)
	}

	readFile, err := os.Open("test.png")
	if err != nil {
		t.Fatalf("Failde to create file writer. %s", err)
	}
	defer readFile.Close()
	io.Copy(fileWriter, readFile)

	// --- タイトルデータ ---
	titleWriter, err := writer.CreateFormField("title")
	if err != nil {
		t.Fatalf("Failed to create file writer. %s", err)
	}
	titleWriter.Write([]byte("test title"))

	// --- フィールドの書き込みが終了後にClose ---
	writer.Close()

	req := httptest.NewRequest("POST", "/api/new", &buffer)
	req.Header.Add("Content-Type", writer.FormDataContentType())
	w := httptest.NewRecorder()

	// session, err := store.New(req, "session")
	// --- sessionにユーザーIDを保存処理 ---
	session, err := store.Get(req, "session")
	if err != nil {
		fmt.Println(err)
		return
	}
	session.Values["userId"] = 1
	err = session.Save(req, w)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("71:", session.Values["userId"])

	// --- prepareMock ---
	// --- interfaces/database/todo_repository.go Storeにて実行されるExecuteのMock + 戻り値のResultのMock

	result := mock_database.NewMockResult(c)
	statement := "insert into todos(user_id, title, content, image_path, isFinished, created_at) value (?, ?, ?, ?, ?, ?)"
	todo := domain.Todo{
		UserID:     1,
		Title:      "test title",
		Content:    "test content",
		ImagePath:  "testimage",
		IsFinished: false,
		CreatedAt:  time.Now(),
	}

	exeMock := func(m *mock_database.MockSqlHandler, r *mock_database.MockResult, statement string, todo domain.Todo) {
		m.EXPECT().
			Execute(statement, todo.UserID, todo.Title, todo.Content, todo.ImagePath, false, time.Now()).
			Return(r.EXPECT().RowsAffected().Return(int64(0), nil), nil)
	}

	exeMock(sqlhandler, result, statement, todo)

	// --- prepareMock ---

	// --- テスト実行 ---
	t.Run("test135", func(t *testing.T) {
		// ctrl.Create(w, req)
		// mux.ServeHTTP(w, req)
		ctrl.Create(w, req)
		fmt.Println(w.Body.String())
		if w.Code != 200 {
			t.Error(w.Code)
		}
	})
	fmt.Println("59:", req.Body)

	if w.Code != 200 {
		t.Error(w.Code)
	}
}
