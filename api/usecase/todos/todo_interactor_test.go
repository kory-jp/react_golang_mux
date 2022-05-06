package usecase_test

import (
	"database/sql"
	"errors"
	"reflect"
	"strings"
	"testing"

	mock_transaction "github.com/kory-jp/react_golang_mux/api/usecase/transaction/mock"

	"github.com/stretchr/testify/assert"

	mock_usecase "github.com/kory-jp/react_golang_mux/api/usecase/todos/mock"

	"github.com/golang/mock/gomock"
	"github.com/kory-jp/react_golang_mux/api/domain"
	usecase "github.com/kory-jp/react_golang_mux/api/usecase/todos"
)

func TestAdd(t *testing.T) {
	var tx *sql.Tx
	// mockを作成
	inter, TodoRepository, Transaction := setMock(t)

	// テストケースを作成
	cases := []struct {
		name               string
		args               domain.Todo
		tags               []int
		prepareTrasMockFn  func(m *mock_transaction.MockSqlHandler)
		prepareStoreMockFn func(m *mock_usecase.MockTodoRepository, tx *sql.Tx, args domain.Todo)
		wantMess           usecase.TodoMessage
		wantErr            error
	}{
		{
			name: "create = success",
			args: domain.Todo{
				UserID:     1,
				Title:      "test title",
				Content:    "test content",
				ImagePath:  "testImg",
				IsFinished: false,
				Importance: 1,
				Urgency:    1,
			},
			tags: []int{1, 2, 3},
			prepareTrasMockFn: func(m *mock_transaction.MockSqlHandler) {
				m.EXPECT().DoInTx(gomock.Any()).Return(nil, nil).AnyTimes()
			},
			prepareStoreMockFn: func(m *mock_usecase.MockTodoRepository, tx *sql.Tx, args domain.Todo) {
				m.EXPECT().TransStore(tx, args).Return(int64(1), nil)
			},
			wantMess: usecase.TodoMessage{
				Message: "保存しました!",
			},
		},
		{
			name: "when UserID = nil, create = fail",
			args: domain.Todo{
				Title:      "test title",
				Content:    "test content",
				ImagePath:  "testImg",
				IsFinished: false,
				Importance: 1,
				Urgency:    1,
			},
			prepareTrasMockFn: func(m *mock_transaction.MockSqlHandler) {
				m.EXPECT().DoInTx(gomock.Any()).Return(nil, nil).AnyTimes()
			},
			prepareStoreMockFn: func(m *mock_usecase.MockTodoRepository, tx *sql.Tx, args domain.Todo) {
				m.EXPECT().TransStore(tx, args).Return(int64(0), nil).AnyTimes()
			},
			wantErr: errors.New("ユーザーIDは必須です。"),
		},
		{
			name: "when title = nil, create = fail",
			args: domain.Todo{
				UserID:     1,
				Content:    "test content",
				ImagePath:  "testImg",
				IsFinished: false,
				Importance: 1,
				Urgency:    1,
			},
			prepareTrasMockFn: func(m *mock_transaction.MockSqlHandler) {
				m.EXPECT().DoInTx(gomock.Any()).Return(nil, nil).AnyTimes()
			},
			prepareStoreMockFn: func(m *mock_usecase.MockTodoRepository, tx *sql.Tx, args domain.Todo) {
				m.EXPECT().TransStore(tx, args).Return(int64(0), nil).AnyTimes()
			},
			wantErr: errors.New("タイトルは必須です。"),
		},
		{
			name: "when title > 50, create = fail",
			args: domain.Todo{
				UserID:     1,
				Title:      strings.Repeat("t", 51),
				Content:    "test Content",
				ImagePath:  "testImg",
				IsFinished: false,
				Importance: 1,
				Urgency:    1,
			},
			prepareTrasMockFn: func(m *mock_transaction.MockSqlHandler) {
				m.EXPECT().DoInTx(gomock.Any()).Return(nil, nil).AnyTimes()
			},
			prepareStoreMockFn: func(m *mock_usecase.MockTodoRepository, tx *sql.Tx, args domain.Todo) {
				m.EXPECT().TransStore(tx, args).Return(int64(0), nil).AnyTimes()
			},
			wantErr: errors.New("タイトルは50文字未満の入力になります。"),
		},
		{
			name: "when memo > 1999, create = fail",
			args: domain.Todo{
				UserID:     1,
				Title:      "test title",
				Content:    strings.Repeat("c", 2000),
				ImagePath:  "testImg",
				IsFinished: false,
				Importance: 1,
				Urgency:    1,
			},
			prepareTrasMockFn: func(m *mock_transaction.MockSqlHandler) {
				m.EXPECT().DoInTx(gomock.Any()).Return(nil, nil).AnyTimes()
			},
			prepareStoreMockFn: func(m *mock_usecase.MockTodoRepository, tx *sql.Tx, args domain.Todo) {
				m.EXPECT().TransStore(tx, args).Return(int64(0), nil).AnyTimes()
			},
			wantErr: errors.New("メモは2000文字未満の入力になります。"),
		},
		{
			name: "when importance = nil, create = fail",
			args: domain.Todo{
				UserID:     1,
				Title:      "test title",
				Content:    "test content",
				ImagePath:  "testImg",
				IsFinished: false,
				Importance: 0,
				Urgency:    1,
			},
			prepareTrasMockFn: func(m *mock_transaction.MockSqlHandler) {
				m.EXPECT().DoInTx(gomock.Any()).Return(nil, nil).AnyTimes()
			},
			prepareStoreMockFn: func(m *mock_usecase.MockTodoRepository, tx *sql.Tx, args domain.Todo) {
				m.EXPECT().TransStore(tx, args).Return(int64(0), nil).AnyTimes()
			},
			wantErr: errors.New("重要度は必須です。"),
		},
		{
			name: "when importance > 2, create = fail",
			args: domain.Todo{
				UserID:     1,
				Title:      "test title",
				Content:    "test content",
				ImagePath:  "testImg",
				IsFinished: false,
				Importance: 3,
				Urgency:    1,
			},
			prepareTrasMockFn: func(m *mock_transaction.MockSqlHandler) {
				m.EXPECT().DoInTx(gomock.Any()).Return(nil, nil).AnyTimes()
			},
			prepareStoreMockFn: func(m *mock_usecase.MockTodoRepository, tx *sql.Tx, args domain.Todo) {
				m.EXPECT().TransStore(tx, args).Return(int64(0), nil).AnyTimes()
			},
			wantErr: errors.New("重要度に異常な値が入力されました"),
		},
		{
			name: "when urgency = nil, create = fail",
			args: domain.Todo{
				UserID:     1,
				Title:      "test title",
				Content:    "test content",
				ImagePath:  "testImg",
				IsFinished: false,
				Importance: 1,
				Urgency:    0,
			},
			prepareTrasMockFn: func(m *mock_transaction.MockSqlHandler) {
				m.EXPECT().DoInTx(gomock.Any()).Return(nil, nil).AnyTimes()
			},
			prepareStoreMockFn: func(m *mock_usecase.MockTodoRepository, tx *sql.Tx, args domain.Todo) {
				m.EXPECT().TransStore(tx, args).Return(int64(0), nil).AnyTimes()
			},
			wantErr: errors.New("緊急度は必須です。"),
		},
		{
			name: "urgency > 2, create = fail",
			args: domain.Todo{
				UserID:     1,
				Title:      "test title",
				Content:    "test content",
				ImagePath:  "testImg",
				IsFinished: false,
				Importance: 1,
				Urgency:    3,
			},
			prepareTrasMockFn: func(m *mock_transaction.MockSqlHandler) {
				m.EXPECT().DoInTx(gomock.Any()).Return(nil, nil).AnyTimes()
			},
			prepareStoreMockFn: func(m *mock_usecase.MockTodoRepository, tx *sql.Tx, args domain.Todo) {
				m.EXPECT().TransStore(tx, args).Return(int64(0), nil).AnyTimes()
			},
			wantErr: errors.New("緊急度に異常な値が入力されました"),
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepareTrasMockFn(Transaction)
			tt.prepareStoreMockFn(TodoRepository, tx, tt.args)
			mess, err := inter.Add(tt.args, tt.tags)
			if err == nil {
				assert.Equal(t, tt.wantMess, *mess)
			} else if err.Error() != tt.wantErr.Error() {
				assert.Equal(t, tt.wantErr, err)
			}
		})
	}
}

func TestTodos(t *testing.T) {
	inter, TodoRepository, _ := setMock(t)

	cases := []struct {
		name          string
		userId        int
		page          int
		prepareMockFn func(m *mock_usecase.MockTodoRepository, userId int, page int)
		wantTodos     domain.Todos
		wantSumPage   int
		wantErr       error
	}{
		{
			name:   "getTodos = success",
			userId: 1,
			page:   1,
			prepareMockFn: func(m *mock_usecase.MockTodoRepository, userId int, page int) {
				m.EXPECT().FindByUserId(userId, page).Return([]domain.Todo{{UserID: userId}}, 1.0, nil)
			},
			wantTodos: []domain.Todo{{UserID: 1}},
		},
		{
			name: "when userId = nil, getTodos = fail",
			page: 1,
			prepareMockFn: func(m *mock_usecase.MockTodoRepository, userId int, page int) {
				m.EXPECT().FindByUserId(0, page).Return([]domain.Todo{{UserID: 0}}, 0.0, nil).AnyTimes()
			},
			wantErr: errors.New("データ取得に失敗しました"),
		},
		{
			name:   "when page = nil, getTodos = fail",
			userId: 1,
			prepareMockFn: func(m *mock_usecase.MockTodoRepository, userId int, page int) {
				m.EXPECT().FindByUserId(userId, 0).Return([]domain.Todo{{UserID: userId}}, 0.0, nil).AnyTimes()
			},
			wantErr: errors.New("データ取得に失敗しました"),
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepareMockFn(TodoRepository, tt.userId, tt.page)
			todos, _, err := inter.Todos(tt.userId, tt.page)
			if err != nil {
				if err.Error() != tt.wantErr.Error() {
					t.Error("actual:", err, "want:", tt.wantErr)
					return
				}
				return
			}
			if !reflect.DeepEqual(todos, tt.wantTodos) {
				assert.Equal(t, tt.wantTodos, todos)
			}
		})
	}
}

func TestTodoByIdAndUserId(t *testing.T) {
	inter, TodoRepository, _ := setMock(t)

	cases := []struct {
		name          string
		id            int
		userId        int
		prepareMockFn func(m *mock_usecase.MockTodoRepository, id int, userId int)
		wantTodo      *domain.Todo
		wantErr       error
	}{
		{
			name:   "getTodo = success",
			id:     1,
			userId: 1,
			prepareMockFn: func(m *mock_usecase.MockTodoRepository, id int, userId int) {
				m.EXPECT().FindByIdAndUserId(id, userId).Return(&domain.Todo{ID: id, UserID: userId}, nil)
			},
			wantTodo: &domain.Todo{ID: 1, UserID: 1},
		},
		{
			name:   "when id = nil, getTodo = fail",
			userId: 1,
			prepareMockFn: func(m *mock_usecase.MockTodoRepository, id int, userId int) {
				m.EXPECT().FindByIdAndUserId(0, userId).Return(&domain.Todo{ID: 0, UserID: userId}, nil).AnyTimes()
			},
			wantErr: errors.New("データ取得に失敗しました"),
		},
		{
			name: "when userId = nil, getTodo = fail",
			id:   1,
			prepareMockFn: func(m *mock_usecase.MockTodoRepository, id int, userId int) {
				m.EXPECT().FindByIdAndUserId(id, 0).Return(&domain.Todo{ID: id, UserID: 0}, nil).AnyTimes()
			},
			wantErr: errors.New("データ取得に失敗しました"),
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepareMockFn(TodoRepository, tt.id, tt.userId)
			todo, err := inter.TodoByIdAndUserId(tt.id, tt.userId)
			if err != nil {
				if err.Error() != tt.wantErr.Error() {
					t.Error("actual:", err, "want:", tt.wantErr)
					return
				}
				return
			}
			if !reflect.DeepEqual(todo, tt.wantTodo) {
				assert.Equal(t, tt.wantTodo, todo)
			}
		})
	}
}

func TestSearch(t *testing.T) {
	inter, TodoRepository, _ := setMock(t)

	cases := []struct {
		name          string
		tagId         int
		userId        int
		args          domain.Todo
		page          int
		prepareMockFn func(m *mock_usecase.MockTodoRepository, tagId int, importanceScore int, urfgencyScore int, userId int, page int)
		wantTodos     domain.Todos
		wantErr       error
	}{
		{
			name:   "search = success",
			tagId:  1,
			userId: 1,
			args: domain.Todo{
				UserID:     1,
				Importance: 1,
				Urgency:    1,
			},
			page: 1,
			prepareMockFn: func(m *mock_usecase.MockTodoRepository, tagId int, importanceScore int, urfgencyScore int, userId int, page int) {
				m.EXPECT().Search(tagId, importanceScore, urfgencyScore, userId, page).Return([]domain.Todo{{UserID: userId, Importance: 1, Urgency: 1, Tags: []domain.Tag{{ID: tagId}}}}, float64(1), nil)
			},
			wantTodos: []domain.Todo{{UserID: 1, Importance: 1, Urgency: 1, Tags: []domain.Tag{{ID: 1}}}},
		},
		{
			// 検索条件(importance)に不適切な値(3以上の値)が入力された場合、データ取得失敗
			name:   "when importance > 2, search = fail",
			tagId:  1,
			userId: 1,
			args: domain.Todo{
				UserID:     1,
				Importance: 3,
				Urgency:    1,
			},
			page: 1,
			prepareMockFn: func(m *mock_usecase.MockTodoRepository, tagId int, importanceScore int, urfgencyScore int, userId int, page int) {
				m.EXPECT().Search(tagId, importanceScore, urfgencyScore, userId, page).Return([]domain.Todo{{UserID: userId, Importance: 1, Urgency: 1, Tags: []domain.Tag{{ID: tagId}}}}, float64(1), nil)
			},
			wantErr: errors.New("データ取得に失敗しました"),
		},
		{
			// 検索条件(urgency)に不適切な値(3以上の値)が入力された場合、データ取得失敗
			name:   "when urgency > 2, search = fail",
			tagId:  1,
			userId: 1,
			args: domain.Todo{
				UserID:     1,
				Importance: 1,
				Urgency:    3,
			},
			page: 1,
			prepareMockFn: func(m *mock_usecase.MockTodoRepository, tagId int, importanceScore int, urfgencyScore int, userId int, page int) {
				m.EXPECT().Search(tagId, importanceScore, urfgencyScore, userId, page).Return([]domain.Todo{{UserID: userId, Importance: 1, Urgency: 1, Tags: []domain.Tag{{ID: tagId}}}}, float64(1), nil)
			},
			wantErr: errors.New("データ取得に失敗しました"),
		},
		{
			// 検索条件(tagId,importance,urgency)に値が入力されていない場合、データ取得失敗
			name:   "when tagId = 0 && importance = 0 && urgency = 0, search =fail",
			tagId:  0,
			userId: 1,
			args: domain.Todo{
				UserID:     1,
				Importance: 0,
				Urgency:    0,
			},
			page: 1,
			prepareMockFn: func(m *mock_usecase.MockTodoRepository, tagId int, importanceScore int, urfgencyScore int, userId int, page int) {
				m.EXPECT().Search(tagId, importanceScore, urfgencyScore, userId, page).Return([]domain.Todo{{UserID: userId, Importance: 1, Urgency: 1, Tags: []domain.Tag{{ID: tagId}}}}, float64(1), nil)
			},
			wantErr: errors.New("データ取得に失敗しました"),
		},
		{
			name:   "when userId = 0, search = fail",
			tagId:  1,
			userId: 0,
			args: domain.Todo{
				UserID:     1,
				Importance: 1,
				Urgency:    1,
			},
			page: 1,
			prepareMockFn: func(m *mock_usecase.MockTodoRepository, tagId int, importanceScore int, urfgencyScore int, userId int, page int) {
				m.EXPECT().Search(tagId, importanceScore, urfgencyScore, userId, page).Return([]domain.Todo{{UserID: userId, Importance: 1, Urgency: 1, Tags: []domain.Tag{{ID: tagId}}}}, float64(1), nil)
			},
			wantErr: errors.New("データ取得に失敗しました"),
		},
		{
			name:   "when page = 0, search =fail",
			tagId:  1,
			userId: 1,
			page:   0,
			prepareMockFn: func(m *mock_usecase.MockTodoRepository, tagId int, importanceScore int, urfgencyScore int, userId int, page int) {
				m.EXPECT().Search(tagId, importanceScore, urfgencyScore, userId, page).Return([]domain.Todo{{UserID: userId, Importance: 1, Urgency: 1, Tags: []domain.Tag{{ID: tagId}}}}, float64(1), nil)
			},
			wantErr: errors.New("データ取得に失敗しました"),
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepareMockFn(TodoRepository, tt.tagId, tt.args.Importance, tt.args.Urgency, tt.userId, tt.page)
			todos, _, err := inter.Search(tt.tagId, tt.args.Importance, tt.args.Urgency, tt.userId, tt.page)
			if err != nil {
				if err.Error() != tt.wantErr.Error() {
					t.Error("actual:", err, "want:", tt.wantErr)
					return
				}
				return
			}
			if !reflect.DeepEqual(todos, tt.wantTodos) {
				assert.Equal(t, tt.wantTodos, todos)
			}
		})
	}
}

func TestUpdateTodo(t *testing.T) {
	inter, TodoRepository, Transaction := setMock(t)
	var tx *sql.Tx

	cases := []struct {
		name                   string
		args                   domain.Todo
		tags                   []int
		prepareTrasMockFn      func(m *mock_transaction.MockSqlHandler)
		prepareOverWriteMockFn func(m *mock_usecase.MockTodoRepository, tx *sql.Tx, args domain.Todo)
		wantMess               usecase.TodoMessage
		wantErr                error
	}{
		{
			name: "update = success",
			args: domain.Todo{
				UserID:     1,
				Title:      "test update title",
				Content:    "test update content",
				ImagePath:  "testUpdateImage",
				IsFinished: false,
				Importance: 1,
				Urgency:    1,
			},
			tags: []int{1, 2},
			prepareTrasMockFn: func(m *mock_transaction.MockSqlHandler) {
				m.EXPECT().DoInTx(gomock.Any()).Return(nil, nil).AnyTimes()
			},
			prepareOverWriteMockFn: func(m *mock_usecase.MockTodoRepository, tx *sql.Tx, args domain.Todo) {
				m.EXPECT().TransOverwrite(tx, args).Return(nil)
			},
			wantMess: usecase.TodoMessage{
				Message: "更新しました",
			},
		},
		{
			name: "when UserID = nil, update = fail",
			args: domain.Todo{
				Title:      "test title",
				Content:    "test content",
				ImagePath:  "testImg",
				IsFinished: false,
				Importance: 1,
				Urgency:    1,
			},
			tags: []int{1, 2},
			prepareTrasMockFn: func(m *mock_transaction.MockSqlHandler) {
				m.EXPECT().DoInTx(gomock.Any()).Return(nil, nil).AnyTimes()
			},
			prepareOverWriteMockFn: func(m *mock_usecase.MockTodoRepository, tx *sql.Tx, args domain.Todo) {
				m.EXPECT().TransOverwrite(tx, args).Return(nil).AnyTimes()
			},
			wantErr: errors.New("ユーザーIDは必須です。"),
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepareTrasMockFn(Transaction)
			tt.prepareOverWriteMockFn(TodoRepository, tx, tt.args)
			mess, err := inter.UpdateTodo(tt.args, tt.tags)
			if err == nil {
				if mess.Message != tt.wantMess.Message {
					assert.Equal(t, tt.wantMess, mess)
				}
			} else if err.Error() != tt.wantErr.Error() {
				assert.Equal(t, tt.wantErr, err)
			}
		})
	}
}

func TestIsFinishedTodo(t *testing.T) {
	inter, TodoRepository, _ := setMock(t)

	cases := []struct {
		name                  string
		id                    int
		userId                int
		todo                  domain.Todo
		prepareMockChangeBlFn func(m *mock_usecase.MockTodoRepository, id int, userId int, todo domain.Todo)
		prepareMockFindByFn   func(m *mock_usecase.MockTodoRepository, id int, userId int)
		wantMess              usecase.TodoMessage
		wantErr               error
	}{
		{
			name:   "isFinished = true, message = 完了しました",
			id:     1,
			userId: 1,
			todo: domain.Todo{
				ID:         1,
				UserID:     1,
				Title:      "test isFinished title",
				Content:    "test isFinished content",
				ImagePath:  "testIsFinishedImg",
				IsFinished: true,
			},
			prepareMockChangeBlFn: func(m *mock_usecase.MockTodoRepository, id int, userId int, todo domain.Todo) {
				m.EXPECT().ChangeBoolean(id, userId, todo).Return(nil)
			},
			prepareMockFindByFn: func(m *mock_usecase.MockTodoRepository, id int, userId int) {
				m.EXPECT().FindByIdAndUserId(id, userId).Return(&domain.Todo{ID: id, UserID: userId, IsFinished: true}, nil)
			},
			wantMess: usecase.TodoMessage{
				Message: "完了しました",
			},
		},
		{
			name:   "isFinished = false, message = 未完了の項目が追加されました",
			id:     1,
			userId: 1,
			todo: domain.Todo{
				ID:         1,
				UserID:     1,
				Title:      "test isFinished title",
				Content:    "test isFinished content",
				ImagePath:  "testIsFinishedImg",
				IsFinished: false,
			},
			prepareMockChangeBlFn: func(m *mock_usecase.MockTodoRepository, id int, userId int, todo domain.Todo) {
				m.EXPECT().ChangeBoolean(id, userId, todo).Return(nil)
			},
			prepareMockFindByFn: func(m *mock_usecase.MockTodoRepository, id int, userId int) {
				m.EXPECT().FindByIdAndUserId(id, userId).Return(&domain.Todo{ID: id, UserID: userId, IsFinished: false}, nil)
			},
			wantMess: usecase.TodoMessage{
				Message: "未完了の項目が追加されました",
			},
		},
		{
			name:   "idがnilの場合、isFinishedの更新の失敗",
			userId: 1,
			todo: domain.Todo{
				ID:         1,
				UserID:     1,
				Title:      "test isFinished title",
				Content:    "test isFinished content",
				ImagePath:  "testIsFinishedImg",
				IsFinished: false,
			},
			prepareMockChangeBlFn: func(m *mock_usecase.MockTodoRepository, id int, userId int, todo domain.Todo) {
				m.EXPECT().ChangeBoolean(0, userId, todo).Return(nil).AnyTimes()
			},
			prepareMockFindByFn: func(m *mock_usecase.MockTodoRepository, id int, userId int) {
				m.EXPECT().FindByIdAndUserId(0, userId).Return(&domain.Todo{ID: 0, UserID: userId, IsFinished: false}, nil).AnyTimes()
			},
			wantErr: errors.New("データ取得に失敗しました"),
		},
		{
			name: "when userId = nil, changeIsFinished = fail",
			id:   1,
			todo: domain.Todo{
				ID:         1,
				UserID:     1,
				Title:      "test isFinished title",
				Content:    "test isFinished content",
				ImagePath:  "testIsFinishedImg",
				IsFinished: false,
			},
			prepareMockChangeBlFn: func(m *mock_usecase.MockTodoRepository, id int, userId int, todo domain.Todo) {
				m.EXPECT().ChangeBoolean(id, userId, todo).Return(nil).AnyTimes()
			},
			prepareMockFindByFn: func(m *mock_usecase.MockTodoRepository, id int, userId int) {
				m.EXPECT().FindByIdAndUserId(id, 0).Return(&domain.Todo{ID: 1, UserID: 0, IsFinished: false}, nil).AnyTimes()
			},
			wantErr: errors.New("データ取得に失敗しました"),
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepareMockChangeBlFn(TodoRepository, tt.id, tt.userId, tt.todo)
			tt.prepareMockFindByFn(TodoRepository, tt.id, tt.userId)
			mess, err := inter.IsFinishedTodo(tt.id, tt.todo, tt.userId)
			if err == nil {
				if mess.Message != tt.wantMess.Message {
					assert.Equal(t, tt.wantMess, mess)
					return
				}
				return
			} else if err.Error() != tt.wantErr.Error() {
				assert.Equal(t, tt.wantErr, err)
			}
		})
	}
}

func TestDeleteTodo(t *testing.T) {
	inter, TodoRepository, _ := setMock(t)

	cases := []struct {
		name          string
		id            int
		userId        int
		prepareMockFn func(m *mock_usecase.MockTodoRepository, id int, userId int)
		wantMess      usecase.TodoMessage
		wantErr       error
	}{
		{
			name:   "delete = success",
			id:     1,
			userId: 1,
			prepareMockFn: func(m *mock_usecase.MockTodoRepository, id int, userId int) {
				m.EXPECT().Erasure(id, userId).Return(nil)
			},
			wantMess: usecase.TodoMessage{
				Message: "削除しました",
			},
		},
		{
			name:   "when id = nil, delete = fail",
			userId: 1,
			prepareMockFn: func(m *mock_usecase.MockTodoRepository, id int, userId int) {
				m.EXPECT().Erasure(0, userId).Return(nil).AnyTimes()
			},
			wantErr: errors.New("データ取得に失敗しました"),
		},
		{
			name: "when userId = nil, delete = fail",
			id:   1,
			prepareMockFn: func(m *mock_usecase.MockTodoRepository, id int, userId int) {
				m.EXPECT().Erasure(id, 0).Return(nil).AnyTimes()
			},
			wantErr: errors.New("データ取得に失敗しました"),
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepareMockFn(TodoRepository, tt.id, tt.userId)
			mess, err := inter.DeleteTodo(tt.id, tt.userId)
			if err == nil {
				if mess.Message != tt.wantMess.Message {
					assert.Equal(t, tt.wantMess, mess)
				}
			} else if err.Error() != tt.wantErr.Error() {
				assert.Equal(t, tt.wantErr, err)
			}
		})
	}
}

func TestDeleteTodoIndex(t *testing.T) {
	inter, TodoRepository, _ := setMock(t)

	cases := []struct {
		name                 string
		id                   int
		userId               int
		page                 int
		prepareMockErasureFn func(m *mock_usecase.MockTodoRepository, id int, userId int)
		prepareMockFindByFn  func(m *mock_usecase.MockTodoRepository, userId int, page int)
		wantTodos            domain.Todos
		wantSumPage          float64
		wantMess             usecase.TodoMessage
		wantErr              error
	}{
		{
			name:   "getTodos = success",
			id:     1,
			userId: 1,
			page:   1,
			prepareMockErasureFn: func(m *mock_usecase.MockTodoRepository, id int, userId int) {
				m.EXPECT().Erasure(id, userId).Return(nil)
			},
			prepareMockFindByFn: func(m *mock_usecase.MockTodoRepository, userId int, page int) {
				m.EXPECT().FindByUserId(userId, page).Return([]domain.Todo{{UserID: userId}}, 1.0, nil)
			},
			wantTodos:   []domain.Todo{{UserID: 1}},
			wantSumPage: 1.0,
			wantMess: usecase.TodoMessage{
				Message: "削除しました",
			},
		},
		{
			name:   "when id = nil, getTodo = fail",
			userId: 1,
			page:   1,
			prepareMockErasureFn: func(m *mock_usecase.MockTodoRepository, id int, userId int) {
				m.EXPECT().Erasure(0, userId).Return(nil).AnyTimes()
			},
			prepareMockFindByFn: func(m *mock_usecase.MockTodoRepository, userId int, page int) {
				m.EXPECT().FindByUserId(userId, page).Return([]domain.Todo{{UserID: userId}}, 1.0, nil).AnyTimes()
			},
			wantErr: errors.New("データ取得に失敗しました"),
		},
		{
			name: "when userID = nil, getTodo = fail",
			id:   1,
			page: 1,
			prepareMockErasureFn: func(m *mock_usecase.MockTodoRepository, id int, userId int) {
				m.EXPECT().Erasure(id, 0).Return(nil).AnyTimes()
			},
			prepareMockFindByFn: func(m *mock_usecase.MockTodoRepository, userId int, page int) {
				m.EXPECT().FindByUserId(0, page).Return([]domain.Todo{{UserID: 0}}, 0.0, nil).AnyTimes()
			},
			wantErr: errors.New("データ取得に失敗しました"),
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepareMockErasureFn(TodoRepository, tt.id, tt.userId)
			tt.prepareMockFindByFn(TodoRepository, tt.userId, tt.page)
			todos, sumPage, mess, err := inter.DeleteTodoInIndex(tt.id, tt.userId, tt.page)
			if err == nil {
				if !reflect.DeepEqual(todos, tt.wantTodos) {
					assert.Equal(t, tt.wantTodos, todos)
					return
				}
				if sumPage != tt.wantSumPage {
					assert.Equal(t, tt.wantSumPage, sumPage)
					return
				}
				if mess.Message != tt.wantMess.Message {
					assert.Equal(t, tt.wantMess, mess)
					return
				}
			} else if err.Error() != tt.wantErr.Error() {
				assert.Equal(t, tt.wantErr, err)
			}
		})
	}
}

// --- 各種Mockをインスタンス ---
func setMock(t *testing.T) (inter *usecase.TodoInteractor, TodoRepository *mock_usecase.MockTodoRepository, Transaction *mock_transaction.MockSqlHandler) {
	// mockを作成
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	TodoRepository = mock_usecase.NewMockTodoRepository(ctrl)
	Transaction = mock_transaction.NewMockSqlHandler(ctrl)
	inter = &usecase.TodoInteractor{}
	inter.TodoRepository = TodoRepository
	inter.Transaction = Transaction
	return
}
