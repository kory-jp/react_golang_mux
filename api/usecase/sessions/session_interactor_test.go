package usecase_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/kory-jp/react_golang_mux/api/domain"

	"github.com/golang/mock/gomock"
	usecase "github.com/kory-jp/react_golang_mux/api/usecase/sessions"
	mock_usecase "github.com/kory-jp/react_golang_mux/api/usecase/sessions/mock"
)

func TestLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	SessionRepository := mock_usecase.NewMockSessionRepository(ctrl)
	inter := &usecase.SessionInteractor{}
	inter.SessionRepository = SessionRepository
	var pUser *domain.User

	// --- テストケース ---
	cases := []struct {
		name                 string
		args                 domain.User
		prepareByEmailMockFn func(m *mock_usecase.MockSessionRepository, args domain.User)
		wantErr              error
	}{
		{
			// EmailとPasswordが入力されPasswordが一致した場合、ログイン成功
			name: "email = correct password = correct, login = success",
			args: domain.User{
				Email:    "test@exm.com",
				Password: "testPassword",
			},
			prepareByEmailMockFn: func(m *mock_usecase.MockSessionRepository, args domain.User) {
				pUser = &domain.User{
					Email: args.Email,
					// --- 戻り値のuser.PasswordをEncryptメソッドで暗号化した文字列に変換 ---
					// --- 引数で渡たすargs.Password(暗号化前のuser.Passwordと同一文字列)と
					// 暗号化したuser.PasswordをCompareHashAndPasswordメソッドを用いて比較 ---
					Password: args.Encrypt(args.Password),
				}
				m.EXPECT().FindByEmail(args).Return(pUser, nil)
			},
			wantErr: nil,
		},
		{
			name: "when email = nil, login = fail",
			args: domain.User{
				Email:    "",
				Password: "testPassword",
			},
			prepareByEmailMockFn: func(m *mock_usecase.MockSessionRepository, args domain.User) {
				pUser = &domain.User{
					Email:    args.Email,
					Password: args.Encrypt(args.Password),
				}
				m.EXPECT().FindByEmail(args).Return(pUser, nil).AnyTimes()
			},
			wantErr: errors.New("認証に失敗しました"),
		},
		{
			name: "when password = nil, login = fail",
			args: domain.User{
				Email:    "test@exm.com",
				Password: "",
			},
			prepareByEmailMockFn: func(m *mock_usecase.MockSessionRepository, args domain.User) {
				pUser = &domain.User{
					Email:    args.Email,
					Password: args.Encrypt(args.Password),
				}
				m.EXPECT().FindByEmail(args).Return(pUser, nil).AnyTimes()
			},
			wantErr: errors.New("認証に失敗しました"),
		},
		{
			// EmailとPasswordが入力されたがPasswordが一致しない場合、ログイン失敗
			name: "when password is not equal, login = fail",
			args: domain.User{
				Email:    "test@exm.com",
				Password: "testPassword",
			},
			prepareByEmailMockFn: func(m *mock_usecase.MockSessionRepository, args domain.User) {
				pUser = &domain.User{
					Email: args.Email,
					// --- 引数で渡されるargs.Passwordと異なる文字列のPasswordを戻り値に設定 ---
					Password: args.Encrypt("notEqualPassword"),
				}
				m.EXPECT().FindByEmail(args).Return(pUser, nil).AnyTimes()
			},
			wantErr: errors.New("認証に失敗しました"),
		},
	}

	// --- テスト実行 ---
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepareByEmailMockFn(SessionRepository, tt.args)
			_, err := inter.Login(tt.args)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestIsLoggedIn(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	SessionRepository := mock_usecase.NewMockSessionRepository(ctrl)
	inter := &usecase.SessionInteractor{}
	inter.SessionRepository = SessionRepository
	var pUser *domain.User

	cases := []struct {
		name              string
		userId            int
		prepareByIdMockFn func(m *mock_usecase.MockSessionRepository, userId int)
		wantErr           error
	}{
		{
			name:   "authentication = success",
			userId: 1,
			prepareByIdMockFn: func(m *mock_usecase.MockSessionRepository, userId int) {
				m.EXPECT().FindById(userId).Return(pUser, nil)
			},
			wantErr: nil,
		},
		{
			name:   "when userId = 0, authentication = fail",
			userId: 0,
			prepareByIdMockFn: func(m *mock_usecase.MockSessionRepository, userId int) {
				m.EXPECT().FindById(userId).Return(pUser, nil).AnyTimes()
			},
			wantErr: errors.New("認証に失敗しました"),
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepareByIdMockFn(SessionRepository, tt.userId)
			_, err := inter.IsLoggedin(tt.userId)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
