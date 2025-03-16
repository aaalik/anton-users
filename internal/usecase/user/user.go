package user

import (
	"context"

	"github.com/aaalik/anton-users/internal/model"
	"github.com/aaalik/anton-users/internal/service"
	"github.com/jmoiron/sqlx"
)

func (uu *UserUsecase) CreateUser(ctx context.Context, request *service.RequestCreateUser) (*model.User, error) {
	pwd, err := uu.hu.HashPassword(request.Password)
	if err != nil {
		return nil, err
	}

	user := model.User{
		Id:       uu.ru.UniqueID(),
		Username: request.Username,
		Password: pwd,
		Name:     request.Name,
		Dob:      request.Dob,
		Gender:   request.Gender,
	}

	err = uu.dbu.ExecuteTx(ctx, nil, func(ctx context.Context, tx *sqlx.Tx) error {
		return uu.ur.CreateUser(ctx, tx, &user)
	})
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (uu *UserUsecase) UpdateUser(ctx context.Context, request *service.RequestUpdateUser) (*model.User, error) {
	user, err := uu.DetailUser(ctx, request.Id)
	if err != nil {
		return nil, err
	}

	reqUser := model.User{
		Id:     user.Id,
		Name:   request.Name,
		Dob:    request.Dob,
		Gender: request.Gender,
	}

	err = uu.dbu.ExecuteTx(ctx, nil, func(ctx context.Context, tx *sqlx.Tx) error {
		return uu.ur.UpdateUser(ctx, tx, &reqUser)
	})
	if err != nil {
		return nil, err
	}

	return &reqUser, nil
}

func (uu *UserUsecase) DeleteUser(ctx context.Context, id string) error {
	return uu.dbu.ExecuteTx(ctx, nil, func(ctx context.Context, tx *sqlx.Tx) error {
		return uu.ur.DeleteUser(ctx, tx, id)
	})
}

func (uu *UserUsecase) DetailUser(ctx context.Context, id string) (*model.User, error) {
	return uu.ur.DetailUser(ctx, id)
}

func (uu *UserUsecase) ListUser(ctx context.Context, request *service.RequestListUser) ([]*model.User, int32, error) {
	users, err := uu.ur.ListUser(ctx, request)
	if err != nil {
		return nil, 0, err
	}

	count, err := uu.ur.CountUsers(ctx, request)
	if err != nil {
		return nil, 0, err
	}

	return users, count, nil
}
