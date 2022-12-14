// Code generated by MockGen. DO NOT EDIT.
// Source: movie.trpc.go

// Package pb_hub is a generated GoMock package.
package pb_hub

import (
	context "context"
	client "git.code.oa.com/trpc-go/trpc-go/client"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockUserService is a mock of UserService interface
type MockUserService struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceMockRecorder
}

// MockUserServiceMockRecorder is the mock recorder for MockUserService
type MockUserServiceMockRecorder struct {
	mock *MockUserService
}

// NewMockUserService creates a new mock instance
func NewMockUserService(ctrl *gomock.Controller) *MockUserService {
	mock := &MockUserService{ctrl: ctrl}
	mock.recorder = &MockUserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserService) EXPECT() *MockUserServiceMockRecorder {
	return m.recorder
}

// Register mocks base method
func (m *MockUserService) Register(ctx context.Context, req *RegisterReq, rsp *RegisterRsp) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", ctx, req, rsp)
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register
func (mr *MockUserServiceMockRecorder) Register(ctx, req, rsp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockUserService)(nil).Register), ctx, req, rsp)
}

// CheckUserName mocks base method
func (m *MockUserService) CheckUserName(ctx context.Context, req *CheckUserNameReq, rsp *CheckUserNameRsp) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckUserName", ctx, req, rsp)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckUserName indicates an expected call of CheckUserName
func (mr *MockUserServiceMockRecorder) CheckUserName(ctx, req, rsp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckUserName", reflect.TypeOf((*MockUserService)(nil).CheckUserName), ctx, req, rsp)
}

// Login mocks base method
func (m *MockUserService) Login(ctx context.Context, req *LoginReq, rsp *LoginRsp) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", ctx, req, rsp)
	ret0, _ := ret[0].(error)
	return ret0
}

// Login indicates an expected call of Login
func (mr *MockUserServiceMockRecorder) Login(ctx, req, rsp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockUserService)(nil).Login), ctx, req, rsp)
}

// MockListService is a mock of ListService interface
type MockListService struct {
	ctrl     *gomock.Controller
	recorder *MockListServiceMockRecorder
}

// MockListServiceMockRecorder is the mock recorder for MockListService
type MockListServiceMockRecorder struct {
	mock *MockListService
}

// NewMockListService creates a new mock instance
func NewMockListService(ctrl *gomock.Controller) *MockListService {
	mock := &MockListService{ctrl: ctrl}
	mock.recorder = &MockListServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockListService) EXPECT() *MockListServiceMockRecorder {
	return m.recorder
}

// GetList mocks base method
func (m *MockListService) GetList(ctx context.Context, req *GetListReq, rsp *GetListRsp) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetList", ctx, req, rsp)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetList indicates an expected call of GetList
func (mr *MockListServiceMockRecorder) GetList(ctx, req, rsp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetList", reflect.TypeOf((*MockListService)(nil).GetList), ctx, req, rsp)
}

// GetLeaderboard mocks base method
func (m *MockListService) GetLeaderboard(ctx context.Context, req *GetLeaderboardReq, rsp *GetLeaderboardRsp) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLeaderboard", ctx, req, rsp)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetLeaderboard indicates an expected call of GetLeaderboard
func (mr *MockListServiceMockRecorder) GetLeaderboard(ctx, req, rsp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLeaderboard", reflect.TypeOf((*MockListService)(nil).GetLeaderboard), ctx, req, rsp)
}

// MockInfoService is a mock of InfoService interface
type MockInfoService struct {
	ctrl     *gomock.Controller
	recorder *MockInfoServiceMockRecorder
}

// MockInfoServiceMockRecorder is the mock recorder for MockInfoService
type MockInfoServiceMockRecorder struct {
	mock *MockInfoService
}

// NewMockInfoService creates a new mock instance
func NewMockInfoService(ctrl *gomock.Controller) *MockInfoService {
	mock := &MockInfoService{ctrl: ctrl}
	mock.recorder = &MockInfoServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInfoService) EXPECT() *MockInfoServiceMockRecorder {
	return m.recorder
}

// GetInfo mocks base method
func (m *MockInfoService) GetInfo(ctx context.Context, req *GetInfoReq, rsp *GetInfoRsp) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInfo", ctx, req, rsp)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetInfo indicates an expected call of GetInfo
func (mr *MockInfoServiceMockRecorder) GetInfo(ctx, req, rsp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInfo", reflect.TypeOf((*MockInfoService)(nil).GetInfo), ctx, req, rsp)
}

// GetRecord mocks base method
func (m *MockInfoService) GetRecord(ctx context.Context, req *RecordReq, rsp *RecordRsp) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRecord", ctx, req, rsp)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetRecord indicates an expected call of GetRecord
func (mr *MockInfoServiceMockRecorder) GetRecord(ctx, req, rsp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRecord", reflect.TypeOf((*MockInfoService)(nil).GetRecord), ctx, req, rsp)
}

// AddRecord mocks base method
func (m *MockInfoService) AddRecord(ctx context.Context, req *RecordReq, rsp *RecordRsp) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRecord", ctx, req, rsp)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddRecord indicates an expected call of AddRecord
func (mr *MockInfoServiceMockRecorder) AddRecord(ctx, req, rsp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRecord", reflect.TypeOf((*MockInfoService)(nil).AddRecord), ctx, req, rsp)
}

// DelRecord mocks base method
func (m *MockInfoService) DelRecord(ctx context.Context, req *RecordReq, rsp *RecordRsp) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DelRecord", ctx, req, rsp)
	ret0, _ := ret[0].(error)
	return ret0
}

// DelRecord indicates an expected call of DelRecord
func (mr *MockInfoServiceMockRecorder) DelRecord(ctx, req, rsp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DelRecord", reflect.TypeOf((*MockInfoService)(nil).DelRecord), ctx, req, rsp)
}

// MockManageService is a mock of ManageService interface
type MockManageService struct {
	ctrl     *gomock.Controller
	recorder *MockManageServiceMockRecorder
}

// MockManageServiceMockRecorder is the mock recorder for MockManageService
type MockManageServiceMockRecorder struct {
	mock *MockManageService
}

// NewMockManageService creates a new mock instance
func NewMockManageService(ctrl *gomock.Controller) *MockManageService {
	mock := &MockManageService{ctrl: ctrl}
	mock.recorder = &MockManageServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockManageService) EXPECT() *MockManageServiceMockRecorder {
	return m.recorder
}

// AddInfo mocks base method
func (m *MockManageService) AddInfo(ctx context.Context, req *AddInfoReq, rsp *ManageInfoRsp) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddInfo", ctx, req, rsp)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddInfo indicates an expected call of AddInfo
func (mr *MockManageServiceMockRecorder) AddInfo(ctx, req, rsp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddInfo", reflect.TypeOf((*MockManageService)(nil).AddInfo), ctx, req, rsp)
}

// UpdateInfo mocks base method
func (m *MockManageService) UpdateInfo(ctx context.Context, req *UpdateInfoReq, rsp *ManageInfoRsp) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateInfo", ctx, req, rsp)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateInfo indicates an expected call of UpdateInfo
func (mr *MockManageServiceMockRecorder) UpdateInfo(ctx, req, rsp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateInfo", reflect.TypeOf((*MockManageService)(nil).UpdateInfo), ctx, req, rsp)
}

// UpdateInfoStatus mocks base method
func (m *MockManageService) UpdateInfoStatus(ctx context.Context, req *UpdateInfoStatusReq, rsp *ManageInfoRsp) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateInfoStatus", ctx, req, rsp)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateInfoStatus indicates an expected call of UpdateInfoStatus
func (mr *MockManageServiceMockRecorder) UpdateInfoStatus(ctx, req, rsp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateInfoStatus", reflect.TypeOf((*MockManageService)(nil).UpdateInfoStatus), ctx, req, rsp)
}

// DelInfo mocks base method
func (m *MockManageService) DelInfo(ctx context.Context, req *DelInfoReq, rsp *ManageInfoRsp) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DelInfo", ctx, req, rsp)
	ret0, _ := ret[0].(error)
	return ret0
}

// DelInfo indicates an expected call of DelInfo
func (mr *MockManageServiceMockRecorder) DelInfo(ctx, req, rsp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DelInfo", reflect.TypeOf((*MockManageService)(nil).DelInfo), ctx, req, rsp)
}

// AddProp mocks base method
func (m *MockManageService) AddProp(ctx context.Context, req *AddPropReq, rsp *ManagePropRsp) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddProp", ctx, req, rsp)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddProp indicates an expected call of AddProp
func (mr *MockManageServiceMockRecorder) AddProp(ctx, req, rsp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddProp", reflect.TypeOf((*MockManageService)(nil).AddProp), ctx, req, rsp)
}

// GetProp mocks base method
func (m *MockManageService) GetProp(ctx context.Context, req *GetPropReq, rsp *GetPropRsp) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProp", ctx, req, rsp)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetProp indicates an expected call of GetProp
func (mr *MockManageServiceMockRecorder) GetProp(ctx, req, rsp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProp", reflect.TypeOf((*MockManageService)(nil).GetProp), ctx, req, rsp)
}

// UpdateProp mocks base method
func (m *MockManageService) UpdateProp(ctx context.Context, req *UpdatePropReq, rsp *ManagePropRsp) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProp", ctx, req, rsp)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProp indicates an expected call of UpdateProp
func (mr *MockManageServiceMockRecorder) UpdateProp(ctx, req, rsp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProp", reflect.TypeOf((*MockManageService)(nil).UpdateProp), ctx, req, rsp)
}

// DelProp mocks base method
func (m *MockManageService) DelProp(ctx context.Context, req *DelPropReq, rsp *ManagePropRsp) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DelProp", ctx, req, rsp)
	ret0, _ := ret[0].(error)
	return ret0
}

// DelProp indicates an expected call of DelProp
func (mr *MockManageServiceMockRecorder) DelProp(ctx, req, rsp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DelProp", reflect.TypeOf((*MockManageService)(nil).DelProp), ctx, req, rsp)
}

// MockUserClientProxy is a mock of UserClientProxy interface
type MockUserClientProxy struct {
	ctrl     *gomock.Controller
	recorder *MockUserClientProxyMockRecorder
}

// MockUserClientProxyMockRecorder is the mock recorder for MockUserClientProxy
type MockUserClientProxyMockRecorder struct {
	mock *MockUserClientProxy
}

// NewMockUserClientProxy creates a new mock instance
func NewMockUserClientProxy(ctrl *gomock.Controller) *MockUserClientProxy {
	mock := &MockUserClientProxy{ctrl: ctrl}
	mock.recorder = &MockUserClientProxyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserClientProxy) EXPECT() *MockUserClientProxyMockRecorder {
	return m.recorder
}

// Register mocks base method
func (m *MockUserClientProxy) Register(ctx context.Context, req *RegisterReq, opts ...client.Option) (*RegisterRsp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, req}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Register", varargs...)
	ret0, _ := ret[0].(*RegisterRsp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Register indicates an expected call of Register
func (mr *MockUserClientProxyMockRecorder) Register(ctx, req interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, req}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockUserClientProxy)(nil).Register), varargs...)
}

// CheckUserName mocks base method
func (m *MockUserClientProxy) CheckUserName(ctx context.Context, req *CheckUserNameReq, opts ...client.Option) (*CheckUserNameRsp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, req}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CheckUserName", varargs...)
	ret0, _ := ret[0].(*CheckUserNameRsp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckUserName indicates an expected call of CheckUserName
func (mr *MockUserClientProxyMockRecorder) CheckUserName(ctx, req interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, req}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckUserName", reflect.TypeOf((*MockUserClientProxy)(nil).CheckUserName), varargs...)
}

// Login mocks base method
func (m *MockUserClientProxy) Login(ctx context.Context, req *LoginReq, opts ...client.Option) (*LoginRsp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, req}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Login", varargs...)
	ret0, _ := ret[0].(*LoginRsp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login
func (mr *MockUserClientProxyMockRecorder) Login(ctx, req interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, req}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockUserClientProxy)(nil).Login), varargs...)
}

// MockListClientProxy is a mock of ListClientProxy interface
type MockListClientProxy struct {
	ctrl     *gomock.Controller
	recorder *MockListClientProxyMockRecorder
}

// MockListClientProxyMockRecorder is the mock recorder for MockListClientProxy
type MockListClientProxyMockRecorder struct {
	mock *MockListClientProxy
}

// NewMockListClientProxy creates a new mock instance
func NewMockListClientProxy(ctrl *gomock.Controller) *MockListClientProxy {
	mock := &MockListClientProxy{ctrl: ctrl}
	mock.recorder = &MockListClientProxyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockListClientProxy) EXPECT() *MockListClientProxyMockRecorder {
	return m.recorder
}

// GetList mocks base method
func (m *MockListClientProxy) GetList(ctx context.Context, req *GetListReq, opts ...client.Option) (*GetListRsp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, req}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetList", varargs...)
	ret0, _ := ret[0].(*GetListRsp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetList indicates an expected call of GetList
func (mr *MockListClientProxyMockRecorder) GetList(ctx, req interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, req}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetList", reflect.TypeOf((*MockListClientProxy)(nil).GetList), varargs...)
}

// GetLeaderboard mocks base method
func (m *MockListClientProxy) GetLeaderboard(ctx context.Context, req *GetLeaderboardReq, opts ...client.Option) (*GetLeaderboardRsp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, req}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetLeaderboard", varargs...)
	ret0, _ := ret[0].(*GetLeaderboardRsp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLeaderboard indicates an expected call of GetLeaderboard
func (mr *MockListClientProxyMockRecorder) GetLeaderboard(ctx, req interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, req}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLeaderboard", reflect.TypeOf((*MockListClientProxy)(nil).GetLeaderboard), varargs...)
}

// MockInfoClientProxy is a mock of InfoClientProxy interface
type MockInfoClientProxy struct {
	ctrl     *gomock.Controller
	recorder *MockInfoClientProxyMockRecorder
}

// MockInfoClientProxyMockRecorder is the mock recorder for MockInfoClientProxy
type MockInfoClientProxyMockRecorder struct {
	mock *MockInfoClientProxy
}

// NewMockInfoClientProxy creates a new mock instance
func NewMockInfoClientProxy(ctrl *gomock.Controller) *MockInfoClientProxy {
	mock := &MockInfoClientProxy{ctrl: ctrl}
	mock.recorder = &MockInfoClientProxyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInfoClientProxy) EXPECT() *MockInfoClientProxyMockRecorder {
	return m.recorder
}

// GetInfo mocks base method
func (m *MockInfoClientProxy) GetInfo(ctx context.Context, req *GetInfoReq, opts ...client.Option) (*GetInfoRsp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, req}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetInfo", varargs...)
	ret0, _ := ret[0].(*GetInfoRsp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInfo indicates an expected call of GetInfo
func (mr *MockInfoClientProxyMockRecorder) GetInfo(ctx, req interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, req}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInfo", reflect.TypeOf((*MockInfoClientProxy)(nil).GetInfo), varargs...)
}

// GetRecord mocks base method
func (m *MockInfoClientProxy) GetRecord(ctx context.Context, req *RecordReq, opts ...client.Option) (*RecordRsp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, req}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRecord", varargs...)
	ret0, _ := ret[0].(*RecordRsp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRecord indicates an expected call of GetRecord
func (mr *MockInfoClientProxyMockRecorder) GetRecord(ctx, req interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, req}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRecord", reflect.TypeOf((*MockInfoClientProxy)(nil).GetRecord), varargs...)
}

// AddRecord mocks base method
func (m *MockInfoClientProxy) AddRecord(ctx context.Context, req *RecordReq, opts ...client.Option) (*RecordRsp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, req}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddRecord", varargs...)
	ret0, _ := ret[0].(*RecordRsp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddRecord indicates an expected call of AddRecord
func (mr *MockInfoClientProxyMockRecorder) AddRecord(ctx, req interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, req}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRecord", reflect.TypeOf((*MockInfoClientProxy)(nil).AddRecord), varargs...)
}

// DelRecord mocks base method
func (m *MockInfoClientProxy) DelRecord(ctx context.Context, req *RecordReq, opts ...client.Option) (*RecordRsp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, req}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DelRecord", varargs...)
	ret0, _ := ret[0].(*RecordRsp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DelRecord indicates an expected call of DelRecord
func (mr *MockInfoClientProxyMockRecorder) DelRecord(ctx, req interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, req}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DelRecord", reflect.TypeOf((*MockInfoClientProxy)(nil).DelRecord), varargs...)
}

// MockManageClientProxy is a mock of ManageClientProxy interface
type MockManageClientProxy struct {
	ctrl     *gomock.Controller
	recorder *MockManageClientProxyMockRecorder
}

// MockManageClientProxyMockRecorder is the mock recorder for MockManageClientProxy
type MockManageClientProxyMockRecorder struct {
	mock *MockManageClientProxy
}

// NewMockManageClientProxy creates a new mock instance
func NewMockManageClientProxy(ctrl *gomock.Controller) *MockManageClientProxy {
	mock := &MockManageClientProxy{ctrl: ctrl}
	mock.recorder = &MockManageClientProxyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockManageClientProxy) EXPECT() *MockManageClientProxyMockRecorder {
	return m.recorder
}

// AddInfo mocks base method
func (m *MockManageClientProxy) AddInfo(ctx context.Context, req *AddInfoReq, opts ...client.Option) (*ManageInfoRsp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, req}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddInfo", varargs...)
	ret0, _ := ret[0].(*ManageInfoRsp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddInfo indicates an expected call of AddInfo
func (mr *MockManageClientProxyMockRecorder) AddInfo(ctx, req interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, req}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddInfo", reflect.TypeOf((*MockManageClientProxy)(nil).AddInfo), varargs...)
}

// UpdateInfo mocks base method
func (m *MockManageClientProxy) UpdateInfo(ctx context.Context, req *UpdateInfoReq, opts ...client.Option) (*ManageInfoRsp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, req}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateInfo", varargs...)
	ret0, _ := ret[0].(*ManageInfoRsp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateInfo indicates an expected call of UpdateInfo
func (mr *MockManageClientProxyMockRecorder) UpdateInfo(ctx, req interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, req}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateInfo", reflect.TypeOf((*MockManageClientProxy)(nil).UpdateInfo), varargs...)
}

// UpdateInfoStatus mocks base method
func (m *MockManageClientProxy) UpdateInfoStatus(ctx context.Context, req *UpdateInfoStatusReq, opts ...client.Option) (*ManageInfoRsp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, req}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateInfoStatus", varargs...)
	ret0, _ := ret[0].(*ManageInfoRsp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateInfoStatus indicates an expected call of UpdateInfoStatus
func (mr *MockManageClientProxyMockRecorder) UpdateInfoStatus(ctx, req interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, req}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateInfoStatus", reflect.TypeOf((*MockManageClientProxy)(nil).UpdateInfoStatus), varargs...)
}

// DelInfo mocks base method
func (m *MockManageClientProxy) DelInfo(ctx context.Context, req *DelInfoReq, opts ...client.Option) (*ManageInfoRsp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, req}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DelInfo", varargs...)
	ret0, _ := ret[0].(*ManageInfoRsp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DelInfo indicates an expected call of DelInfo
func (mr *MockManageClientProxyMockRecorder) DelInfo(ctx, req interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, req}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DelInfo", reflect.TypeOf((*MockManageClientProxy)(nil).DelInfo), varargs...)
}

// AddProp mocks base method
func (m *MockManageClientProxy) AddProp(ctx context.Context, req *AddPropReq, opts ...client.Option) (*ManagePropRsp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, req}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddProp", varargs...)
	ret0, _ := ret[0].(*ManagePropRsp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddProp indicates an expected call of AddProp
func (mr *MockManageClientProxyMockRecorder) AddProp(ctx, req interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, req}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddProp", reflect.TypeOf((*MockManageClientProxy)(nil).AddProp), varargs...)
}

// GetProp mocks base method
func (m *MockManageClientProxy) GetProp(ctx context.Context, req *GetPropReq, opts ...client.Option) (*GetPropRsp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, req}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetProp", varargs...)
	ret0, _ := ret[0].(*GetPropRsp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProp indicates an expected call of GetProp
func (mr *MockManageClientProxyMockRecorder) GetProp(ctx, req interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, req}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProp", reflect.TypeOf((*MockManageClientProxy)(nil).GetProp), varargs...)
}

// UpdateProp mocks base method
func (m *MockManageClientProxy) UpdateProp(ctx context.Context, req *UpdatePropReq, opts ...client.Option) (*ManagePropRsp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, req}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateProp", varargs...)
	ret0, _ := ret[0].(*ManagePropRsp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateProp indicates an expected call of UpdateProp
func (mr *MockManageClientProxyMockRecorder) UpdateProp(ctx, req interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, req}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProp", reflect.TypeOf((*MockManageClientProxy)(nil).UpdateProp), varargs...)
}

// DelProp mocks base method
func (m *MockManageClientProxy) DelProp(ctx context.Context, req *DelPropReq, opts ...client.Option) (*ManagePropRsp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, req}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DelProp", varargs...)
	ret0, _ := ret[0].(*ManagePropRsp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DelProp indicates an expected call of DelProp
func (mr *MockManageClientProxyMockRecorder) DelProp(ctx, req interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, req}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DelProp", reflect.TypeOf((*MockManageClientProxy)(nil).DelProp), varargs...)
}
