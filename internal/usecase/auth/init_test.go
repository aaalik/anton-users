package auth

import (
	"reflect"
	"testing"

	gomock "github.com/golang/mock/gomock"
)

func TestNew(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRP := NewMockiUserRepo(ctrl)
	mockJWTConfUtils := NewMockiJwtConfUtils(ctrl)
	mockDatabaseUtils := NewMockiDatabaseUtils(ctrl)
	mockRandomUtils := NewMockiRandomUtils(ctrl)
	mockHasherUtils := NewMockiHasherUtils(ctrl)
	mockJWTUtils := NewMockiJwtUtils(ctrl)

	type args struct {
		ur  iUserRepo
		jcu iJwtConfUtils
		dbu iDatabaseUtils
		ru  iRandomUtils
		hu  iHasherUtils
		jwu iJwtUtils
	}
	tests := []struct {
		name string
		args args
		want *AuthUsecase
	}{
		{
			name: "success",
			args: args{
				ur:  mockUserRP,
				jcu: mockJWTConfUtils,
				dbu: mockDatabaseUtils,
				ru:  mockRandomUtils,
				hu:  mockHasherUtils,
				jwu: mockJWTUtils,
			},
			want: &AuthUsecase{
				ur:  mockUserRP,
				jcu: mockJWTConfUtils,
				dbu: mockDatabaseUtils,
				ru:  mockRandomUtils,
				hu:  mockHasherUtils,
				jwu: mockJWTUtils,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.ur, tt.args.jcu, tt.args.dbu, tt.args.ru, tt.args.hu, tt.args.jwu); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
