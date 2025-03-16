package user

import (
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestNew(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRP := NewMockiUserRepo(ctrl)
	mockDatabaseUtils := NewMockiDatabaseUtils(ctrl)
	mockRandomUtils := NewMockiRandomUtils(ctrl)
	mockHasherUtils := NewMockiHasherUtils(ctrl)

	type args struct {
		ur  iUserRepo
		dbu iDatabaseUtils
		ru  iRandomUtils
		hu  iHasherUtils
	}
	tests := []struct {
		name string
		args args
		want *UserUsecase
	}{
		{
			name: "success",
			args: args{
				ur:  mockUserRP,
				dbu: mockDatabaseUtils,
				ru:  mockRandomUtils,
				hu:  mockHasherUtils,
			},
			want: &UserUsecase{
				ur:  mockUserRP,
				dbu: mockDatabaseUtils,
				ru:  mockRandomUtils,
				hu:  mockHasherUtils,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.ur, tt.args.dbu, tt.args.ru, tt.args.hu); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
