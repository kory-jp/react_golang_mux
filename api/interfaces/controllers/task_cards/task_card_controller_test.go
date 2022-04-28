package controllers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	"github.com/kory-jp/react_golang_mux/api/domain"
	"github.com/stretchr/testify/assert"

	"github.com/golang/mock/gomock"
	mock_database "github.com/kory-jp/react_golang_mux/api/interfaces/mock"
	usecase "github.com/kory-jp/react_golang_mux/api/usecase/task_cards"
	mock_usecase "github.com/kory-jp/react_golang_mux/api/usecase/task_cards/mock"

	controllers "github.com/kory-jp/react_golang_mux/api/interfaces/controllers/task_cards"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
var response *controllers.Response

func TestCreate(t *testing.T) {
	ctrl, taskCardRepository := setMock(t)

	cases := []struct {
		name          string
		args          domain.TaskCard
		requestBody   bool
		prepareMockFn func(m *mock_usecase.MockTaskCardRepository, args domain.TaskCard)
		response      controllers.Response
	}{
		{
			name: "success",
			args: domain.TaskCard{
				UserID:  1,
				TodoID:  1,
				Title:   "test title",
				Purpose: "test purpose",
				Content: "test content",
				Memo:    "test memo",
			},
			requestBody: true,
			prepareMockFn: func(m *mock_usecase.MockTaskCardRepository, args domain.TaskCard) {
				m.EXPECT().Store(args).Return(nil)
			},
			response: controllers.Response{
				Status:  200,
				Message: "保存しました",
			},
		},
		{
			name: "when userId = 0, result = fail",
			args: domain.TaskCard{
				UserID:  0,
				TodoID:  1,
				Title:   "test title",
				Purpose: "test purpose",
				Content: "test content",
				Memo:    "test memo",
			},
			requestBody: true,
			prepareMockFn: func(m *mock_usecase.MockTaskCardRepository, args domain.TaskCard) {
				m.EXPECT().Store(args).Return(nil)
			},
			response: controllers.Response{
				Status:  401,
				Message: "ログインをしてください",
			},
		},
		{
			name: "when requestBody = nil, result = fail",
			args: domain.TaskCard{
				UserID:  1,
				TodoID:  1,
				Title:   "test title",
				Purpose: "test purpose",
				Content: "test content",
				Memo:    "test memo",
			},
			requestBody: false,
			prepareMockFn: func(m *mock_usecase.MockTaskCardRepository, args domain.TaskCard) {
				m.EXPECT().Store(args).Return(nil)
			},
			response: controllers.Response{
				Status:  400,
				Message: "データ取得に失敗しました",
			},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			var req *http.Request
			jsonArgs, _ := json.Marshal(tt.args)
			apiURL := "/api/taskcard/new"
			if tt.requestBody {
				req = httptest.NewRequest("POST", apiURL, bytes.NewBuffer(jsonArgs))
			} else {
				req = httptest.NewRequest("POST", apiURL, nil)
			}
			w := httptest.NewRecorder()

			SetSessionUserId(t, w, req, tt.args.UserID)
			tt.prepareMockFn(taskCardRepository, tt.args)

			ctrl.Create(w, req)
			buf, _ := ioutil.ReadAll(w.Body)
			json.Unmarshal(buf, &response)
			assert.Equal(t, tt.response.Status, response.Status)
			assert.Equal(t, tt.response.Message, response.Message)
		})
	}
}

func TestIndex(t *testing.T) {
	ctrl, taskCardRepository := setMock(t)

	cases := []struct {
		name          string
		loginUserId   int
		todoId        int
		page          int
		prepareMockFn func(m *mock_usecase.MockTaskCardRepository, todoIn int, userId int, page int)
		response      controllers.Response
	}{
		{
			name:        "success",
			loginUserId: 1,
			todoId:      1,
			page:        1,
			prepareMockFn: func(m *mock_usecase.MockTaskCardRepository, todoId int, userId int, page int) {
				m.EXPECT().FindByTodoIdAndUserId(todoId, userId, page).Return([]domain.TaskCard{{UserID: 1, TodoID: 1}}, 1.0, nil)
			},
			response: controllers.Response{
				Status:  200,
				Message: "タスクカード一覧取得",
			},
		},
		{
			name:        "when loginUserId = 0, resutl = fail",
			loginUserId: 0,
			todoId:      1,
			page:        1,
			prepareMockFn: func(m *mock_usecase.MockTaskCardRepository, todoId int, userId int, page int) {
				m.EXPECT().FindByTodoIdAndUserId(todoId, userId, page).Return(nil, 0.0, nil)
			},
			response: controllers.Response{
				Status:  401,
				Message: "ログインをしてください",
			},
		},
		{
			name:        "when todoId = 0, resutl = fail",
			loginUserId: 1,
			todoId:      0,
			page:        1,
			prepareMockFn: func(m *mock_usecase.MockTaskCardRepository, todoId int, userId int, page int) {
				m.EXPECT().FindByTodoIdAndUserId(todoId, userId, page).Return(nil, 0.0, nil)
			},
			response: controllers.Response{
				Status:  400,
				Message: "データ取得に失敗しました",
			},
		},
		{
			name:        "when page = 0, resutl = fail",
			loginUserId: 1,
			todoId:      1,
			page:        0,
			prepareMockFn: func(m *mock_usecase.MockTaskCardRepository, todoId int, userId int, page int) {
				m.EXPECT().FindByTodoIdAndUserId(todoId, userId, page).Return(nil, 0.0, nil)
			},
			response: controllers.Response{
				Status:  400,
				Message: "データ取得に失敗しました",
			},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			var req *http.Request
			apiURL := fmt.Sprintf("/api/todo/%s/taskcard?page=%s", strconv.Itoa(tt.todoId), strconv.Itoa(tt.page))
			req = httptest.NewRequest("GET", apiURL, nil)
			req = mux.SetURLVars(req, map[string]string{
				"id": strconv.Itoa(tt.todoId),
			})
			w := httptest.NewRecorder()
			SetSessionUserId(t, w, req, tt.loginUserId)
			tt.prepareMockFn(taskCardRepository, tt.todoId, tt.loginUserId, tt.page)

			ctrl.Index(w, req)
			buf, _ := ioutil.ReadAll(w.Body)
			json.Unmarshal(buf, &response)
			assert.Equal(t, tt.response.Status, response.Status)
			assert.Equal(t, tt.response.Message, response.Message)
		})
	}
}

func TestShow(t *testing.T) {
	ctrl, taskCardRepository := setMock(t)

	cases := []struct {
		name          string
		loginUserId   int
		todoId        int
		prepareMockFn func(m *mock_usecase.MockTaskCardRepository, todoIn int, loginUserId int)
		response      controllers.Response
	}{
		{
			name:        "success",
			loginUserId: 1,
			todoId:      1,
			prepareMockFn: func(m *mock_usecase.MockTaskCardRepository, todoId, loginUserId int) {
				m.EXPECT().FindByIdAndUserId(todoId, loginUserId).Return(&domain.TaskCard{ID: 1, UserID: 1}, nil)
			},
			response: controllers.Response{
				Status:  200,
				Message: "タスクカード詳細取得",
			},
		},
		{
			name:        "when loginUserId = 0, result = fail",
			loginUserId: 0,
			todoId:      1,
			prepareMockFn: func(m *mock_usecase.MockTaskCardRepository, todoId, loginUserId int) {
				m.EXPECT().FindByIdAndUserId(todoId, loginUserId).Return(&domain.TaskCard{ID: 1, UserID: 1}, nil)
			},
			response: controllers.Response{
				Status:  401,
				Message: "ログインをしてください",
			},
		},
		{
			name:        "when todoId = 0, result = fail",
			loginUserId: 1,
			todoId:      0,
			prepareMockFn: func(m *mock_usecase.MockTaskCardRepository, todoId, loginUserId int) {
				m.EXPECT().FindByIdAndUserId(todoId, loginUserId).Return(&domain.TaskCard{ID: 1, UserID: 1}, nil)
			},
			response: controllers.Response{
				Status:  400,
				Message: "データ取得に失敗しました",
			},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			var req *http.Request
			apiURL := fmt.Sprintf("/api/taskcard/%s", strconv.Itoa(tt.todoId))
			req = httptest.NewRequest("GET", apiURL, nil)
			req = mux.SetURLVars(req, map[string]string{
				"id": strconv.Itoa(tt.todoId),
			})
			w := httptest.NewRecorder()

			SetSessionUserId(t, w, req, tt.loginUserId)
			tt.prepareMockFn(taskCardRepository, tt.todoId, tt.loginUserId)

			ctrl.Show(w, req)
			buf, _ := ioutil.ReadAll(w.Body)
			json.Unmarshal(buf, &response)
			assert.Equal(t, tt.response.Status, response.Status)
			assert.Equal(t, tt.response.Message, response.Message)
		})
	}
}

func TestUpdate(t *testing.T) {
	ctrl, taskCardRepository := setMock(t)

	cases := []struct {
		name          string
		args          domain.TaskCard
		requestBody   bool
		prepareMockFn func(m *mock_usecase.MockTaskCardRepository, args domain.TaskCard)
		response      controllers.Response
	}{
		{
			name: "success",
			args: domain.TaskCard{
				ID:         1,
				UserID:     1,
				TodoID:     1,
				Title:      "update test title",
				Purpose:    "update test purpose",
				Content:    "update test content",
				Memo:       "update test memo",
				IsFinished: false,
			},
			requestBody: true,
			prepareMockFn: func(m *mock_usecase.MockTaskCardRepository, args domain.TaskCard) {
				m.EXPECT().Overwrite(args).Return(nil)
			},
			response: controllers.Response{
				Status:  200,
				Message: "更新しました",
			},
		},
		{
			name: "when UserID = 0, result = fail",
			args: domain.TaskCard{
				ID:         1,
				UserID:     0,
				TodoID:     1,
				Title:      "update test title",
				Purpose:    "update test purpose",
				Content:    "update test content",
				Memo:       "update test memo",
				IsFinished: false,
			},
			requestBody: true,
			prepareMockFn: func(m *mock_usecase.MockTaskCardRepository, args domain.TaskCard) {
				m.EXPECT().Overwrite(args).Return(nil)
			},
			response: controllers.Response{
				Status:  401,
				Message: "ログインをしてください",
			},
		},
		{
			name: "when ID = 0, result = fail",
			args: domain.TaskCard{
				ID:         0,
				UserID:     1,
				TodoID:     1,
				Title:      "update test title",
				Purpose:    "update test purpose",
				Content:    "update test content",
				Memo:       "update test memo",
				IsFinished: false,
			},
			requestBody: true,
			prepareMockFn: func(m *mock_usecase.MockTaskCardRepository, args domain.TaskCard) {
				m.EXPECT().Overwrite(args).Return(nil)
			},
			response: controllers.Response{
				Status:  400,
				Message: "データ取得に失敗しました",
			},
		},
		{
			name: "when requestBody = nil, result = fail",
			args: domain.TaskCard{
				ID:         1,
				UserID:     1,
				TodoID:     1,
				Title:      "update test title",
				Purpose:    "update test purpose",
				Content:    "update test content",
				Memo:       "update test memo",
				IsFinished: false,
			},
			requestBody: false,
			prepareMockFn: func(m *mock_usecase.MockTaskCardRepository, args domain.TaskCard) {
				m.EXPECT().Overwrite(args).Return(nil)
			},
			response: controllers.Response{
				Status:  400,
				Message: "データ取得に失敗しました",
			},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			var req *http.Request
			jsonArgs, _ := json.Marshal(tt.args)
			apiURL := fmt.Sprintf("/api/taskcard/%s", strconv.Itoa(tt.args.ID))
			if tt.requestBody {
				req = httptest.NewRequest("POST", apiURL, bytes.NewBuffer(jsonArgs))
			} else {
				req = httptest.NewRequest("POST", apiURL, nil)
			}
			req = mux.SetURLVars(req, map[string]string{
				"id": strconv.Itoa(tt.args.ID),
			})
			w := httptest.NewRecorder()

			SetSessionUserId(t, w, req, tt.args.UserID)
			tt.prepareMockFn(taskCardRepository, tt.args)

			ctrl.Update(w, req)
			buf, _ := ioutil.ReadAll(w.Body)
			json.Unmarshal(buf, &response)
			assert.Equal(t, tt.response.Status, response.Status)
			assert.Equal(t, tt.response.Message, response.Message)
		})
	}
}

func TestIsFinished(t *testing.T) {
	ctrl, taskCardRepository := setMock(t)

	cases := []struct {
		name                  string
		taskCardId            int
		loginUserId           int
		args                  domain.TaskCard
		requestBody           bool
		prepareMockChangeBlFn func(m *mock_usecase.MockTaskCardRepository, taskCardId int, loginUserId int, args domain.TaskCard)
		prepareMockFindByFn   func(m *mock_usecase.MockTaskCardRepository, taskCardId int, loginUserId int)
		response              controllers.Response
	}{
		{
			name:        "success",
			taskCardId:  1,
			loginUserId: 1,
			args: domain.TaskCard{
				IsFinished: true,
			},
			requestBody: true,
			prepareMockChangeBlFn: func(m *mock_usecase.MockTaskCardRepository, taskCardId int, loginUserId int, args domain.TaskCard) {
				m.EXPECT().ChangeBoolean(taskCardId, loginUserId, args).Return(nil)
			},
			prepareMockFindByFn: func(m *mock_usecase.MockTaskCardRepository, taskCardId int, loginUserId int) {
				m.EXPECT().FindByIdAndUserId(taskCardId, loginUserId).Return(&domain.TaskCard{IsFinished: true}, nil)
			},
			response: controllers.Response{
				Status:  200,
				Message: "完了しました",
			},
		},
		{
			name:        "when taskCardId = 0, result = fail",
			taskCardId:  0,
			loginUserId: 1,
			args: domain.TaskCard{
				IsFinished: true,
			},
			requestBody: true,
			prepareMockChangeBlFn: func(m *mock_usecase.MockTaskCardRepository, taskCardId int, loginUserId int, args domain.TaskCard) {
				m.EXPECT().ChangeBoolean(taskCardId, loginUserId, args).Return(nil)
			},
			prepareMockFindByFn: func(m *mock_usecase.MockTaskCardRepository, taskCardId int, loginUserId int) {
				m.EXPECT().FindByIdAndUserId(taskCardId, loginUserId).Return(&domain.TaskCard{IsFinished: true}, nil)
			},
			response: controllers.Response{
				Status:  400,
				Message: "データ取得に失敗しました",
			},
		},
		{
			name:        "when loginUserId = 0, result = fail",
			taskCardId:  1,
			loginUserId: 0,
			args: domain.TaskCard{
				IsFinished: true,
			},
			requestBody: true,
			prepareMockChangeBlFn: func(m *mock_usecase.MockTaskCardRepository, taskCardId int, loginUserId int, args domain.TaskCard) {
				m.EXPECT().ChangeBoolean(taskCardId, loginUserId, args).Return(nil)
			},
			prepareMockFindByFn: func(m *mock_usecase.MockTaskCardRepository, taskCardId int, loginUserId int) {
				m.EXPECT().FindByIdAndUserId(taskCardId, loginUserId).Return(&domain.TaskCard{IsFinished: true}, nil)
			},
			response: controllers.Response{
				Status:  401,
				Message: "ログインをしてください",
			},
		},
		{
			name:        "when requestBody = nil, result = fail",
			taskCardId:  1,
			loginUserId: 1,
			args: domain.TaskCard{
				IsFinished: true,
			},
			requestBody: false,
			prepareMockChangeBlFn: func(m *mock_usecase.MockTaskCardRepository, taskCardId int, loginUserId int, args domain.TaskCard) {
				m.EXPECT().ChangeBoolean(taskCardId, loginUserId, args).Return(nil)
			},
			prepareMockFindByFn: func(m *mock_usecase.MockTaskCardRepository, taskCardId int, loginUserId int) {
				m.EXPECT().FindByIdAndUserId(taskCardId, loginUserId).Return(&domain.TaskCard{IsFinished: true}, nil)
			},
			response: controllers.Response{
				Status:  400,
				Message: "データ取得に失敗しました",
			},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			var req *http.Request
			jsonArgs, _ := json.Marshal(tt.args)
			apiURL := fmt.Sprintf("/api/taskcard/isfinished/%s", strconv.Itoa(tt.taskCardId))
			if tt.requestBody {
				req = httptest.NewRequest("POST", apiURL, bytes.NewBuffer(jsonArgs))
			} else {
				req = httptest.NewRequest("POST", apiURL, nil)
			}
			req = mux.SetURLVars(req, map[string]string{
				"id": strconv.Itoa(tt.taskCardId),
			})
			w := httptest.NewRecorder()

			SetSessionUserId(t, w, req, tt.loginUserId)

			tt.prepareMockChangeBlFn(taskCardRepository, tt.taskCardId, tt.loginUserId, tt.args)
			tt.prepareMockFindByFn(taskCardRepository, tt.taskCardId, tt.loginUserId)

			ctrl.IsFinished(w, req)
			buf, _ := ioutil.ReadAll(w.Body)
			json.Unmarshal(buf, &response)
			assert.Equal(t, tt.response.Status, response.Status)
			assert.Equal(t, tt.response.Message, response.Message)
		})
	}
}

func TestDelete(t *testing.T) {
	ctrl, taskCardRepository := setMock(t)

	cases := []struct {
		name          string
		taskCardId    int
		loginUserId   int
		prepareMockFn func(m *mock_usecase.MockTaskCardRepository, taskCardId int, loginUserId int)
		response      controllers.Response
	}{
		{
			name:        "success",
			taskCardId:  1,
			loginUserId: 1,
			prepareMockFn: func(m *mock_usecase.MockTaskCardRepository, taskCardId int, loginUserId int) {
				m.EXPECT().Erasure(taskCardId, loginUserId).Return(nil)
			},
			response: controllers.Response{
				Status:  200,
				Message: "削除しました",
			},
		},
		{
			name:        "when taskCardId = 0, result = fail",
			taskCardId:  0,
			loginUserId: 1,
			prepareMockFn: func(m *mock_usecase.MockTaskCardRepository, taskCardId int, loginUserId int) {
				m.EXPECT().Erasure(taskCardId, loginUserId).Return(nil)
			},
			response: controllers.Response{
				Status:  400,
				Message: "データ取得に失敗しました",
			},
		},
		{
			name:        "when loginUserId = 0, result = fail",
			taskCardId:  1,
			loginUserId: 0,
			prepareMockFn: func(m *mock_usecase.MockTaskCardRepository, taskCardId int, loginUserId int) {
				m.EXPECT().Erasure(taskCardId, loginUserId).Return(nil)
			},
			response: controllers.Response{
				Status:  401,
				Message: "ログインをしてください",
			},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			apiURL := fmt.Sprintf("/api/taskcard/isfinished/%s", strconv.Itoa(tt.taskCardId))
			req := httptest.NewRequest("DELETE", apiURL, nil)
			req = mux.SetURLVars(req, map[string]string{
				"id": strconv.Itoa(tt.taskCardId),
			})
			w := httptest.NewRecorder()

			SetSessionUserId(t, w, req, tt.loginUserId)

			tt.prepareMockFn(taskCardRepository, tt.taskCardId, tt.loginUserId)

			ctrl.Delete(w, req)
			buf, _ := ioutil.ReadAll(w.Body)
			json.Unmarshal(buf, &response)
			assert.Equal(t, tt.response.Status, response.Status)
			assert.Equal(t, tt.response.Message, response.Message)
		})
	}
}

func setMock(t *testing.T) (ctrl *controllers.TaskCardController, taskCardRepository *mock_usecase.MockTaskCardRepository) {
	c := gomock.NewController(t)
	defer c.Finish()
	sqlhandler := mock_database.NewMockSqlHandler(c)
	taskCardRepository = mock_usecase.NewMockTaskCardRepository(c)
	ctrl = controllers.NewTaskCardController(sqlhandler)
	Interactor := usecase.TaskCardInteractor{}
	Interactor.TaskCardRepository = taskCardRepository
	ctrl.Interactor = Interactor
	return
}

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
