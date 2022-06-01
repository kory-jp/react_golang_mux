package controllers_test

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	mock_awsS3handlers "github.com/kory-jp/react_golang_mux/api/interfaces/controllers/mock"

	usecase "github.com/kory-jp/react_golang_mux/api/usecase/todos"
	mock_usecase "github.com/kory-jp/react_golang_mux/api/usecase/todos/mock"

	mock_transaction "github.com/kory-jp/react_golang_mux/api/usecase/transaction/mock"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/sessions"
	"github.com/kory-jp/react_golang_mux/api/domain"
	controllers "github.com/kory-jp/react_golang_mux/api/interfaces/controllers/todos"
	mock_database "github.com/kory-jp/react_golang_mux/api/interfaces/mock"
	"github.com/stretchr/testify/assert"
)

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
var response *controllers.Response

func TestCreate(t *testing.T) {
	// --- 各種mockをインスタンス ---
	ctrl, todoRepository, transaction := setMock(t)
	var tx *sql.Tx

	cases := []struct {
		name                   string
		args                   domain.Todo
		isImage                bool
		requestBody            bool
		prepareTrasMockFn      func(m *mock_transaction.MockSqlHandler)
		prepareRepoStoreMockFn func(m *mock_usecase.MockTodoRepository, tx *sql.Tx, args domain.Todo)
		response               controllers.Response
	}{
		{
			name: "when image = ImagePath, create = success",
			args: domain.Todo{
				UserID:     1,
				Title:      "test title",
				Content:    "test content",
				ImagePath:  "test.png",
				Importance: 1,
				Urgency:    1,
			},
			isImage:     true,
			requestBody: true,
			prepareTrasMockFn: func(m *mock_transaction.MockSqlHandler) {
				m.EXPECT().DoInTx(gomock.Any()).Return(nil, nil).AnyTimes()
			},
			prepareRepoStoreMockFn: func(m *mock_usecase.MockTodoRepository, tx *sql.Tx, args domain.Todo) {
				m.EXPECT().TransStore(tx, args).Return(int64(1), nil)
			},
			response: controllers.Response{
				Status:  200,
				Message: "保存しました",
			},
		},
		{
			name: "when image = nil, create = success",
			args: domain.Todo{
				UserID:     1,
				Title:      "test title",
				Content:    "test content",
				Importance: 1,
				Urgency:    1,
			},
			isImage:     false,
			requestBody: true,
			prepareTrasMockFn: func(m *mock_transaction.MockSqlHandler) {
				m.EXPECT().DoInTx(gomock.Any()).Return(nil, nil).AnyTimes()
			},
			prepareRepoStoreMockFn: func(m *mock_usecase.MockTodoRepository, tx *sql.Tx, args domain.Todo) {
				m.EXPECT().TransStore(tx, args).Return(int64(1), nil)
			},
			response: controllers.Response{
				Status:  200,
				Message: "保存しました",
			},
		},
		{
			name: "when userId = 0, create = fail",
			args: domain.Todo{
				UserID:     0,
				Title:      "test title",
				Content:    "test content",
				Importance: 1,
				Urgency:    1,
			},
			isImage:     false,
			requestBody: true,
			prepareTrasMockFn: func(m *mock_transaction.MockSqlHandler) {
				m.EXPECT().DoInTx(gomock.Any()).Return(nil, nil).AnyTimes()
			},
			prepareRepoStoreMockFn: func(m *mock_usecase.MockTodoRepository, tx *sql.Tx, args domain.Todo) {
				m.EXPECT().TransStore(tx, args).Return(int64(1), nil)
			},
			response: controllers.Response{
				Status:  401,
				Message: "ログインをしてください",
			},
		},
		{
			name: "when requestBody = nil, create = fail",
			args: domain.Todo{
				UserID:     1,
				Title:      "test title",
				Content:    "test content",
				Importance: 1,
				Urgency:    1,
			},
			isImage:     false,
			requestBody: false,
			prepareTrasMockFn: func(m *mock_transaction.MockSqlHandler) {
				m.EXPECT().DoInTx(gomock.Any()).Return(nil, nil).AnyTimes()
			},
			prepareRepoStoreMockFn: func(m *mock_usecase.MockTodoRepository, tx *sql.Tx, args domain.Todo) {
				m.EXPECT().TransStore(tx, args).Return(int64(1), nil)
			},
			response: controllers.Response{
				Status:  400,
				Message: "データ取得に失敗しました",
			},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			var buffer bytes.Buffer
			var req *http.Request
			writer := multipart.NewWriter(&buffer)
			// -- 各種フィールドに値を設定 ---
			setField(t, writer, tt.isImage, tt.args.ImagePath, tt.args.Title, tt.args.Content, tt.args.Importance, tt.args.Urgency)
			// --- フィールドの書き込みが終了後にClose ---
			writer.Close()

			if tt.requestBody {
				req = httptest.NewRequest("POST", "/api/new", &buffer)
			} else {
				req = httptest.NewRequest("POST", "/api/new", nil)
			}
			req.Header.Add("Content-Type", writer.FormDataContentType())
			w := httptest.NewRecorder()

			// --- sessionにユーザーIDを保存処理 ---
			SetSessionUserId(t, w, req, tt.args.UserID)

			// --- mock ---
			tt.prepareTrasMockFn(transaction)
			tt.prepareRepoStoreMockFn(todoRepository, tx, tt.args)

			// --- テスト実行 ---
			ctrl.Create(w, req)
			buf, _ := ioutil.ReadAll(w.Body)
			json.Unmarshal(buf, &response)
			assert.Equal(t, tt.response.Status, response.Status)
			assert.Equal(t, tt.response.Message, response.Message)
		})
	}
}

func TestIndex(t *testing.T) {
	ctrl, todoRepository, _ := setMock(t)

	cases := []struct {
		name                          string
		loginUserId                   int
		nowPage                       int
		prepareRepoFindByUserIdMockFn func(m *mock_usecase.MockTodoRepository, userId int, page int)
		response                      controllers.Response
	}{
		{
			name:        "getTodos = success",
			loginUserId: 1,
			nowPage:     1,
			prepareRepoFindByUserIdMockFn: func(m *mock_usecase.MockTodoRepository, userId int, page int) {
				m.EXPECT().FindByUserId(userId, page).Return([]domain.Todo{{UserID: userId}}, 1.0, nil)
			},
			response: controllers.Response{
				Status:  200,
				Message: "Todo一覧取得",
			},
		},
		{
			name:        "when userId = 0, getTodos = fail",
			loginUserId: 0,
			nowPage:     1,
			prepareRepoFindByUserIdMockFn: func(m *mock_usecase.MockTodoRepository, userId int, page int) {
				m.EXPECT().FindByUserId(userId, page).Return([]domain.Todo{{UserID: userId}}, 1.0, nil)
			},
			response: controllers.Response{
				Status:  401,
				Message: "ログインをしてください",
			},
		},
		{
			name:        "when nowPage = 0, getTodos = fail",
			loginUserId: 1,
			nowPage:     0,
			prepareRepoFindByUserIdMockFn: func(m *mock_usecase.MockTodoRepository, userId int, page int) {
				m.EXPECT().FindByUserId(userId, page).Return([]domain.Todo{{UserID: userId}}, 1.0, nil)
			},
			response: controllers.Response{
				Status:  400,
				Message: "データ取得に失敗しました",
			},
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

			tt.prepareRepoFindByUserIdMockFn(todoRepository, tt.loginUserId, tt.nowPage)

			ctrl.Index(w, req)
			buf, _ := ioutil.ReadAll(w.Body)
			json.Unmarshal(buf, &response)
			assert.Equal(t, tt.response.Status, response.Status)
			assert.Equal(t, tt.response.Message, response.Message)
		})
	}
}

func TestShow(t *testing.T) {
	ctrl, todoRepository, _ := setMock(t)

	cases := []struct {
		name                               string
		todoId                             int
		loginUserId                        int
		prepareRepoFindByIdAndUserIdMockFn func(m *mock_usecase.MockTodoRepository, id int, userId int)
		response                           controllers.Response
	}{
		{
			name:        "getTodo = success",
			todoId:      1,
			loginUserId: 1,
			prepareRepoFindByIdAndUserIdMockFn: func(m *mock_usecase.MockTodoRepository, id int, userId int) {
				m.EXPECT().FindByIdAndUserId(id, userId).Return(&domain.Todo{ID: id, UserID: userId}, nil)
			},
			response: controllers.Response{
				Status:  200,
				Message: "投稿詳細取得",
			},
		},
		{
			name:        "when todoId = 0, getTodo = fail",
			todoId:      0,
			loginUserId: 1,
			prepareRepoFindByIdAndUserIdMockFn: func(m *mock_usecase.MockTodoRepository, id int, userId int) {
				m.EXPECT().FindByIdAndUserId(id, userId).Return(&domain.Todo{ID: id, UserID: userId}, nil)
			},
			response: controllers.Response{
				Status:  400,
				Message: "データ取得に失敗しました",
			},
		},
		{
			name:        "when userId = 0, getTodo = fail",
			todoId:      1,
			loginUserId: 0,
			prepareRepoFindByIdAndUserIdMockFn: func(m *mock_usecase.MockTodoRepository, id int, userId int) {
				m.EXPECT().FindByIdAndUserId(id, userId).Return(&domain.Todo{ID: id, UserID: userId}, nil)
			},
			response: controllers.Response{
				Status:  401,
				Message: "ログインをしてください",
			},
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

			tt.prepareRepoFindByIdAndUserIdMockFn(todoRepository, tt.todoId, tt.loginUserId)

			ctrl.Show(w, req)
			buf, _ := ioutil.ReadAll(w.Body)
			json.Unmarshal(buf, &response)
			assert.Equal(t, tt.response.Status, response.Status)
			assert.Equal(t, tt.response.Message, response.Message)
		})
	}
}

func TestSearch(t *testing.T) {
	ctrl, todoRepository, _ := setMock(t)

	cases := []struct {
		name                         string
		tagId                        int
		args                         domain.Todo
		loginUserId                  int
		nowPage                      int
		apiURL                       string
		prepareRepoFindByTagIdMockFn func(m *mock_usecase.MockTodoRepository, tagId int, importanceScore int, urfgencyScore int, userId int, page int)
		response                     controllers.Response
	}{
		{
			// 検索可能なURLが取得された場合、データ取得に成功
			name:  "current URL, search = success",
			tagId: 1,
			args: domain.Todo{
				UserID:     1,
				Importance: 1,
				Urgency:    1,
			},
			loginUserId: 1,
			nowPage:     1,
			apiURL:      "/api/todos/search?tagId=1&importance=1&urgency=1&page=1",
			prepareRepoFindByTagIdMockFn: func(m *mock_usecase.MockTodoRepository, tagId int, importanceScore int, urfgencyScore int, userId int, page int) {
				m.EXPECT().Search(tagId, importanceScore, urfgencyScore, userId, page).Return([]domain.Todo{{UserID: userId, Importance: 1, Urgency: 1, Tags: []domain.Tag{{ID: tagId}}}}, float64(1), nil)
			},
			response: controllers.Response{
				Status:  200,
				Message: "タグ検索成功",
			},
		},
		{
			// URLにtag情報が存在しない場合、データ取得に失敗
			name:  "not include tagInfo at URL, search = fail",
			tagId: 1,
			args: domain.Todo{
				UserID:     1,
				Importance: 1,
				Urgency:    1,
			},
			loginUserId: 1,
			nowPage:     1,
			apiURL:      "/api/todos/search?importance=1&urgency=1&page=1",
			prepareRepoFindByTagIdMockFn: func(m *mock_usecase.MockTodoRepository, tagId int, importanceScore int, urfgencyScore int, userId int, page int) {
				m.EXPECT().Search(tagId, importanceScore, urfgencyScore, userId, page).Return([]domain.Todo{{UserID: userId, Importance: 1, Urgency: 1, Tags: []domain.Tag{{ID: tagId}}}}, float64(1), nil)
			},
			response: controllers.Response{
				Status:  400,
				Message: "データ取得に失敗しました",
			},
		},
		{
			// URLにimportance情報が存在しない場合、データ取得に失敗
			name:  "not include importanceInfo at URL, search = fail",
			tagId: 1,
			args: domain.Todo{
				UserID:     1,
				Importance: 1,
				Urgency:    1,
			},
			loginUserId: 1,
			nowPage:     1,
			apiURL:      "/api/todos/search?tagId=1&urgency=1&page=1",
			prepareRepoFindByTagIdMockFn: func(m *mock_usecase.MockTodoRepository, tagId int, importanceScore int, urfgencyScore int, userId int, page int) {
				m.EXPECT().Search(tagId, importanceScore, urfgencyScore, userId, page).Return([]domain.Todo{{UserID: userId, Importance: 1, Urgency: 1, Tags: []domain.Tag{{ID: tagId}}}}, float64(1), nil)
			},
			response: controllers.Response{
				Status:  400,
				Message: "データ取得に失敗しました",
			},
		},
		{
			// URLにurgency情報が存在しない場合、データ取得に失敗
			name:  "not include urgencyInfo at URL, search = fail",
			tagId: 1,
			args: domain.Todo{
				UserID:     1,
				Importance: 1,
				Urgency:    1,
			},
			loginUserId: 1,
			nowPage:     1,
			apiURL:      "/api/todos/search?importance=1&page=1",
			prepareRepoFindByTagIdMockFn: func(m *mock_usecase.MockTodoRepository, tagId int, importanceScore int, urfgencyScore int, userId int, page int) {
				m.EXPECT().Search(tagId, importanceScore, urfgencyScore, userId, page).Return([]domain.Todo{{UserID: userId, Importance: 1, Urgency: 1, Tags: []domain.Tag{{ID: tagId}}}}, float64(1), nil)
			},
			response: controllers.Response{
				Status:  400,
				Message: "データ取得に失敗しました",
			},
		},
		{
			// URLにpage情報が存在しない場合、データ取得に失敗
			name:  "not include pageInfo at URL, search = fail",
			tagId: 1,
			args: domain.Todo{
				UserID:     1,
				Importance: 1,
				Urgency:    1,
			},
			loginUserId: 1,
			nowPage:     1,
			apiURL:      "/api/todos/search?importance=1&urgency=1",
			prepareRepoFindByTagIdMockFn: func(m *mock_usecase.MockTodoRepository, tagId int, importanceScore int, urfgencyScore int, userId int, page int) {
				m.EXPECT().Search(tagId, importanceScore, urfgencyScore, userId, page).Return([]domain.Todo{{UserID: userId, Importance: 1, Urgency: 1, Tags: []domain.Tag{{ID: tagId}}}}, float64(1), nil)
			},
			response: controllers.Response{
				Status:  400,
				Message: "データ取得に失敗しました",
			},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			var buffer bytes.Buffer
			writer := multipart.NewWriter(&buffer)
			writer.Close()
			req := httptest.NewRequest("GET", tt.apiURL, &buffer)
			req.Header.Add("Content-Type", writer.FormDataContentType())
			w := httptest.NewRecorder()
			SetSessionUserId(t, w, req, tt.loginUserId)

			tt.prepareRepoFindByTagIdMockFn(todoRepository, tt.tagId, tt.args.Importance, tt.args.Urgency, tt.loginUserId, tt.nowPage)

			ctrl.Search(w, req)
			buf, _ := ioutil.ReadAll(w.Body)
			json.Unmarshal(buf, &response)
			assert.Equal(t, tt.response.Status, response.Status)
			assert.Equal(t, tt.response.Message, response.Message)
		})
	}
}

func TestUpdate(t *testing.T) {
	ctrl, todoRepository, transaction := setMock(t)
	var tx *sql.Tx

	cases := []struct {
		name                       string
		args                       domain.Todo
		loginUserId                int
		isImage                    bool
		requestBody                bool
		prepareTrasMockFn          func(m *mock_transaction.MockSqlHandler)
		prepareRepoOverwriteMockFn func(m *mock_usecase.MockTodoRepository, tx *sql.Tx)
		response                   controllers.Response
	}{
		{
			name: "when image = ImagePath, update = success",
			args: domain.Todo{
				ID:         1,
				UserID:     1,
				Title:      "test title",
				Content:    "test content",
				ImagePath:  "test.png",
				Importance: 1,
				Urgency:    1,
			},
			loginUserId: 1,
			isImage:     true,
			requestBody: true,
			prepareTrasMockFn: func(m *mock_transaction.MockSqlHandler) {
				m.EXPECT().DoInTx(gomock.Any()).Return(nil, nil).AnyTimes()
			},
			prepareRepoOverwriteMockFn: func(m *mock_usecase.MockTodoRepository, tx *sql.Tx) {
				m.EXPECT().TransOverwrite(tx, gomock.Any()).Return(nil)
			},
			response: controllers.Response{
				Status:  200,
				Message: "更新しました",
			},
		},
		{
			name: "when image = nil, update = success",
			args: domain.Todo{
				ID:         1,
				UserID:     1,
				Title:      "test title",
				Content:    "test content",
				ImagePath:  "test.png",
				Importance: 1,
				Urgency:    1,
			},
			loginUserId: 1,
			isImage:     false,
			requestBody: true,
			prepareTrasMockFn: func(m *mock_transaction.MockSqlHandler) {
				m.EXPECT().DoInTx(gomock.Any()).Return(nil, nil).AnyTimes()
			},
			prepareRepoOverwriteMockFn: func(m *mock_usecase.MockTodoRepository, tx *sql.Tx) {
				m.EXPECT().TransOverwrite(tx, gomock.Any()).Return(nil)
			},
			response: controllers.Response{
				Status:  200,
				Message: "更新しました",
			},
		},
		{
			name: "when loginUserId = 0, update = fail",
			args: domain.Todo{
				ID:         1,
				UserID:     1,
				Title:      "test title",
				Content:    "test content",
				ImagePath:  "test.png",
				Importance: 1,
				Urgency:    1,
			},
			loginUserId: 0,
			isImage:     false,
			requestBody: true,
			prepareTrasMockFn: func(m *mock_transaction.MockSqlHandler) {
				m.EXPECT().DoInTx(gomock.Any()).Return(nil, nil).AnyTimes()
			},
			prepareRepoOverwriteMockFn: func(m *mock_usecase.MockTodoRepository, tx *sql.Tx) {
				m.EXPECT().TransOverwrite(tx, gomock.Any()).Return(nil)
			},
			response: controllers.Response{
				Status:  401,
				Message: "ログインをしてください",
			},
		},
		{
			name: "when requestBody = nil, update = fail",
			args: domain.Todo{
				UserID:     1,
				Title:      "test title",
				Content:    "test content",
				Importance: 1,
				Urgency:    1,
			},
			isImage:     false,
			requestBody: false,
			prepareTrasMockFn: func(m *mock_transaction.MockSqlHandler) {
				m.EXPECT().DoInTx(gomock.Any()).Return(nil, nil).AnyTimes()
			},
			prepareRepoOverwriteMockFn: func(m *mock_usecase.MockTodoRepository, tx *sql.Tx) {
				m.EXPECT().TransOverwrite(tx, gomock.Any()).Return(nil)
			},
			response: controllers.Response{
				Status:  400,
				Message: "データ取得に失敗しました",
			},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			var buffer bytes.Buffer
			var req *http.Request
			writer := multipart.NewWriter(&buffer)
			// -- 各種フィールドに値を設定 ---
			setField(t, writer, tt.isImage, tt.args.ImagePath, tt.args.Title, tt.args.Content, tt.args.Importance, tt.args.Urgency)
			// --- フィールドの書き込みが終了後にClose ---
			writer.Close()
			// ---
			apiURL := "/api/todos/update/" + strconv.Itoa(tt.args.ID)
			// req := httptest.NewRequest("POST", apiURL, &buffer)
			if tt.requestBody {
				req = httptest.NewRequest("POST", apiURL, &buffer)
			} else {
				req = httptest.NewRequest("POST", apiURL, nil)
			}
			req.Header.Add("Content-Type", writer.FormDataContentType())
			w := httptest.NewRecorder()

			// --- sessionにユーザーIDを保存処理 ---
			SetSessionUserId(t, w, req, tt.loginUserId)

			// --- mock ---
			tt.prepareTrasMockFn(transaction)
			tt.prepareRepoOverwriteMockFn(todoRepository, tx)

			// --- テスト実行 ---
			ctrl.Update(w, req)
			buf, _ := ioutil.ReadAll(w.Body)
			json.Unmarshal(buf, &response)
			assert.Equal(t, tt.response.Status, response.Status)
			assert.Equal(t, tt.response.Message, response.Message)
		})
	}
}

func TestIsFinished(t *testing.T) {
	ctrl, todoRepository, _ := setMock(t)

	cases := []struct {
		name                  string
		todoId                int
		args                  domain.Todo
		loginUserId           int
		prepareMockChangeBlFn func(m *mock_usecase.MockTodoRepository, id int, userId int, todo domain.Todo)
		prepareMockFindByFn   func(m *mock_usecase.MockTodoRepository, id int, userId int)
		response              controllers.Response
	}{
		{
			name:   "changeIsFinished = success",
			todoId: 1,
			args: domain.Todo{
				IsFinished: true,
			},
			loginUserId: 1,
			prepareMockChangeBlFn: func(m *mock_usecase.MockTodoRepository, id int, userId int, todo domain.Todo) {
				m.EXPECT().ChangeBoolean(id, userId, todo).Return(nil)
			},
			prepareMockFindByFn: func(m *mock_usecase.MockTodoRepository, id int, userId int) {
				m.EXPECT().FindByIdAndUserId(id, userId).Return(&domain.Todo{ID: id, UserID: userId, IsFinished: true}, nil)
			},
			response: controllers.Response{
				Status:  200,
				Message: "完了しました",
			},
		},
		{
			name:   "when todoId = 0, changeIsFinished = fail",
			todoId: 0,
			args: domain.Todo{
				IsFinished: true,
			},
			loginUserId: 1,
			prepareMockChangeBlFn: func(m *mock_usecase.MockTodoRepository, id int, userId int, todo domain.Todo) {
				m.EXPECT().ChangeBoolean(id, userId, todo).Return(nil)
			},
			prepareMockFindByFn: func(m *mock_usecase.MockTodoRepository, id int, userId int) {
				m.EXPECT().FindByIdAndUserId(id, userId).Return(&domain.Todo{ID: id, UserID: userId, IsFinished: true}, nil)
			},
			response: controllers.Response{
				Status:  400,
				Message: "データ取得に失敗しました",
			},
		},
		{
			name:   "when loginUserId = 0, changeIsFinished = fail",
			todoId: 1,
			args: domain.Todo{
				IsFinished: true,
			},
			loginUserId: 0,
			prepareMockChangeBlFn: func(m *mock_usecase.MockTodoRepository, id int, userId int, todo domain.Todo) {
				m.EXPECT().ChangeBoolean(id, userId, todo).Return(nil)
			},
			prepareMockFindByFn: func(m *mock_usecase.MockTodoRepository, id int, userId int) {
				m.EXPECT().FindByIdAndUserId(id, userId).Return(&domain.Todo{ID: id, UserID: userId, IsFinished: true}, nil)
			},
			response: controllers.Response{
				Status:  401,
				Message: "ログインをしてください",
			},
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
			tt.prepareMockChangeBlFn(todoRepository, tt.todoId, tt.loginUserId, tt.args)
			tt.prepareMockFindByFn(todoRepository, tt.todoId, tt.loginUserId)

			// --- テスト実行 ---
			ctrl.IsFinished(w, req)
			buf, _ := ioutil.ReadAll(w.Body)
			json.Unmarshal(buf, &response)
			assert.Equal(t, tt.response.Status, response.Status)
			assert.Equal(t, tt.response.Message, response.Message)
		})
	}
}

func TestDelete(t *testing.T) {
	ctrl, todoRepository, _ := setMock(t)

	cases := []struct {
		name                     string
		todoId                   int
		loginUserId              int
		prepareRepoErasureMockFn func(m *mock_usecase.MockTodoRepository, id int, userId int)
		response                 controllers.Response
	}{
		{
			name:        "delete = success",
			todoId:      1,
			loginUserId: 1,
			prepareRepoErasureMockFn: func(m *mock_usecase.MockTodoRepository, id int, userId int) {
				m.EXPECT().Erasure(id, userId).Return(nil)
			},
			response: controllers.Response{
				Status:  200,
				Message: "削除しました",
			},
		},
		{
			name:        "when todoId = 0, delete = fail",
			todoId:      0,
			loginUserId: 1,
			prepareRepoErasureMockFn: func(m *mock_usecase.MockTodoRepository, id int, userId int) {
				m.EXPECT().Erasure(id, userId).Return(nil)
			},
			response: controllers.Response{
				Status:  400,
				Message: "データ取得に失敗しました",
			},
		},
		{
			name:        "when loginUserId = 0, delete = fail",
			todoId:      1,
			loginUserId: 0,
			prepareRepoErasureMockFn: func(m *mock_usecase.MockTodoRepository, id int, userId int) {
				m.EXPECT().Erasure(id, userId).Return(nil)
			},
			response: controllers.Response{
				Status:  401,
				Message: "ログインをしてください",
			},
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
			tt.prepareRepoErasureMockFn(todoRepository, tt.todoId, tt.loginUserId)

			// --- テスト実行 ---
			ctrl.Delete(w, req)
			buf, _ := ioutil.ReadAll(w.Body)
			json.Unmarshal(buf, &response)
			assert.Equal(t, tt.response.Status, response.Status)
			assert.Equal(t, tt.response.Message, response.Message)
		})
	}
}

func TestDeleteIndex(t *testing.T) {
	ctrl, todoRepository, _ := setMock(t)

	cases := []struct {
		name                     string
		todoId                   int
		loginUserId              int
		page                     int
		prepareRepoErasureMockFn func(m *mock_usecase.MockTodoRepository, id int, userId int)
		prepareRepoFindByMockFn  func(m *mock_usecase.MockTodoRepository, userId int, page int)
		response                 controllers.Response
	}{
		{
			name:        "getTodos = success",
			todoId:      2,
			loginUserId: 1,
			page:        1,
			prepareRepoErasureMockFn: func(m *mock_usecase.MockTodoRepository, id int, userId int) {
				m.EXPECT().Erasure(id, userId).Return(nil)
			},
			prepareRepoFindByMockFn: func(m *mock_usecase.MockTodoRepository, userId int, page int) {
				m.EXPECT().FindByUserId(userId, page).Return([]domain.Todo{{UserID: userId}}, 1.0, nil)
			},
			response: controllers.Response{
				Status:  200,
				Message: "削除しました",
			},
		},
		{
			name:        "when todoId = 0, getTodos = fail",
			todoId:      0,
			loginUserId: 1,
			page:        1,
			prepareRepoErasureMockFn: func(m *mock_usecase.MockTodoRepository, id int, userId int) {
				m.EXPECT().Erasure(id, userId).Return(nil)
			},
			prepareRepoFindByMockFn: func(m *mock_usecase.MockTodoRepository, userId int, page int) {
				m.EXPECT().FindByUserId(userId, page).Return([]domain.Todo{{UserID: userId}}, 1.0, nil)
			},
			response: controllers.Response{
				Status:  400,
				Message: "データ取得に失敗しました",
			},
		},
		{
			name:        "when loginUserId = 0, getTodos = fail",
			todoId:      1,
			loginUserId: 0,
			page:        1,
			prepareRepoErasureMockFn: func(m *mock_usecase.MockTodoRepository, id int, userId int) {
				m.EXPECT().Erasure(id, userId).Return(nil)
			},
			prepareRepoFindByMockFn: func(m *mock_usecase.MockTodoRepository, userId int, page int) {
				m.EXPECT().FindByUserId(userId, page).Return([]domain.Todo{{UserID: userId}}, 1.0, nil)
			},
			response: controllers.Response{
				Status:  401,
				Message: "ログインをしてください",
			},
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
			tt.prepareRepoErasureMockFn(todoRepository, tt.todoId, tt.loginUserId)
			tt.prepareRepoFindByMockFn(todoRepository, tt.loginUserId, tt.page)

			// --- テスト実行 ---
			ctrl.DeleteInIndex(w, req)
			buf, _ := ioutil.ReadAll(w.Body)
			json.Unmarshal(buf, &response)
			assert.Equal(t, tt.response.Status, response.Status)
			assert.Equal(t, tt.response.Message, response.Message)
		})
	}
}

// --- 各種Mockのインスタンス処理 ---
func setMock(t *testing.T) (ctrl *controllers.TodoController, todoRepository *mock_usecase.MockTodoRepository, transaction *mock_transaction.MockSqlHandler) {
	c := gomock.NewController(t)
	defer c.Finish()
	// --- api/interfaces/database/sqlhandlerのモック ---
	sqlhandler := mock_database.NewMockSqlHandler(c)
	awsS3handler := mock_awsS3handlers.NewMockS3(c)
	transaction = mock_transaction.NewMockSqlHandler(c)
	todoRepository = mock_usecase.NewMockTodoRepository(c)
	ctrl = controllers.NewTodoController(sqlhandler, awsS3handler)
	Interactor := usecase.TodoInteractor{}
	Interactor.Transaction = transaction
	Interactor.TodoRepository = todoRepository
	ctrl.Interactor = Interactor
	return
}

// --- todo投稿における各種にフィールドに値をセット ---
func setField(t *testing.T, writer *multipart.Writer, isImage bool, imagePath string, title string, content string, importance int, urgency int) {
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

	// --- importance ---
	importanceWriter, err := writer.CreateFormField("importance")
	if err != nil {
		t.Fatalf("Failed to create file writer. %s", err)
	}
	importanceWriter.Write([]byte(strconv.Itoa(importance)))

	// --- importance ---
	urgencyWriter, err := writer.CreateFormField("urgency")
	if err != nil {
		t.Fatalf("Failed to create file writer. %s", err)
	}
	urgencyWriter.Write([]byte(strconv.Itoa(urgency)))
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
