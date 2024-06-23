package user

import (
	"context"

	"github.com/aaalik/anton-users/internal/model"
	"github.com/aaalik/anton-users/internal/service"
	"github.com/aaalik/anton-users/pkg/hasher"
	"github.com/aaalik/anton-users/pkg/utils"
)

func (uu *UserUsecase) CreateUser(ctx context.Context, request *service.RequestCreateUser) (*model.User, error) {
	pwd, err := hasher.HashPassword(request.Password)
	if err != nil {
		return nil, err
	}

	user := model.User{
		Id:       utils.UniqueID(),
		Username: request.Username,
		Password: pwd,
		Name:     request.Name,
		Dob:      request.Dob,
		Gender:   request.Gender,
	}

	tx, err := uu.ur.CreateTx(ctx)
	if err != nil {
		return nil, err
	}
	defer uu.ur.RollbackTx(ctx, tx)

	err = uu.ur.CreateUser(ctx, tx, &user)
	if err != nil {
		return nil, err
	}

	err = uu.ur.CommitTx(ctx, tx)
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
		Id:   user.Id,
		Name: request.Name,
		Dob:  request.Dob,
	}

	tx, err := uu.ur.CreateTx(ctx)
	if err != nil {
		return nil, err
	}
	defer uu.ur.RollbackTx(ctx, tx)

	err = uu.ur.UpdateUser(ctx, tx, &reqUser)
	if err != nil {
		return nil, err
	}

	err = uu.ur.CommitTx(ctx, tx)
	if err != nil {
		return nil, err
	}

	return &reqUser, nil
}

func (uu *UserUsecase) DeleteUser(ctx context.Context, id string) error {
	tx, err := uu.ur.CreateTx(ctx)
	if err != nil {
		return err
	}
	defer uu.ur.RollbackTx(ctx, tx)

	err = uu.ur.DeleteUser(ctx, tx, id)
	if err != nil {
		return err
	}

	err = uu.ur.CommitTx(ctx, tx)
	if err != nil {
		return err
	}

	return nil
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
