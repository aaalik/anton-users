package user

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/aaalik/anton-users/internal/model"
	"github.com/aaalik/anton-users/internal/service"
	"github.com/aaalik/anton-users/pkg/testfiles"
	"github.com/golang/mock/gomock"
)

func TestUserUsecase_CreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRP := NewMockiUserRepo(ctrl)
	mockDatabaseUtils := NewMockiDatabaseUtils(ctrl)
	mockHasherUtils := NewMockiHasherUtils(ctrl)
	mockRandomUtils := NewMockiRandomUtils(ctrl)

	reqCreate := service.RequestCreateUser{
		Username: "uname",
		Password: "123123",
		Name:     "name",
	}

	reqCreatePasswordFailed := reqCreate
	reqCreatePasswordFailed.Password = "asdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasdf"

	user := &model.User{
		Username: "uname",
		Password: "123123",
		Name:     "name",
	}

	tests := []struct {
		name    string
		req     *service.RequestCreateUser
		mock    func()
		want    *model.User
		wantErr bool
	}{
		{
			name: "success",
			req:  &reqCreate,
			mock: func() {
				mockHasherUtils.EXPECT().HashPassword(reqCreate.Password).Return(user.Password, nil)
				mockRandomUtils.EXPECT().UniqueID().Return(user.Id)
				mockDatabaseUtils.EXPECT().ExecuteTx(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(testfiles.DoExecuteTx)
				mockUserRP.EXPECT().CreateUser(gomock.Any(), gomock.Any(), user).Return(nil)
			},
			want:    user,
			wantErr: false,
		},
		{
			name: "failed - hash password",
			req:  &reqCreatePasswordFailed,
			mock: func() {
				mockHasherUtils.EXPECT().HashPassword(reqCreatePasswordFailed.Password).Return("", errors.New(gomock.Any().String()))
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "failed - execute tx",
			req:  &reqCreate,
			mock: func() {
				mockHasherUtils.EXPECT().HashPassword(reqCreate.Password).Return(user.Password, nil)
				mockRandomUtils.EXPECT().UniqueID().Return(user.Id)
				mockDatabaseUtils.EXPECT().ExecuteTx(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New(gomock.Any().String()))
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "failed - create user",
			req:  &reqCreate,
			mock: func() {
				mockHasherUtils.EXPECT().HashPassword(reqCreate.Password).Return(user.Password, nil)
				mockRandomUtils.EXPECT().UniqueID().Return(user.Id)
				mockDatabaseUtils.EXPECT().ExecuteTx(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(testfiles.DoExecuteTx)
				mockUserRP.EXPECT().CreateUser(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New(gomock.Any().String()))
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uu := &UserUsecase{
				ur:  mockUserRP,
				dbu: mockDatabaseUtils,
				hu:  mockHasherUtils,
				ru:  mockRandomUtils,
			}

			if tt.mock != nil {
				tt.mock()
			}

			got, err := uu.CreateUser(context.Background(), tt.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserUsecase.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserUsecase.CreateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserUsecase_UpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRP := NewMockiUserRepo(ctrl)
	mockDatabaseUtils := NewMockiDatabaseUtils(ctrl)
	mockHasherUtils := NewMockiHasherUtils(ctrl)
	mockRandomUtils := NewMockiRandomUtils(ctrl)

	reqUpdate := service.RequestUpdateUser{
		Id:     "1",
		Name:   "name",
		Dob:    "2003-01-02",
		Gender: model.USER_GENDER_MALE,
	}

	user := &model.User{
		Id:     "1",
		Name:   "name",
		Dob:    "2003-01-02",
		Gender: model.USER_GENDER_MALE,
	}

	tests := []struct {
		name    string
		req     *service.RequestUpdateUser
		mock    func()
		want    *model.User
		wantErr bool
	}{
		{
			name: "success",
			req:  &reqUpdate,
			mock: func() {
				mockUserRP.EXPECT().DetailUser(gomock.Any(), reqUpdate.Id).Return(user, nil)
				mockDatabaseUtils.EXPECT().ExecuteTx(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(testfiles.DoExecuteTx)
				mockUserRP.EXPECT().UpdateUser(gomock.Any(), gomock.Any(), user).Return(nil)
			},
			want:    user,
			wantErr: false,
		},
		{
			name: "failed - get detail user",
			req:  &reqUpdate,
			mock: func() {
				mockUserRP.EXPECT().DetailUser(gomock.Any(), reqUpdate.Id).Return(nil, errors.New(gomock.Any().String()))
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "failed - execute tx",
			req:  &reqUpdate,
			mock: func() {
				mockUserRP.EXPECT().DetailUser(gomock.Any(), reqUpdate.Id).Return(user, nil)
				mockDatabaseUtils.EXPECT().ExecuteTx(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New(gomock.Any().String()))
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "failed - update user",
			req:  &reqUpdate,
			mock: func() {
				mockUserRP.EXPECT().DetailUser(gomock.Any(), reqUpdate.Id).Return(user, nil)
				mockDatabaseUtils.EXPECT().ExecuteTx(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(testfiles.DoExecuteTx)
				mockUserRP.EXPECT().UpdateUser(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New(gomock.Any().String()))
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uu := &UserUsecase{
				ur:  mockUserRP,
				dbu: mockDatabaseUtils,
				hu:  mockHasherUtils,
				ru:  mockRandomUtils,
			}

			if tt.mock != nil {
				tt.mock()
			}

			got, err := uu.UpdateUser(context.Background(), tt.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserUsecase.UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserUsecase.UpdateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserUsecase_DeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRP := NewMockiUserRepo(ctrl)
	mockDatabaseUtils := NewMockiDatabaseUtils(ctrl)

	id := "1"

	tests := []struct {
		name    string
		id      string
		mock    func()
		wantErr bool
	}{
		{
			name: "success",
			id:   id,
			mock: func() {
				mockDatabaseUtils.EXPECT().ExecuteTx(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(testfiles.DoExecuteTx)
				mockUserRP.EXPECT().DeleteUser(gomock.Any(), gomock.Any(), id).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "failed - delete user",
			id:   id,
			mock: func() {
				mockDatabaseUtils.EXPECT().ExecuteTx(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(testfiles.DoExecuteTx)
				mockUserRP.EXPECT().DeleteUser(gomock.Any(), gomock.Any(), id).Return(errors.New(gomock.Any().String()))
			},
			wantErr: true,
		},
		{
			name: "failed - execute query",
			id:   id,
			mock: func() {
				mockDatabaseUtils.EXPECT().ExecuteTx(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New(gomock.Any().String()))
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uu := &UserUsecase{
				ur:  mockUserRP,
				dbu: mockDatabaseUtils,
			}

			if tt.mock != nil {
				tt.mock()
			}

			if err := uu.DeleteUser(context.Background(), tt.id); (err != nil) != tt.wantErr {
				t.Errorf("UserUsecase.DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserUsecase_DetailUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRP := NewMockiUserRepo(ctrl)

	id := "1"

	user := &model.User{
		Id:        "1",
		Username:  "uname",
		Name:      "name",
		Dob:       "2003-01-02",
		Gender:    model.USER_GENDER_MALE,
		CreatedAt: 1,
		UpdatedAt: 1,
		DeletedAt: 0,
	}

	tests := []struct {
		name    string
		id      string
		mock    func()
		want    *model.User
		wantErr bool
	}{
		{
			name: "success",
			id:   id,
			mock: func() {
				mockUserRP.EXPECT().DetailUser(gomock.Any(), id).Return(user, nil)
			},
			want:    user,
			wantErr: false,
		},
		{
			name: "failed",
			id:   id,
			mock: func() {
				mockUserRP.EXPECT().DetailUser(gomock.Any(), id).Return(nil, errors.New(gomock.Any().String()))
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uu := &UserUsecase{
				ur: mockUserRP,
			}

			if tt.mock != nil {
				tt.mock()
			}

			got, err := uu.DetailUser(context.Background(), tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserUsecase.DetailUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserUsecase.DetailUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserUsecase_ListUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRP := NewMockiUserRepo(ctrl)

	reqList := &service.RequestListUser{
		Includes: &service.RequestFilterUser{
			Ids:  []string{"1"},
			Dobs: []string{"2003-01-02"},
			CreatedAt: &service.Range{
				GTE: 0,
				LTE: 10,
			},
			DeletedAt: &service.Range{
				GTE: 0,
				LTE: 10,
			},
		},
		Queries: &service.Queries{
			Page: 1,
			Rows: 10,
			Sort: &service.SortBy{
				Field: "created_at",
				Order: service.ORDER_DESC,
			},
			Keyword:     "name",
			WithDeleted: []bool{true},
		},
	}

	users := []*model.User{
		{
			Id:        "1",
			Username:  "uname",
			Name:      "name",
			Dob:       "2003-01-02",
			Gender:    model.USER_GENDER_MALE,
			CreatedAt: 1,
			UpdatedAt: 1,
			DeletedAt: 0,
		},
		{
			Id:        "2",
			Username:  "uname2",
			Name:      "name2",
			Dob:       "2003-01-02",
			Gender:    model.USER_GENDER_MALE,
			CreatedAt: 1,
			UpdatedAt: 1,
			DeletedAt: 10,
		},
	}

	count := int32(len(users))

	tests := []struct {
		name    string
		req     *service.RequestListUser
		mock    func()
		want    []*model.User
		want1   int32
		wantErr bool
	}{
		{
			name: "success",
			req:  reqList,
			mock: func() {
				mockUserRP.EXPECT().ListUser(gomock.Any(), reqList).Return(users, nil)
				mockUserRP.EXPECT().CountUsers(gomock.Any(), reqList).Return(count, nil)
			},
			want:    users,
			want1:   count,
			wantErr: false,
		},
		{
			name: "failed - get list user",
			req:  reqList,
			mock: func() {
				mockUserRP.EXPECT().ListUser(gomock.Any(), reqList).Return(nil, errors.New(gomock.Any().String()))
			},
			want:    nil,
			want1:   0,
			wantErr: true,
		},
		{
			name: "failed - get count user",
			req:  reqList,
			mock: func() {
				mockUserRP.EXPECT().ListUser(gomock.Any(), reqList).Return(users, nil)
				mockUserRP.EXPECT().CountUsers(gomock.Any(), reqList).Return(int32(0), errors.New(gomock.Any().String()))
			},
			want:    nil,
			want1:   0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uu := &UserUsecase{
				ur: mockUserRP,
			}

			if tt.mock != nil {
				tt.mock()
			}

			got, got1, err := uu.ListUser(context.Background(), tt.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserUsecase.ListUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserUsecase.ListUser() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("UserUsecase.ListUser() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
