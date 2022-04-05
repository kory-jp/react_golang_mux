package usecase_test

import (
	"errors"
	"reflect"
	"strings"
	"testing"

	mock_usecase "github.com/kory-jp/react_golang_mux/api/usecase/todo/mock"

	"github.com/golang/mock/gomock"
	"github.com/kory-jp/react_golang_mux/api/domain"
	usecase "github.com/kory-jp/react_golang_mux/api/usecase/todo"
)

func TestAdd(t *testing.T) {
	// mockを作成
	inter, TodoRepository := setMock(t)

	// テストケースを作成
	cases := []struct {
		name          string
		args          domain.Todo
		prepareMockFn func(m *mock_usecase.MockTodoRepository, args domain.Todo)
		wantMess      usecase.TodoMessage
		wantErr       error
	}{
		{
			name: "必須項目が入力された場合、データ保存成功",
			args: domain.Todo{
				UserID:     1,
				Title:      "test title",
				Content:    "test content",
				ImagePath:  "testImg",
				IsFinished: false,
			},
			prepareMockFn: func(m *mock_usecase.MockTodoRepository, args domain.Todo) {
				m.EXPECT().Store(args).Return(nil)
			},
			wantMess: usecase.TodoMessage{
				Message: "保存しました",
			},
		},
		{
			name: "ユーザーIDがnilの場合、エラー返却",
			args: domain.Todo{
				Title:      "test title",
				Content:    "test content",
				ImagePath:  "testImg",
				IsFinished: false,
			},
			prepareMockFn: func(m *mock_usecase.MockTodoRepository, args domain.Todo) {
				m.EXPECT().Store(args).Return(nil).AnyTimes()
			},
			wantErr: errors.New("ユーザーIDは必須です。"),
		},
		{
			name: "タイトルがnilの場合、エラー返却",
			args: domain.Todo{
				UserID:     1,
				Content:    "test content",
				ImagePath:  "testImg",
				IsFinished: false,
			},
			prepareMockFn: func(m *mock_usecase.MockTodoRepository, args domain.Todo) {
				m.EXPECT().Store(args).Return(nil).AnyTimes()
			},
			wantErr: errors.New("タイトルは必須です。"),
		},
		{
			name: "タイトルが50文字以上の場合、エラー返却",
			args: domain.Todo{
				UserID:     1,
				Title:      strings.Repeat("t", 51),
				Content:    "test Content",
				ImagePath:  "testImg",
				IsFinished: false,
			},
			prepareMockFn: func(m *mock_usecase.MockTodoRepository, args domain.Todo) {
				m.EXPECT().Store(args).Return(nil).AnyTimes()
			},
			wantErr: errors.New("タイトルは50文字未満の入力になります。"),
		},
		{
			name: "メモが2001文字以上の場合、エラー返却",
			args: domain.Todo{
				UserID:     1,
				Title:      "test title",
				Content:    strings.Repeat("c", 2001),
				ImagePath:  "testImg",
				IsFinished: false,
			},
			prepareMockFn: func(m *mock_usecase.MockTodoRepository, args domain.Todo) {
				m.EXPECT().Store(args).Return(nil).AnyTimes()
			},
			wantErr: errors.New("メモは2000文字を超えて入力はできません。"),
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepareMockFn(TodoRepository, tt.args)
			mess, err := inter.Add(tt.args)
			if err == nil {
				if mess.Message != tt.wantMess.Message {
					t.Error("actual:", mess, "want:", tt.wantMess)
				}
			} else if err.Error() != tt.wantErr.Error() {
				t.Error("actual:", err, "want:", tt.wantErr)
			}
		})
	}
}

func TestTodos(t *testing.T) {
	inter, TodoRepository := setMock(t)

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
			name:   "userIdと現在ページ情報が正しい場合,Todo一覧を取得",
			userId: 1,
			page:   1,
			prepareMockFn: func(m *mock_usecase.MockTodoRepository, userId int, page int) {
				m.EXPECT().FindByUserId(userId, page).Return([]domain.Todo{{UserID: userId}}, 1.0, nil)
			},
			wantTodos: []domain.Todo{{UserID: 1}},
		},
		{
			name: "userIdがnilの場合,エラーを取得",
			page: 1,
			prepareMockFn: func(m *mock_usecase.MockTodoRepository, userId int, page int) {
				m.EXPECT().FindByUserId(0, page).Return([]domain.Todo{{UserID: 0}}, 0.0, nil).AnyTimes()
			},
			wantErr: errors.New("データ取得に失敗しました"),
		},
		{
			name:   "pageがnilの場合,エラーを取得",
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
				t.Error("actual:", todos, "want:", tt.wantTodos)
			}
		})
	}
}

func TestTodoByIdAndUserId(t *testing.T) {
	inter, TodoRepository := setMock(t)

	cases := []struct {
		name          string
		id            int
		userId        int
		prepareMockFn func(m *mock_usecase.MockTodoRepository, id int, userId int)
		wantTodo      *domain.Todo
		wantErr       error
	}{
		{
			name:   "idとuserIdが正しい場合、該当するtodoデータを取得",
			id:     1,
			userId: 1,
			prepareMockFn: func(m *mock_usecase.MockTodoRepository, id int, userId int) {
				m.EXPECT().FindByIdAndUserId(id, userId).Return(&domain.Todo{ID: id, UserID: userId}, nil)
			},
			wantTodo: &domain.Todo{ID: 1, UserID: 1},
		},
		{
			name:   "idがnilの場合、todoデータの取得の失敗",
			userId: 1,
			prepareMockFn: func(m *mock_usecase.MockTodoRepository, id int, userId int) {
				m.EXPECT().FindByIdAndUserId(0, userId).Return(&domain.Todo{ID: 0, UserID: userId}, nil).AnyTimes()
			},
			wantErr: errors.New("データ取得に失敗しました"),
		},
		{
			name: "userIdがnilの場合、todoデータの取得の失敗",
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
				t.Error("actual:", todo, "want:", tt.wantTodo)
			}
		})
	}
}

func TestUpdateTodo(t *testing.T) {
	inter, TodoRepository := setMock(t)

	cases := []struct {
		name          string
		args          domain.Todo
		prepareMockFn func(m *mock_usecase.MockTodoRepository, args domain.Todo)
		wantMess      usecase.TodoMessage
		wantErr       error
	}{
		{
			name: "必須項目が入力された場合、データ更新成功",
			args: domain.Todo{
				UserID:     1,
				Title:      "test update title",
				Content:    "test update content",
				ImagePath:  "testUpdateImage",
				IsFinished: false,
			},
			prepareMockFn: func(m *mock_usecase.MockTodoRepository, args domain.Todo) {
				m.EXPECT().Overwrite(args).Return(nil)
			},
			wantMess: usecase.TodoMessage{
				Message: "更新しました",
			},
		},
		{
			name: "ユーザーIDがnilの場合、エラー返却",
			args: domain.Todo{
				Title:      "test title",
				Content:    "test content",
				ImagePath:  "testImg",
				IsFinished: false,
			},
			prepareMockFn: func(m *mock_usecase.MockTodoRepository, args domain.Todo) {
				m.EXPECT().Overwrite(args).Return(nil).AnyTimes()
			},
			wantErr: errors.New("ユーザーIDは必須です。"),
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepareMockFn(TodoRepository, tt.args)
			mess, err := inter.UpdateTodo(tt.args)
			if err == nil {
				if mess.Message != tt.wantMess.Message {
					t.Error("actual:", mess, "want:", tt.wantMess)
				}
			} else if err.Error() != tt.wantErr.Error() {
				t.Error("actual:", err, "want:", tt.wantErr)
			}
		})
	}
}

func TestIsFinishedTodo(t *testing.T) {
	inter, TodoRepository := setMock(t)

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
			name:   "isFinishedがtrueの場合、メッセージ=[完了しました]",
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
			name:   "isFinishedがfalseの場合、メッセージ=[未完了の項目が追加されました]",
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
			name: "userIdがnilの場合、isFinishedの更新の失敗",
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
					t.Error("actual:", mess, "want:", tt.wantMess)
					return
				}
				return
			} else if err.Error() != tt.wantErr.Error() {
				t.Error("actual:", err, "want:", tt.wantErr)
			}
		})
	}
}

func TestDeleteTodo(t *testing.T) {
	inter, TodoRepository := setMock(t)

	cases := []struct {
		name          string
		id            int
		userId        int
		prepareMockFn func(m *mock_usecase.MockTodoRepository, id int, userId int)
		wantMess      usecase.TodoMessage
		wantErr       error
	}{
		{
			name:   "idとuserIdが正しい場合、Todoを削除成功",
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
			name:   "idがnilの場合、Todoを削除失敗",
			userId: 1,
			prepareMockFn: func(m *mock_usecase.MockTodoRepository, id int, userId int) {
				m.EXPECT().Erasure(0, userId).Return(nil).AnyTimes()
			},
			wantErr: errors.New("データ取得に失敗しました"),
		},
		{
			name: "idがnilの場合、Todoを削除失敗",
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
					t.Error("actual:", mess, "want:", tt.wantMess)
				}
			} else if err.Error() != tt.wantErr.Error() {
				t.Error("actual:", err, "want:", tt.wantErr)
			}
		})
	}
}

func TestDeleteTodoIndex(t *testing.T) {
	inter, TodoRepository := setMock(t)

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
			name:   "id、userIdとpageが正しい場合、todos,sumPageとmessageを取得",
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
			name:   "idがnilの場合、エラーメッセージeを取得",
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
			name: "userIdがnilの場合、エラーメッセージeを取得",
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
					t.Error("actual:", todos, "want:", tt.wantTodos)
					return
				}
				if sumPage != tt.wantSumPage {
					t.Error("actual:", sumPage, "want:", tt.wantSumPage)
					return
				}
				if mess.Message != tt.wantMess.Message {
					t.Error("actual:", mess, "want:", tt.wantMess)
					return
				}
			} else if err.Error() != tt.wantErr.Error() {
				t.Error("actual:", err, "want:", tt.wantErr)
			}
		})
	}
}

// --- 各種Mockをインスタンス ---
func setMock(t *testing.T) (inter *usecase.TodoInteractor, TodoRepository *mock_usecase.MockTodoRepository) {
	// mockを作成
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	TodoRepository = mock_usecase.NewMockTodoRepository(ctrl)
	inter = &usecase.TodoInteractor{}
	inter.TodoRepository = TodoRepository
	return
}
