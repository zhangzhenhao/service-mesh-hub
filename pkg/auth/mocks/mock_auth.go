// Code generated by MockGen. DO NOT EDIT.
// Source: ./interfaces.go

// Package mock_auth is a generated GoMock package.
package mock_auth

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/core/v1"
	v10 "k8s.io/api/rbac/v1"
	v11 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// MockServiceAccountClient is a mock of ServiceAccountClient interface
type MockServiceAccountClient struct {
	ctrl     *gomock.Controller
	recorder *MockServiceAccountClientMockRecorder
}

// MockServiceAccountClientMockRecorder is the mock recorder for MockServiceAccountClient
type MockServiceAccountClientMockRecorder struct {
	mock *MockServiceAccountClient
}

// NewMockServiceAccountClient creates a new mock instance
func NewMockServiceAccountClient(ctrl *gomock.Controller) *MockServiceAccountClient {
	mock := &MockServiceAccountClient{ctrl: ctrl}
	mock.recorder = &MockServiceAccountClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockServiceAccountClient) EXPECT() *MockServiceAccountClientMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockServiceAccountClient) Create(serviceAccount *v1.ServiceAccount) (*v1.ServiceAccount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", serviceAccount)
	ret0, _ := ret[0].(*v1.ServiceAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockServiceAccountClientMockRecorder) Create(serviceAccount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockServiceAccountClient)(nil).Create), serviceAccount)
}

// Get mocks base method
func (m *MockServiceAccountClient) Get(namespace, name string, options v11.GetOptions) (*v1.ServiceAccount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", namespace, name, options)
	ret0, _ := ret[0].(*v1.ServiceAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockServiceAccountClientMockRecorder) Get(namespace, name, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockServiceAccountClient)(nil).Get), namespace, name, options)
}

// Update mocks base method
func (m *MockServiceAccountClient) Update(serviceAccount *v1.ServiceAccount) (*v1.ServiceAccount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", serviceAccount)
	ret0, _ := ret[0].(*v1.ServiceAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockServiceAccountClientMockRecorder) Update(serviceAccount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockServiceAccountClient)(nil).Update), serviceAccount)
}

// MockSecretClient is a mock of SecretClient interface
type MockSecretClient struct {
	ctrl     *gomock.Controller
	recorder *MockSecretClientMockRecorder
}

// MockSecretClientMockRecorder is the mock recorder for MockSecretClient
type MockSecretClientMockRecorder struct {
	mock *MockSecretClient
}

// NewMockSecretClient creates a new mock instance
func NewMockSecretClient(ctrl *gomock.Controller) *MockSecretClient {
	mock := &MockSecretClient{ctrl: ctrl}
	mock.recorder = &MockSecretClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSecretClient) EXPECT() *MockSecretClientMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockSecretClient) Get(namespace, name string, options v11.GetOptions) (*v1.Secret, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", namespace, name, options)
	ret0, _ := ret[0].(*v1.Secret)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockSecretClientMockRecorder) Get(namespace, name, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockSecretClient)(nil).Get), namespace, name, options)
}

// MockRbacClient is a mock of RbacClient interface
type MockRbacClient struct {
	ctrl     *gomock.Controller
	recorder *MockRbacClientMockRecorder
}

// MockRbacClientMockRecorder is the mock recorder for MockRbacClient
type MockRbacClientMockRecorder struct {
	mock *MockRbacClient
}

// NewMockRbacClient creates a new mock instance
func NewMockRbacClient(ctrl *gomock.Controller) *MockRbacClient {
	mock := &MockRbacClient{ctrl: ctrl}
	mock.recorder = &MockRbacClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRbacClient) EXPECT() *MockRbacClientMockRecorder {
	return m.recorder
}

// BindClusterRolesToServiceAccount mocks base method
func (m *MockRbacClient) BindClusterRolesToServiceAccount(targetServiceAccount *v1.ServiceAccount, roles []*v10.ClusterRole) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BindClusterRolesToServiceAccount", targetServiceAccount, roles)
	ret0, _ := ret[0].(error)
	return ret0
}

// BindClusterRolesToServiceAccount indicates an expected call of BindClusterRolesToServiceAccount
func (mr *MockRbacClientMockRecorder) BindClusterRolesToServiceAccount(targetServiceAccount, roles interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BindClusterRolesToServiceAccount", reflect.TypeOf((*MockRbacClient)(nil).BindClusterRolesToServiceAccount), targetServiceAccount, roles)
}