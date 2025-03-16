package auth

import (
	"context"
	"time"

	"github.com/RoseRocket/xerrs"
	cons "github.com/aaalik/anton-users/internal/constant"
	"github.com/aaalik/anton-users/internal/model"
	"github.com/aaalik/anton-users/internal/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/jmoiron/sqlx"
)

func (au *AuthUsecase) Login(ctx context.Context, username, password string) (*service.ResponseLogin, error) {
	user, err := au.ur.FetchUserLogin(ctx, username)
	if err != nil {
		return nil, xerrs.Mask(err, cons.ErrorInvalidLogin)
	}

	isValid := au.hu.CheckPasswordHash(password, user.Password)
	if !isValid {
		return nil, xerrs.Extend(cons.ErrorInvalidLogin)
	}

	expire := au.jcu.GetExpire()

	// generate access token
	expires := time.Now().Add((time.Second * time.Duration(expire))).Unix()

	claims := jwt.MapClaims{
		"authorized": true,
		"username":   user.Username,
		"user_id":    user.Id,
		"exp":        expires,
	}

	token, err := au.jwu.GenerateToken(claims)
	if err != nil {
		return nil, xerrs.Mask(err, cons.ErrorInvalidLogin)
	}

	// generate refresh token
	expiresRefresh := time.Now().Add((7 * time.Second * time.Duration(expire))).Unix()

	refreshClaims := jwt.MapClaims{
		"authorized": true,
		"username":   user.Username,
		"user_id":    user.Id,
		"exp":        expiresRefresh,
	}

	refreshToken, err := au.jwu.GenerateToken(refreshClaims)
	if err != nil {
		return nil, xerrs.Mask(err, cons.ErrorInvalidLogin)
	}

	resp := &service.ResponseLogin{
		Token:                 token,
		ExpiresAt:             expires,
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: expiresRefresh,
	}

	return resp, nil
}

func (au *AuthUsecase) RefreshToken(ctx context.Context, refreshToken string) (string, int64, error) {
	rfrshToken, err := au.jwu.ParseToken(refreshToken)
	if err != nil {
		return "", 0, xerrs.Mask(err, cons.ErrorInvalidLogin)
	}

	if !rfrshToken.Valid {
		return "", 0, xerrs.Extend(cons.ErrorInvalidLogin)
	}

	expires := time.Now().Add((time.Second * time.Duration(au.jcu.GetExpire()))).Unix()

	accessClaims := jwt.MapClaims{
		"authorized": true,
		"username":   rfrshToken.Claims.(jwt.MapClaims)["username"],
		"user_id":    rfrshToken.Claims.(jwt.MapClaims)["user_id"],
		"exp":        expires,
	}

	token, err := au.jwu.GenerateToken(accessClaims)
	if err != nil {
		return "", 0, xerrs.Mask(err, cons.ErrorInvalidLogin)
	}

	return token, expires, nil
}

func (au *AuthUsecase) Register(ctx context.Context, request *service.RequestRegister) (*service.ResponseRegister, error) {
	pwd, err := au.hu.HashPassword(request.Password)
	if err != nil {
		return nil, err
	}

	user := model.User{
		Id:       au.ru.UniqueID(),
		Username: request.Username,
		Password: pwd,
		Name:     request.Name,
	}

	err = au.dbu.ExecuteTx(ctx, nil, func(ctx context.Context, tx *sqlx.Tx) error {
		return au.ur.CreateUser(ctx, tx, &user)
	})
	if err != nil {
		return nil, err
	}

	return &service.ResponseRegister{Username: user.Username, Name: user.Name}, nil
}
