package usecase_test

import (
	"errors"
	"reflect"
	"strings"
	"testing"

	"github.com/kory-jp/react_golang_mux/api/domain"
	"github.com/stretchr/testify/assert"

	"github.com/golang/mock/gomock"
	usecase "github.com/kory-jp/react_golang_mux/api/usecase/task_cards"
	mock_usecase "github.com/kory-jp/react_golang_mux/api/usecase/task_cards/mock"
)

func TestAdd(t *testing.T) {
	inter, TaskCardRepository := setMock(t)

	cases := []struct {
		name          string
		args          domain.TaskCard
		prepareMockFn func(m *mock_usecase.MockTaskCardRepository, args domain.TaskCard)
		wantMess      usecase.TaskCardMessage
		wantErr       error
	}{
		{
			name: "create = success",
			args: domain.TaskCard{
				UserID:  1,
				TodoID:  1,
				Title:   "test title",
				Purpose: "test purpose",
				Content: "test content",
				Memo:    "test memo",
			},
			prepareMockFn: func(m *mock_usecase.MockTaskCardRepository, arges domain.TaskCard) {
				m.EXPECT().Store(arges).Return(nil)
			},
			wantMess: usecase.TaskCardMessage{
				Message: "保存しました",
			},
		},
		{
			name: "when userId = 0, create = fail",
			args: domain.TaskCard{
				TodoID:  1,
				Title:   "test title",
				Purpose: "test purpose",
				Content: "test content",
				Memo:    "test memo",
			},
			prepareMockFn: func(m *mock_usecase.MockTaskCardRepository, arges domain.TaskCard) {
				m.EXPECT().Store(arges).Return(nil)
			},
			wantErr: errors.New("ユーザーIDは必須です。"),
		},
		{
			name: "when title = nil, create = fail",
			args: domain.TaskCard{
				UserID:  1,
				TodoID:  1,
				Purpose: "test purpose",
				Content: "test content",
				Memo:    "test memo",
			},
			prepareMockFn: func(m *mock_usecase.MockTaskCardRepository, arges domain.TaskCard) {
				m.EXPECT().Store(arges).Return(nil)
			},
			wantErr: errors.New("タイトルは必須です。"),
		},
		{
			name: "when title > 49, create = fail",
			args: domain.TaskCard{
				UserID:  1,
				TodoID:  1,
				Title:   strings.Repeat("c", 50),
				Purpose: "test purpose",
				Content: "test content",
				Memo:    "test memo",
			},
			prepareMockFn: func(m *mock_usecase.MockTaskCardRepository, arges domain.TaskCard) {
				m.EXPECT().Store(arges).Return(nil)
			},
			wantErr: errors.New("タイトルは50文字未満の入力になります。"),
		},
		{
			name: "when purpose > 1999, create = fail",
			args: domain.TaskCard{
				UserID:  1,
				TodoID:  1,
				Title:   "test title",
				Purpose: strings.Repeat("c", 2000),
				Content: "test content",
				Memo:    "test memo",
			},
			prepareMockFn: func(m *mock_usecase.MockTaskCardRepository, arges domain.TaskCard) {
				m.EXPECT().Store(arges).Return(nil)
			},
			wantErr: errors.New("目的は2000文字未満の入力になります。"),
		},
		{
			name: "when content > 1999, create = fail",
			args: domain.TaskCard{
				UserID:  1,
				TodoID:  1,
				Title:   "test title",
				Purpose: "test purpose",
				Content: strings.Repeat("c", 2000),
				Memo:    "test memo",
			},
			prepareMockFn: func(m *mock_usecase.MockTaskCardRepository, arges domain.TaskCard) {
				m.EXPECT().Store(arges).Return(nil)
			},
			wantErr: errors.New("作業内容は2000文字未満の入力になります。"),
		},
		{
			name: "when memo > 1999, create = fail",
			args: domain.TaskCard{
				UserID:  1,
				TodoID:  1,
				Title:   "test title",
				Purpose: "test purpose",
				Content: "test content",
				Memo:    strings.Repeat("c", 2000),
			},
			prepareMockFn: func(m *mock_usecase.MockTaskCardRepository, arges domain.TaskCard) {
				m.EXPECT().Store(arges).Return(nil)
			},
			wantErr: errors.New("メモは2000文字未満の入力になります。"),
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepareMockFn(TaskCardRepository, tt.args)
			mess, err := inter.Add(tt.args)
			if err == nil {
				assert.Equal(t, tt.wantMess, *mess)
			} else if err.Error() != tt.wantErr.Error() {
				assert.Equal(t, tt.wantErr, err)
			}
		})
	}
}

func TestTaskCards(t *testing.T) {
	inter, TaskCardRepository := setMock(t)

	cases := []struct {
		name          string
		userId        int
		todoId        int
		page          int
		prepareMockFn func(m *mock_usecase.MockTaskCardRepository, userId int, todoId int, page int)
		wantTaskCards domain.TaskCards
		wantErr       error
	}{
		{
			name:   "getTaskCards = success",
			userId: 1,
			todoId: 1,
			page:   1,
			prepareMockFn: func(m *mock_usecase.MockTaskCardRepository, userId int, todoId int, page int) {
				m.EXPECT().FindByTodoIdAndUserId(todoId, userId, page).Return([]domain.TaskCard{{UserID: 1, TodoID: 1}}, 0.0, nil)
			},
			wantTaskCards: []domain.TaskCard{{UserID: 1, TodoID: 1}},
		},
		{
			name:   "when userId = 0, getTaskCards = fail",
			userId: 0,
			todoId: 1,
			page:   1,
			prepareMockFn: func(m *mock_usecase.MockTaskCardRepository, userId int, todoId int, page int) {
				m.EXPECT().FindByTodoIdAndUserId(todoId, userId, page).Return([]domain.TaskCard{{UserID: 0, TodoID: 1}}, 0.0, nil)
			},
			wantErr: errors.New("データ取得に失敗しました"),
		},
		{
			name:   "when todoId = 0, getTaskCards = fail",
			userId: 1,
			todoId: 0,
			page:   1,
			prepareMockFn: func(m *mock_usecase.MockTaskCardRepository, userId int, todoId int, page int) {
				m.EXPECT().FindByTodoIdAndUserId(todoId, userId, page).Return([]domain.TaskCard{{UserID: 0, TodoID: 1}}, 0.0, nil)
			},
			wantErr: errors.New("データ取得に失敗しました"),
		},
		{
			name:   "when page = 0, getTaskCards = fail",
			userId: 1,
			todoId: 1,
			page:   0,
			prepareMockFn: func(m *mock_usecase.MockTaskCardRepository, userId int, todoId int, page int) {
				m.EXPECT().FindByTodoIdAndUserId(todoId, userId, page).Return([]domain.TaskCard{{UserID: 0, TodoID: 1}}, 0.0, nil)
			},
			wantErr: errors.New("データ取得に失敗しました"),
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepareMockFn(TaskCardRepository, tt.userId, tt.todoId, tt.page)
			taskCards, _, err := inter.TaskCards(tt.todoId, tt.userId, tt.page)
			if err != nil {
				if err.Error() != tt.wantErr.Error() {
					t.Error("actual:", err, "want:", tt.wantErr)
					return
				}
				return
			}
			if !reflect.DeepEqual(taskCards, tt.wantTaskCards) {
				assert.Equal(t, tt.wantTaskCards, taskCards)
			}
		})
	}
}

func TestTaskCardByIdAndUserId(t *testing.T) {
	inter, TaskCardRepository := setMock(t)

	cases := []struct {
		name          string
		userId        int
		taskCardId    int
		prepareMockFn func(m *mock_usecase.MockTaskCardRepository, userId int, taskCardId int)
		wantTaskCard  *domain.TaskCard
		wantErr       error
	}{
		{
			name:       "getTaskCard = success",
			userId:     1,
			taskCardId: 1,
			prepareMockFn: func(m *mock_usecase.MockTaskCardRepository, userId int, taskCardId int) {
				m.EXPECT().FindByIdAndUserId(taskCardId, userId).Return(&domain.TaskCard{ID: 1, UserID: 1}, nil)
			},
			wantTaskCard: &domain.TaskCard{ID: 1, UserID: 1},
		},
		{
			name:       "when userId = 0, getTaskCard = fail",
			userId:     0,
			taskCardId: 1,
			prepareMockFn: func(m *mock_usecase.MockTaskCardRepository, userId int, taskCardId int) {
				m.EXPECT().FindByIdAndUserId(taskCardId, userId).Return(&domain.TaskCard{ID: 1, UserID: 1}, nil)
			},
			wantErr: errors.New("データ取得に失敗しました"),
		},
		{
			name:       "when taskCardId = 0, getTaskCard = fail",
			userId:     1,
			taskCardId: 0,
			prepareMockFn: func(m *mock_usecase.MockTaskCardRepository, userId int, taskCardId int) {
				m.EXPECT().FindByIdAndUserId(taskCardId, userId).Return(&domain.TaskCard{ID: 1, UserID: 1}, nil)
			},
			wantErr: errors.New("データ取得に失敗しました"),
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepareMockFn(TaskCardRepository, tt.userId, tt.taskCardId)
			taskCard, err := inter.TaskCardByIdAndUserId(tt.taskCardId, tt.userId)
			if err != nil {
				if err.Error() != tt.wantErr.Error() {
					t.Error("actual:", err, "want:", tt.wantErr)
					return
				}
				return
			}
			if !reflect.DeepEqual(taskCard, tt.wantTaskCard) {
				assert.Equal(t, tt.wantTaskCard, taskCard)
			}
		})
	}
}

func TestUpdateTaskCard(t *testing.T) {
	inter, TaskCardRepository := setMock(t)

	cases := []struct {
		name          string
		args          domain.TaskCard
		prepareMockFn func(m *mock_usecase.MockTaskCardRepository, args domain.TaskCard)
		wantMess      usecase.TaskCardMessage
		wantErr       error
	}{
		{
			name: "update = success",
			args: domain.TaskCard{
				UserID:  1,
				TodoID:  1,
				Title:   "test title",
				Purpose: "test purpose",
				Content: "test content",
				Memo:    "test memo",
			},
			prepareMockFn: func(m *mock_usecase.MockTaskCardRepository, args domain.TaskCard) {
				m.EXPECT().Overwrite(args).Return(nil)
			},
			wantMess: usecase.TaskCardMessage{
				Message: "更新しました",
			},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepareMockFn(TaskCardRepository, tt.args)
			mess, err := inter.UpdateTaskCard(tt.args)
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

func TestIsFinishedTaskCard(t *testing.T) {
	inter, TaskCardRepository := setMock(t)

	cases := []struct {
		name                  string
		tcId                  int
		userId                int
		args                  domain.TaskCard
		prepareMockChangeBlFn func(m *mock_usecase.MockTaskCardRepository, tcId int, taskCard domain.TaskCard, userId int)
		prepareMockFindByIdFn func(m *mock_usecase.MockTaskCardRepository, tcId int, userId int)
		wantMess              usecase.TaskCardMessage
		wantErr               error
	}{
		{
			name:   "when isFinished = true, result = 完了しました",
			tcId:   1,
			userId: 1,
			args: domain.TaskCard{
				IsFinished: true,
			},
			prepareMockChangeBlFn: func(m *mock_usecase.MockTaskCardRepository, tcId int, args domain.TaskCard, userId int) {
				m.EXPECT().ChangeBoolean(tcId, userId, args).Return(nil)
			},
			prepareMockFindByIdFn: func(m *mock_usecase.MockTaskCardRepository, tcId int, userId int) {
				m.EXPECT().FindByIdAndUserId(tcId, userId).Return(&domain.TaskCard{IsFinished: true}, nil)
			},
			wantMess: usecase.TaskCardMessage{
				Message: "完了しました",
			},
		},
		{
			name:   "when isFinished = false, result = 未完了の項目が追加されました",
			tcId:   1,
			userId: 1,
			args: domain.TaskCard{
				IsFinished: false,
			},
			prepareMockChangeBlFn: func(m *mock_usecase.MockTaskCardRepository, tcId int, args domain.TaskCard, userId int) {
				m.EXPECT().ChangeBoolean(tcId, userId, args).Return(nil)
			},
			prepareMockFindByIdFn: func(m *mock_usecase.MockTaskCardRepository, tcId int, userId int) {
				m.EXPECT().FindByIdAndUserId(tcId, userId).Return(&domain.TaskCard{IsFinished: false}, nil)
			},
			wantMess: usecase.TaskCardMessage{
				Message: "未完了の項目が追加されました",
			},
		},
		{
			name:   "tcId = 0, changeIsFinished = fail",
			tcId:   0,
			userId: 1,
			args: domain.TaskCard{
				IsFinished: false,
			},
			prepareMockChangeBlFn: func(m *mock_usecase.MockTaskCardRepository, tcId int, args domain.TaskCard, userId int) {
				m.EXPECT().ChangeBoolean(tcId, userId, args).Return(nil)
			},
			prepareMockFindByIdFn: func(m *mock_usecase.MockTaskCardRepository, tcId int, userId int) {
				m.EXPECT().FindByIdAndUserId(tcId, userId).Return(&domain.TaskCard{IsFinished: false}, nil)
			},
			wantErr: errors.New("データ取得に失敗しました"),
		},
		{
			name:   "userId = 0, changeIsFinished = fail",
			tcId:   1,
			userId: 0,
			args: domain.TaskCard{
				IsFinished: false,
			},
			prepareMockChangeBlFn: func(m *mock_usecase.MockTaskCardRepository, tcId int, args domain.TaskCard, userId int) {
				m.EXPECT().ChangeBoolean(tcId, userId, args).Return(nil)
			},
			prepareMockFindByIdFn: func(m *mock_usecase.MockTaskCardRepository, tcId int, userId int) {
				m.EXPECT().FindByIdAndUserId(tcId, userId).Return(&domain.TaskCard{IsFinished: false}, nil)
			},
			wantErr: errors.New("データ取得に失敗しました"),
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepareMockChangeBlFn(TaskCardRepository, tt.tcId, tt.args, tt.userId)
			tt.prepareMockFindByIdFn(TaskCardRepository, tt.tcId, tt.userId)
			mess, err := inter.IsFinishedTaskCard(tt.tcId, tt.args, tt.userId)
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

func TestDeleteTaskCard(t *testing.T) {
	inter, TaskCardRepository := setMock(t)

	cases := []struct {
		name          string
		tcId          int
		userId        int
		prepareMockFn func(m *mock_usecase.MockTaskCardRepository, tcId int, userId int)
		wantMess      usecase.TaskCardMessage
		wantErr       error
	}{
		{
			name:   "delete = success",
			tcId:   1,
			userId: 1,
			prepareMockFn: func(m *mock_usecase.MockTaskCardRepository, tcId int, userId int) {
				m.EXPECT().Erasure(tcId, userId).Return(nil)
			},
			wantMess: usecase.TaskCardMessage{
				Message: "削除しました",
			},
		},
		{
			name:   "when tcId = 0, delete = fail",
			tcId:   0,
			userId: 1,
			prepareMockFn: func(m *mock_usecase.MockTaskCardRepository, tcId int, userId int) {
				m.EXPECT().Erasure(tcId, userId).Return(nil)
			},
			wantErr: errors.New("データ取得に失敗しました"),
		},
		{
			name:   "when userId = 0, delete = fail",
			tcId:   1,
			userId: 0,
			prepareMockFn: func(m *mock_usecase.MockTaskCardRepository, tcId int, userId int) {
				m.EXPECT().Erasure(tcId, userId).Return(nil)
			},
			wantErr: errors.New("データ取得に失敗しました"),
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepareMockFn(TaskCardRepository, tt.tcId, tt.userId)
			mess, err := inter.DeleteTaskCard(tt.tcId, tt.userId)
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

func setMock(t *testing.T) (inter *usecase.TaskCardInteractor, TaskCardRepository *mock_usecase.MockTaskCardRepository) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	TaskCardRepository = mock_usecase.NewMockTaskCardRepository(ctrl)
	inter = &usecase.TaskCardInteractor{}
	inter.TaskCardRepository = TaskCardRepository
	return
}
