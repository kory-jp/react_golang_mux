package usecase_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/kory-jp/react_golang_mux/api/domain"

	usecase "github.com/kory-jp/react_golang_mux/api/usecase/users"

	"github.com/golang/mock/gomock"
	mock_usecase "github.com/kory-jp/react_golang_mux/api/usecase/users/mock"
)

func TestAdd(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	UserRepository := mock_usecase.NewMockUserRepository(ctrl)
	inter := &usecase.UserInteractor{}
	inter.UserRepository = UserRepository

	cases := []struct {
		name               string
		args               domain.User
		id                 int
		prepareStoreMockFn func(m *mock_usecase.MockUserRepository, args domain.User)
		prepareFindMockFn  func(m *mock_usecase.MockUserRepository, id int)
		wantErr            error
	}{
		{
			name: "addUser = success",
			args: domain.User{
				Name:     "testName",
				Email:    "test@exm.com",
				Password: "testPassword",
			},
			id: 1,
			prepareStoreMockFn: func(m *mock_usecase.MockUserRepository, args domain.User) {
				m.EXPECT().Store(gomock.Any()).Return(1, nil)
			},
			prepareFindMockFn: func(m *mock_usecase.MockUserRepository, id int) {
				m.EXPECT().FindById(id).Return(nil, nil)
			},
			wantErr: nil,
		},
		{
			name: "when Name = nil, addUser = fail",
			args: domain.User{
				Name:     "",
				Email:    "test@exm.com",
				Password: "testPassword",
			},
			id: 1,
			prepareStoreMockFn: func(m *mock_usecase.MockUserRepository, args domain.User) {
				m.EXPECT().Store(gomock.Any()).Return(1, nil).AnyTimes()
			},
			prepareFindMockFn: func(m *mock_usecase.MockUserRepository, id int) {
				m.EXPECT().FindById(id).Return(nil, nil).AnyTimes()
			},
			wantErr: errors.New("名前は必須です。"),
		},
		{
			name: "when Name.length < 2, addUser = fail",
			args: domain.User{
				Name:     "a",
				Email:    "test@exm.com",
				Password: "testPassword",
			},
			id: 1,
			prepareStoreMockFn: func(m *mock_usecase.MockUserRepository, args domain.User) {
				m.EXPECT().Store(gomock.Any()).Return(1, nil).AnyTimes()
			},
			prepareFindMockFn: func(m *mock_usecase.MockUserRepository, id int) {
				m.EXPECT().FindById(id).Return(nil, nil).AnyTimes()
			},
			wantErr: errors.New("名前は2文字より入力が必須です。"),
		},
		{
			name: "when Name.length > 19, addUser = fail",
			args: domain.User{
				Name:     strings.Repeat("t", 20),
				Email:    "test@exm.com",
				Password: "testPassword",
			},
			id: 1,
			prepareStoreMockFn: func(m *mock_usecase.MockUserRepository, args domain.User) {
				m.EXPECT().Store(gomock.Any()).Return(1, nil).AnyTimes()
			},
			prepareFindMockFn: func(m *mock_usecase.MockUserRepository, id int) {
				m.EXPECT().FindById(id).Return(nil, nil).AnyTimes()
			},
			wantErr: errors.New("名前は20文字未満の入力になります。"),
		},
		{
			name: "when Email = nil, addUser = fail",
			args: domain.User{
				Name:     "testName",
				Email:    "",
				Password: "testPassword",
			},
			id: 1,
			prepareStoreMockFn: func(m *mock_usecase.MockUserRepository, args domain.User) {
				m.EXPECT().Store(gomock.Any()).Return(1, nil).AnyTimes()
			},
			prepareFindMockFn: func(m *mock_usecase.MockUserRepository, id int) {
				m.EXPECT().FindById(id).Return(nil, nil).AnyTimes()
			},
			wantErr: errors.New("メールアドレスは必須です。"),
		},
		{
			// Emailのフォーマットに誤りがある場合、データ保存失敗
			name: "when Email is not fomat, addUser = fail",
			args: domain.User{
				Name:     "testName",
				Email:    "testmail",
				Password: "testPassword",
			},
			id: 1,
			prepareStoreMockFn: func(m *mock_usecase.MockUserRepository, args domain.User) {
				m.EXPECT().Store(gomock.Any()).Return(1, nil).AnyTimes()
			},
			prepareFindMockFn: func(m *mock_usecase.MockUserRepository, id int) {
				m.EXPECT().FindById(id).Return(nil, nil).AnyTimes()
			},
			wantErr: errors.New("メールアドレスのフォーマットに誤りがあります"),
		},
		{
			name: "when Email > 29, addUser = fail",
			args: domain.User{
				Name:     "testName",
				Email:    "12345abcde12345abcde12345abcde@exm.com",
				Password: "testPassword",
			},
			id: 1,
			prepareStoreMockFn: func(m *mock_usecase.MockUserRepository, args domain.User) {
				m.EXPECT().Store(gomock.Any()).Return(1, nil).AnyTimes()
			},
			prepareFindMockFn: func(m *mock_usecase.MockUserRepository, id int) {
				m.EXPECT().FindById(id).Return(nil, nil).AnyTimes()
			},
			wantErr: errors.New("メールアドレスは30文字未満の入力になります。"),
		},
		{
			name: "when Password = nil, addUser = fail",
			args: domain.User{
				Name:     "testName",
				Email:    "test@exm.com",
				Password: "",
			},
			id: 1,
			prepareStoreMockFn: func(m *mock_usecase.MockUserRepository, args domain.User) {
				m.EXPECT().Store(gomock.Any()).Return(1, nil).AnyTimes()
			},
			prepareFindMockFn: func(m *mock_usecase.MockUserRepository, id int) {
				m.EXPECT().FindById(id).Return(nil, nil).AnyTimes()
			},
			wantErr: errors.New("パスワードは必須です。"),
		},
		{
			name: "when Password < 5, addUser = fail",
			args: domain.User{
				Name:     "testName",
				Email:    "test@exm.com",
				Password: strings.Repeat("t", 4),
			},
			id: 1,
			prepareStoreMockFn: func(m *mock_usecase.MockUserRepository, args domain.User) {
				m.EXPECT().Store(gomock.Any()).Return(1, nil).AnyTimes()
			},
			prepareFindMockFn: func(m *mock_usecase.MockUserRepository, id int) {
				m.EXPECT().FindById(id).Return(nil, nil).AnyTimes()
			},
			wantErr: errors.New("パスワードは5文字より入力が必須です。"),
		},
		{
			name: "when Password > 19. addUser = fail",
			args: domain.User{
				Name:     "testName",
				Email:    "test@exm.com",
				Password: strings.Repeat("t", 20),
			},
			id: 1,
			prepareStoreMockFn: func(m *mock_usecase.MockUserRepository, args domain.User) {
				m.EXPECT().Store(gomock.Any()).Return(1, nil).AnyTimes()
			},
			prepareFindMockFn: func(m *mock_usecase.MockUserRepository, id int) {
				m.EXPECT().FindById(id).Return(nil, nil).AnyTimes()
			},
			wantErr: errors.New("パスワードは20文字未満の入力になります。"),
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepareStoreMockFn(UserRepository, tt.args)
			tt.prepareFindMockFn(UserRepository, tt.id)
			_, err := inter.Add(tt.args)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
