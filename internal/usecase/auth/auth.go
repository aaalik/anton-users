package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/RoseRocket/xerrs"
	cons "github.com/aaalik/anton-users/internal/constant"
	"github.com/aaalik/anton-users/internal/model"
	"github.com/aaalik/anton-users/internal/service"
	"github.com/aaalik/anton-users/pkg/hasher"
	"github.com/aaalik/anton-users/pkg/utils"
	"github.com/dgrijalva/jwt-go"
)

func (au *AuthUsecase) Login(ctx context.Context, username, password string) (*service.ResponseLogin, error) {
	user, err := au.ur.FetchUserLogin(ctx, username)
	if err != nil {
		return nil, xerrs.Mask(err, cons.ErrorInvalidLogin)
	}

	isValid := hasher.CheckPasswordHash(password, user.Password)

	if !isValid {
		return nil, cons.ErrorInvalidLogin
	}

	token := jwt.New(jwt.SigningMethodHS256)
	expires := time.Now().Add((time.Second * time.Duration(au.jcu.GetExpire()))).Unix()

	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["username"] = user.Username
	claims["user_id"] = user.Id
	claims["exp"] = expires

	tokenString, err := token.SignedString([]byte(au.jcu.GetSecret()))
	if err != nil {
		return nil, xerrs.Mask(err, cons.ErrorInvalidLogin)
	}

	refreshToken := jwt.New(jwt.SigningMethodHS256)
	expiresRefresh := time.Now().Add((7 * time.Second * time.Duration(au.jcu.GetExpire()))).Unix()

	refreshClaims := refreshToken.Claims.(jwt.MapClaims)
	refreshClaims["username"] = user.Username
	refreshClaims["user_id"] = user.Id
	refreshClaims["exp"] = expiresRefresh

	refreshTokenString, err := refreshToken.SignedString([]byte(au.jcu.GetSecret()))
	if err != nil {
		return nil, xerrs.Mask(err, cons.ErrorInvalidLogin)
	}

	resp := &service.ResponseLogin{
		Token:                 tokenString,
		ExpiresAt:             expires,
		RefreshToken:          refreshTokenString,
		RefreshTokenExpiresAt: expiresRefresh,
	}

	return resp, nil
}

func (au *AuthUsecase) RefreshToken(ctx context.Context, refreshToken string) (string, int64, error) {
	rfrshToken, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(au.jcu.GetSecretRefresh()), nil
	})
	if err != nil {
		return "", 0, xerrs.Mask(err, cons.ErrorInvalidLogin)
	}

	if !rfrshToken.Valid {
		return "", 0, cons.ErrorInvalidLogin
	}

	token := jwt.New(jwt.SigningMethodHS256)
	expires := time.Now().Add((time.Second * time.Duration(au.jcu.GetExpire()))).Unix()

	accessClaims := token.Claims.(jwt.MapClaims)
	accessClaims["authorized"] = true
	accessClaims["username"] = rfrshToken.Claims.(jwt.MapClaims)["username"]
	accessClaims["user_id"] = rfrshToken.Claims.(jwt.MapClaims)["user_id"]
	accessClaims["exp"] = expires

	tokenString, err := token.SignedString([]byte(au.jcu.GetSecret()))
	if err != nil {
		return "", 0, xerrs.Mask(err, cons.ErrorInvalidLogin)
	}

	return tokenString, 0, nil
}

func (au *AuthUsecase) Register(ctx context.Context, request *service.RequestRegister) (*service.ResponseRegister, error) {
	pwd, err := hasher.HashPassword(request.Password)
	if err != nil {
		return nil, err
	}

	user := model.User{
		Id:       utils.UniqueID(),
		Username: request.Username,
		Password: pwd,
		Name:     request.Name,
	}

	tx, err := au.ur.CreateTx(ctx)
	if err != nil {
		return nil, err
	}
	defer au.ur.RollbackTx(ctx, tx)

	err = au.ur.CreateUser(ctx, tx, &user)
	if err != nil {
		return nil, err
	}

	err = au.ur.CommitTx(ctx, tx)
	if err != nil {
		return nil, err
	}

	return &service.ResponseRegister{Username: user.Username, Name: user.Name}, nil
}
