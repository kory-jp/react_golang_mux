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
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	TodoRepository := mock_usecase.NewMockTodoRepository(ctrl)
	inter := &usecase.TodoInteractor{}
	inter.TodoRepository = TodoRepository

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
				if mess != tt.wantMess {
					t.Error("actual:", mess, "want:", tt.wantMess)
				}
			} else if err.Error() != tt.wantErr.Error() {
				t.Error("actual:", err, "want:", tt.wantErr)
			}
		})
	}
}

func TestTodos(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	TodoRepository := mock_usecase.NewMockTodoRepository(ctrl)
	inter := &usecase.TodoInteractor{}
	inter.TodoRepository = TodoRepository

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
