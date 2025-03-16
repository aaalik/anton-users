package auth

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/aaalik/anton-users/internal/model"
	"github.com/aaalik/anton-users/internal/service"
	"github.com/aaalik/anton-users/pkg/testfiles"
	"github.com/dgrijalva/jwt-go"
	gomock "github.com/golang/mock/gomock"
)

func TestAuthUsecase_Login(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRP := NewMockiUserRepo(ctrl)
	mockJWTConfUtils := NewMockiJwtConfUtils(ctrl)
	mockHasherUtils := NewMockiHasherUtils(ctrl)
	mockJWTUtils := NewMockiJwtUtils(ctrl)

	username := "username"
	password := "password"
	expire := 1000
	secret := "secret"
	secretRefresh := "secret_refresh"
	// secretFailed := ""

	user := &model.User{
		Id:        "1",
		Username:  username,
		Password:  password,
		Name:      "name",
		Dob:       "2003-01-02",
		Gender:    model.USER_GENDER_MALE,
		CreatedAt: 1,
		UpdatedAt: 1,
		DeletedAt: 0,
	}

	// generate access token
	expires := time.Now().Add((time.Second * time.Duration(expire))).Unix()
	claims := jwt.MapClaims{
		"authorized": true,
		"username":   user.Username,
		"user_id":    user.Id,
		"exp":        expires,
	}

	// generate refresh token
	expiresRefresh := time.Now().Add((7 * time.Second * time.Duration(expire))).Unix()
	refreshClaims := jwt.MapClaims{
		"authorized": true,
		"username":   user.Username,
		"user_id":    user.Id,
		"exp":        expiresRefresh,
	}

	resp := &service.ResponseLogin{
		Token:                 "token",
		ExpiresAt:             time.Now().Add((time.Second * time.Duration(expire))).Unix(),
		RefreshToken:          "refresh_token",
		RefreshTokenExpiresAt: time.Now().Add((7 * time.Second * time.Duration(expire))).Unix(),
	}

	tests := []struct {
		name     string
		username string
		password string
		mock     func()
		want     *service.ResponseLogin
		wantErr  bool
	}{
		{
			name:     "success",
			username: username,
			password: password,
			mock: func() {
				mockUserRP.EXPECT().FetchUserLogin(gomock.Any(), username).Return(user, nil)
				mockHasherUtils.EXPECT().CheckPasswordHash(password, user.Password).Return(true)
				mockJWTConfUtils.EXPECT().GetExpire().Return(expire)
				mockJWTConfUtils.EXPECT().GetSecret().Return(secret)
				mockJWTUtils.EXPECT().GenerateToken(claims, secret).Return(resp.Token, nil)
				mockJWTConfUtils.EXPECT().GetSecretRefresh().Return(secretRefresh)
				mockJWTUtils.EXPECT().GenerateToken(refreshClaims, secretRefresh).Return(resp.RefreshToken, nil)
			},
			want:    resp,
			wantErr: false,
		},
		{
			name:     "failed - fetch user login",
			username: username,
			password: password,
			mock: func() {
				mockUserRP.EXPECT().FetchUserLogin(gomock.Any(), username).Return(nil, errors.New(gomock.Any().String()))
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:     "failed - check password hash",
			username: username,
			password: password,
			mock: func() {
				mockUserRP.EXPECT().FetchUserLogin(gomock.Any(), username).Return(user, nil)
				mockHasherUtils.EXPECT().CheckPasswordHash(password, user.Password).Return(false)
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:     "failed - generate access token",
			username: username,
			password: password,
			mock: func() {
				mockUserRP.EXPECT().FetchUserLogin(gomock.Any(), username).Return(user, nil)
				mockHasherUtils.EXPECT().CheckPasswordHash(password, user.Password).Return(true)
				mockJWTConfUtils.EXPECT().GetExpire().Return(expire)
				mockJWTConfUtils.EXPECT().GetSecret().Return(secret)
				mockJWTUtils.EXPECT().GenerateToken(claims, secret).Return("", errors.New(gomock.Any().String()))
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:     "failed - generate refresh token",
			username: username,
			password: password,
			mock: func() {
				mockUserRP.EXPECT().FetchUserLogin(gomock.Any(), username).Return(user, nil)
				mockHasherUtils.EXPECT().CheckPasswordHash(password, user.Password).Return(true)
				mockJWTConfUtils.EXPECT().GetExpire().Return(expire)
				mockJWTConfUtils.EXPECT().GetSecret().Return(secret)
				mockJWTUtils.EXPECT().GenerateToken(claims, secret).Return(resp.Token, nil)
				mockJWTConfUtils.EXPECT().GetSecretRefresh().Return(secretRefresh)
				mockJWTUtils.EXPECT().GenerateToken(refreshClaims, secretRefresh).Return("", errors.New(gomock.Any().String()))
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			au := &AuthUsecase{
				ur:  mockUserRP,
				jcu: mockJWTConfUtils,
				hu:  mockHasherUtils,
				jwu: mockJWTUtils,
			}

			if tt.mock != nil {
				tt.mock()
			}

			got, err := au.Login(context.Background(), tt.username, tt.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("AuthUsecase.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuthUsecase.Login() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthUsecase_RefreshToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockJWTConfUtils := NewMockiJwtConfUtils(ctrl)
	mockJWTUtils := NewMockiJwtUtils(ctrl)

	tokenString := "token"
	refreshTokenString := "refresh_token"
	secretRefresh := "secret_refresh"
	secret := "secret"
	expire := 3600
	expires := time.Now().Add((time.Second * time.Duration(expire))).Unix()

	refreshToken := jwt.New(jwt.SigningMethodHS256)
	refreshToken.Valid = true
	refreshToken.Claims = jwt.MapClaims{
		"authorized": true,
		"username":   "username",
		"user_id":    "user_id",
		"exp":        expires,
	}

	refreshTokenInvalid := jwt.New(jwt.SigningMethodHS256)
	refreshTokenInvalid.Valid = false
	refreshTokenInvalid.Claims = jwt.MapClaims{
		"authorized": true,
		"username":   "username",
		"user_id":    "user_id",
		"exp":        expires,
	}

	accessClaims := jwt.MapClaims{
		"authorized": true,
		"username":   refreshToken.Claims.(jwt.MapClaims)["username"],
		"user_id":    refreshToken.Claims.(jwt.MapClaims)["user_id"],
		"exp":        expires,
	}

	tests := []struct {
		name               string
		refreshTokenString string
		mock               func()
		want               string
		want1              int64
		wantErr            bool
	}{
		{
			name:               "success",
			refreshTokenString: refreshTokenString,
			mock: func() {
				mockJWTConfUtils.EXPECT().GetSecretRefresh().Return(secretRefresh)
				mockJWTUtils.EXPECT().ParseToken(refreshTokenString, secretRefresh).Return(refreshToken, nil)
				mockJWTConfUtils.EXPECT().GetExpire().Return(expire)
				mockJWTConfUtils.EXPECT().GetSecret().Return(secret)
				mockJWTUtils.EXPECT().GenerateToken(accessClaims, secret).Return(tokenString, nil)
			},
			want:    tokenString,
			want1:   expires,
			wantErr: false,
		},
		{
			name:               "failed - parse token",
			refreshTokenString: refreshTokenString,
			mock: func() {
				mockJWTConfUtils.EXPECT().GetSecretRefresh().Return(secretRefresh)
				mockJWTUtils.EXPECT().ParseToken(refreshTokenString, secretRefresh).Return(nil, errors.New(gomock.Any().String()))
			},
			want:    "",
			want1:   0,
			wantErr: true,
		},
		{
			name:               "failed - invalid token",
			refreshTokenString: refreshTokenString,
			mock: func() {
				mockJWTConfUtils.EXPECT().GetSecretRefresh().Return(secretRefresh)
				mockJWTUtils.EXPECT().ParseToken(refreshTokenString, secretRefresh).Return(refreshTokenInvalid, nil)
			},
			want:    "",
			want1:   0,
			wantErr: true,
		},
		{
			name:               "failed - generate token",
			refreshTokenString: refreshTokenString,
			mock: func() {
				mockJWTConfUtils.EXPECT().GetSecretRefresh().Return(secretRefresh)
				mockJWTUtils.EXPECT().ParseToken(refreshTokenString, secretRefresh).Return(refreshToken, nil)
				mockJWTConfUtils.EXPECT().GetExpire().Return(expire)
				mockJWTConfUtils.EXPECT().GetSecret().Return(secret)
				mockJWTUtils.EXPECT().GenerateToken(accessClaims, secret).Return("", errors.New(gomock.Any().String()))
			},
			want:    "",
			want1:   0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			au := &AuthUsecase{
				jcu: mockJWTConfUtils,
				jwu: mockJWTUtils,
			}

			if tt.mock != nil {
				tt.mock()
			}

			got, got1, err := au.RefreshToken(context.Background(), tt.refreshTokenString)
			if (err != nil) != tt.wantErr {
				t.Errorf("AuthUsecase.RefreshToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AuthUsecase.RefreshToken() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("AuthUsecase.RefreshToken() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestAuthUsecase_Register(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRP := NewMockiUserRepo(ctrl)
	mockDatabaseUtils := NewMockiDatabaseUtils(ctrl)
	mockHasherUtils := NewMockiHasherUtils(ctrl)
	mockRandomUtils := NewMockiRandomUtils(ctrl)

	reqRegister := &service.RequestRegister{
		Username: "username",
		Password: "password",
		Name:     "name",
	}

	id := "1"
	hashedPassword := "hashed_password"

	user := model.User{
		Id:       id,
		Username: reqRegister.Username,
		Password: hashedPassword,
		Name:     reqRegister.Name,
	}

	respRegister := &service.ResponseRegister{
		Username: user.Username,
		Name:     user.Name,
	}

	tests := []struct {
		name    string
		req     *service.RequestRegister
		mock    func()
		want    *service.ResponseRegister
		wantErr bool
	}{
		{
			name: "success",
			req:  reqRegister,
			mock: func() {
				mockHasherUtils.EXPECT().HashPassword(reqRegister.Password).Return(hashedPassword, nil)
				mockRandomUtils.EXPECT().UniqueID().Return(id)
				mockDatabaseUtils.EXPECT().ExecuteTx(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(testfiles.DoExecuteTx)
				mockUserRP.EXPECT().CreateUser(gomock.Any(), gomock.Any(), &user).Return(nil)
			},
			want:    respRegister,
			wantErr: false,
		},
		{
			name: "failed - hash password",
			req:  reqRegister,
			mock: func() {
				mockHasherUtils.EXPECT().HashPassword(reqRegister.Password).Return("", errors.New(gomock.Any().String()))
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "failed - execute tx",
			req:  reqRegister,
			mock: func() {
				mockHasherUtils.EXPECT().HashPassword(reqRegister.Password).Return(hashedPassword, nil)
				mockRandomUtils.EXPECT().UniqueID().Return(id)
				mockDatabaseUtils.EXPECT().ExecuteTx(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New(gomock.Any().String()))
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "failed - create user",
			req:  reqRegister,
			mock: func() {
				mockHasherUtils.EXPECT().HashPassword(reqRegister.Password).Return(hashedPassword, nil)
				mockRandomUtils.EXPECT().UniqueID().Return(id)
				mockDatabaseUtils.EXPECT().ExecuteTx(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(testfiles.DoExecuteTx)
				mockUserRP.EXPECT().CreateUser(gomock.Any(), gomock.Any(), &user).Return(errors.New(gomock.Any().String()))
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			au := &AuthUsecase{
				ur:  mockUserRP,
				dbu: mockDatabaseUtils,
				hu:  mockHasherUtils,
				ru:  mockRandomUtils,
			}

			if tt.mock != nil {
				tt.mock()
			}

			got, err := au.Register(context.Background(), tt.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("AuthUsecase.Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuthUsecase.Register() = %v, want %v", got, tt.want)
			}
		})
	}
}
