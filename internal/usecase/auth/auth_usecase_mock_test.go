// Code generated by MockGen. DO NOT EDIT.
// Source: type.go

// Package auth is a generated GoMock package.
package auth

import (
	context "context"
	reflect "reflect"

	model "github.com/aaalik/anton-users/internal/model"
	jwt "github.com/dgrijalva/jwt-go"
	gomock "github.com/golang/mock/gomock"
	sqlx "github.com/jmoiron/sqlx"
)

// MockiUserRepo is a mock of iUserRepo interface.
type MockiUserRepo struct {
	ctrl     *gomock.Controller
	recorder *MockiUserRepoMockRecorder
}

// MockiUserRepoMockRecorder is the mock recorder for MockiUserRepo.
type MockiUserRepoMockRecorder struct {
	mock *MockiUserRepo
}

// NewMockiUserRepo creates a new mock instance.
func NewMockiUserRepo(ctrl *gomock.Controller) *MockiUserRepo {
	mock := &MockiUserRepo{ctrl: ctrl}
	mock.recorder = &MockiUserRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockiUserRepo) EXPECT() *MockiUserRepoMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockiUserRepo) CreateUser(ctx context.Context, tx *sqlx.Tx, user *model.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, tx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockiUserRepoMockRecorder) CreateUser(ctx, tx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockiUserRepo)(nil).CreateUser), ctx, tx, user)
}

// FetchUserLogin mocks base method.
func (m *MockiUserRepo) FetchUserLogin(ctx context.Context, username string) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchUserLogin", ctx, username)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchUserLogin indicates an expected call of FetchUserLogin.
func (mr *MockiUserRepoMockRecorder) FetchUserLogin(ctx, username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchUserLogin", reflect.TypeOf((*MockiUserRepo)(nil).FetchUserLogin), ctx, username)
}

// MockiJwtConfUtils is a mock of iJwtConfUtils interface.
type MockiJwtConfUtils struct {
	ctrl     *gomock.Controller
	recorder *MockiJwtConfUtilsMockRecorder
}

// MockiJwtConfUtilsMockRecorder is the mock recorder for MockiJwtConfUtils.
type MockiJwtConfUtilsMockRecorder struct {
	mock *MockiJwtConfUtils
}

// NewMockiJwtConfUtils creates a new mock instance.
func NewMockiJwtConfUtils(ctrl *gomock.Controller) *MockiJwtConfUtils {
	mock := &MockiJwtConfUtils{ctrl: ctrl}
	mock.recorder = &MockiJwtConfUtilsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockiJwtConfUtils) EXPECT() *MockiJwtConfUtilsMockRecorder {
	return m.recorder
}

// GetExpire mocks base method.
func (m *MockiJwtConfUtils) GetExpire() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExpire")
	ret0, _ := ret[0].(int)
	return ret0
}

// GetExpire indicates an expected call of GetExpire.
func (mr *MockiJwtConfUtilsMockRecorder) GetExpire() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExpire", reflect.TypeOf((*MockiJwtConfUtils)(nil).GetExpire))
}

// GetSecret mocks base method.
func (m *MockiJwtConfUtils) GetSecret() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecret")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetSecret indicates an expected call of GetSecret.
func (mr *MockiJwtConfUtilsMockRecorder) GetSecret() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecret", reflect.TypeOf((*MockiJwtConfUtils)(nil).GetSecret))
}

// GetSecretRefresh mocks base method.
func (m *MockiJwtConfUtils) GetSecretRefresh() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecretRefresh")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetSecretRefresh indicates an expected call of GetSecretRefresh.
func (mr *MockiJwtConfUtilsMockRecorder) GetSecretRefresh() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecretRefresh", reflect.TypeOf((*MockiJwtConfUtils)(nil).GetSecretRefresh))
}

// MockiDatabaseUtils is a mock of iDatabaseUtils interface.
type MockiDatabaseUtils struct {
	ctrl     *gomock.Controller
	recorder *MockiDatabaseUtilsMockRecorder
}

// MockiDatabaseUtilsMockRecorder is the mock recorder for MockiDatabaseUtils.
type MockiDatabaseUtilsMockRecorder struct {
	mock *MockiDatabaseUtils
}

// NewMockiDatabaseUtils creates a new mock instance.
func NewMockiDatabaseUtils(ctrl *gomock.Controller) *MockiDatabaseUtils {
	mock := &MockiDatabaseUtils{ctrl: ctrl}
	mock.recorder = &MockiDatabaseUtilsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockiDatabaseUtils) EXPECT() *MockiDatabaseUtilsMockRecorder {
	return m.recorder
}

// ExecuteTx mocks base method.
func (m *MockiDatabaseUtils) ExecuteTx(ctx context.Context, tx *sqlx.Tx, fn func(context.Context, *sqlx.Tx) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExecuteTx", ctx, tx, fn)
	ret0, _ := ret[0].(error)
	return ret0
}

// ExecuteTx indicates an expected call of ExecuteTx.
func (mr *MockiDatabaseUtilsMockRecorder) ExecuteTx(ctx, tx, fn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecuteTx", reflect.TypeOf((*MockiDatabaseUtils)(nil).ExecuteTx), ctx, tx, fn)
}

// MockiRandomUtils is a mock of iRandomUtils interface.
type MockiRandomUtils struct {
	ctrl     *gomock.Controller
	recorder *MockiRandomUtilsMockRecorder
}

// MockiRandomUtilsMockRecorder is the mock recorder for MockiRandomUtils.
type MockiRandomUtilsMockRecorder struct {
	mock *MockiRandomUtils
}

// NewMockiRandomUtils creates a new mock instance.
func NewMockiRandomUtils(ctrl *gomock.Controller) *MockiRandomUtils {
	mock := &MockiRandomUtils{ctrl: ctrl}
	mock.recorder = &MockiRandomUtilsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockiRandomUtils) EXPECT() *MockiRandomUtilsMockRecorder {
	return m.recorder
}

// UniqueID mocks base method.
func (m *MockiRandomUtils) UniqueID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UniqueID")
	ret0, _ := ret[0].(string)
	return ret0
}

// UniqueID indicates an expected call of UniqueID.
func (mr *MockiRandomUtilsMockRecorder) UniqueID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UniqueID", reflect.TypeOf((*MockiRandomUtils)(nil).UniqueID))
}

// MockiHasherUtils is a mock of iHasherUtils interface.
type MockiHasherUtils struct {
	ctrl     *gomock.Controller
	recorder *MockiHasherUtilsMockRecorder
}

// MockiHasherUtilsMockRecorder is the mock recorder for MockiHasherUtils.
type MockiHasherUtilsMockRecorder struct {
	mock *MockiHasherUtils
}

// NewMockiHasherUtils creates a new mock instance.
func NewMockiHasherUtils(ctrl *gomock.Controller) *MockiHasherUtils {
	mock := &MockiHasherUtils{ctrl: ctrl}
	mock.recorder = &MockiHasherUtilsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockiHasherUtils) EXPECT() *MockiHasherUtilsMockRecorder {
	return m.recorder
}

// CheckPasswordHash mocks base method.
func (m *MockiHasherUtils) CheckPasswordHash(password, hash string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckPasswordHash", password, hash)
	ret0, _ := ret[0].(bool)
	return ret0
}

// CheckPasswordHash indicates an expected call of CheckPasswordHash.
func (mr *MockiHasherUtilsMockRecorder) CheckPasswordHash(password, hash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckPasswordHash", reflect.TypeOf((*MockiHasherUtils)(nil).CheckPasswordHash), password, hash)
}

// HashPassword mocks base method.
func (m *MockiHasherUtils) HashPassword(password string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HashPassword", password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HashPassword indicates an expected call of HashPassword.
func (mr *MockiHasherUtilsMockRecorder) HashPassword(password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HashPassword", reflect.TypeOf((*MockiHasherUtils)(nil).HashPassword), password)
}

// MockiJwtUtils is a mock of iJwtUtils interface.
type MockiJwtUtils struct {
	ctrl     *gomock.Controller
	recorder *MockiJwtUtilsMockRecorder
}

// MockiJwtUtilsMockRecorder is the mock recorder for MockiJwtUtils.
type MockiJwtUtilsMockRecorder struct {
	mock *MockiJwtUtils
}

// NewMockiJwtUtils creates a new mock instance.
func NewMockiJwtUtils(ctrl *gomock.Controller) *MockiJwtUtils {
	mock := &MockiJwtUtils{ctrl: ctrl}
	mock.recorder = &MockiJwtUtilsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockiJwtUtils) EXPECT() *MockiJwtUtilsMockRecorder {
	return m.recorder
}

// GenerateToken mocks base method.
func (m *MockiJwtUtils) GenerateToken(claims jwt.MapClaims, secretKey string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateToken", claims, secretKey)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateToken indicates an expected call of GenerateToken.
func (mr *MockiJwtUtilsMockRecorder) GenerateToken(claims, secretKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateToken", reflect.TypeOf((*MockiJwtUtils)(nil).GenerateToken), claims, secretKey)
}

// ParseToken mocks base method.
func (m *MockiJwtUtils) ParseToken(tokenString, secretKey string) (*jwt.Token, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseToken", tokenString, secretKey)
	ret0, _ := ret[0].(*jwt.Token)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseToken indicates an expected call of ParseToken.
func (mr *MockiJwtUtilsMockRecorder) ParseToken(tokenString, secretKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseToken", reflect.TypeOf((*MockiJwtUtils)(nil).ParseToken), tokenString, secretKey)
}
