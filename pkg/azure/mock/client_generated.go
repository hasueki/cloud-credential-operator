// Code generated by MockGen. DO NOT EDIT.
// Source: ./clients.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	authorization "github.com/Azure/azure-sdk-for-go/services/authorization/mgmt/2015-07-01/authorization"
	graphrbac "github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	gomock "github.com/golang/mock/gomock"
)

// MockAppClient is a mock of AppClient interface.
type MockAppClient struct {
	ctrl     *gomock.Controller
	recorder *MockAppClientMockRecorder
}

// MockAppClientMockRecorder is the mock recorder for MockAppClient.
type MockAppClientMockRecorder struct {
	mock *MockAppClient
}

// NewMockAppClient creates a new mock instance.
func NewMockAppClient(ctrl *gomock.Controller) *MockAppClient {
	mock := &MockAppClient{ctrl: ctrl}
	mock.recorder = &MockAppClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAppClient) EXPECT() *MockAppClientMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockAppClient) Create(ctx context.Context, parameters graphrbac.ApplicationCreateParameters) (graphrbac.Application, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, parameters)
	ret0, _ := ret[0].(graphrbac.Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockAppClientMockRecorder) Create(ctx, parameters interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAppClient)(nil).Create), ctx, parameters)
}

// Delete mocks base method.
func (m *MockAppClient) Delete(ctx context.Context, applicationObjectID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, applicationObjectID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockAppClientMockRecorder) Delete(ctx, applicationObjectID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockAppClient)(nil).Delete), ctx, applicationObjectID)
}

// List mocks base method.
func (m *MockAppClient) List(ctx context.Context, filter string) ([]graphrbac.Application, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, filter)
	ret0, _ := ret[0].([]graphrbac.Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockAppClientMockRecorder) List(ctx, filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockAppClient)(nil).List), ctx, filter)
}

// UpdatePasswordCredentials mocks base method.
func (m *MockAppClient) UpdatePasswordCredentials(ctx context.Context, applicationObjectID string, parameters graphrbac.PasswordCredentialsUpdateParameters) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePasswordCredentials", ctx, applicationObjectID, parameters)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePasswordCredentials indicates an expected call of UpdatePasswordCredentials.
func (mr *MockAppClientMockRecorder) UpdatePasswordCredentials(ctx, applicationObjectID, parameters interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePasswordCredentials", reflect.TypeOf((*MockAppClient)(nil).UpdatePasswordCredentials), ctx, applicationObjectID, parameters)
}

// MockServicePrincipalClient is a mock of ServicePrincipalClient interface.
type MockServicePrincipalClient struct {
	ctrl     *gomock.Controller
	recorder *MockServicePrincipalClientMockRecorder
}

// MockServicePrincipalClientMockRecorder is the mock recorder for MockServicePrincipalClient.
type MockServicePrincipalClientMockRecorder struct {
	mock *MockServicePrincipalClient
}

// NewMockServicePrincipalClient creates a new mock instance.
func NewMockServicePrincipalClient(ctrl *gomock.Controller) *MockServicePrincipalClient {
	mock := &MockServicePrincipalClient{ctrl: ctrl}
	mock.recorder = &MockServicePrincipalClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServicePrincipalClient) EXPECT() *MockServicePrincipalClientMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockServicePrincipalClient) Create(ctx context.Context, parameters graphrbac.ServicePrincipalCreateParameters) (graphrbac.ServicePrincipal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, parameters)
	ret0, _ := ret[0].(graphrbac.ServicePrincipal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockServicePrincipalClientMockRecorder) Create(ctx, parameters interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockServicePrincipalClient)(nil).Create), ctx, parameters)
}

// List mocks base method.
func (m *MockServicePrincipalClient) List(ctx context.Context, filter string) ([]graphrbac.ServicePrincipal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, filter)
	ret0, _ := ret[0].([]graphrbac.ServicePrincipal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockServicePrincipalClientMockRecorder) List(ctx, filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockServicePrincipalClient)(nil).List), ctx, filter)
}

// MockRoleAssignmentsClient is a mock of RoleAssignmentsClient interface.
type MockRoleAssignmentsClient struct {
	ctrl     *gomock.Controller
	recorder *MockRoleAssignmentsClientMockRecorder
}

// MockRoleAssignmentsClientMockRecorder is the mock recorder for MockRoleAssignmentsClient.
type MockRoleAssignmentsClientMockRecorder struct {
	mock *MockRoleAssignmentsClient
}

// NewMockRoleAssignmentsClient creates a new mock instance.
func NewMockRoleAssignmentsClient(ctrl *gomock.Controller) *MockRoleAssignmentsClient {
	mock := &MockRoleAssignmentsClient{ctrl: ctrl}
	mock.recorder = &MockRoleAssignmentsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRoleAssignmentsClient) EXPECT() *MockRoleAssignmentsClientMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockRoleAssignmentsClient) Create(ctx context.Context, scope, roleAssignmentName string, parameters authorization.RoleAssignmentCreateParameters) (authorization.RoleAssignment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, scope, roleAssignmentName, parameters)
	ret0, _ := ret[0].(authorization.RoleAssignment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockRoleAssignmentsClientMockRecorder) Create(ctx, scope, roleAssignmentName, parameters interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRoleAssignmentsClient)(nil).Create), ctx, scope, roleAssignmentName, parameters)
}

// DeleteByID mocks base method.
func (m *MockRoleAssignmentsClient) DeleteByID(ctx context.Context, roleAssignmentID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByID", ctx, roleAssignmentID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByID indicates an expected call of DeleteByID.
func (mr *MockRoleAssignmentsClientMockRecorder) DeleteByID(ctx, roleAssignmentID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockRoleAssignmentsClient)(nil).DeleteByID), ctx, roleAssignmentID)
}

// List mocks base method.
func (m *MockRoleAssignmentsClient) List(ctx context.Context, filter string) ([]authorization.RoleAssignment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, filter)
	ret0, _ := ret[0].([]authorization.RoleAssignment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockRoleAssignmentsClientMockRecorder) List(ctx, filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockRoleAssignmentsClient)(nil).List), ctx, filter)
}

// MockRoleDefinitionClient is a mock of RoleDefinitionClient interface.
type MockRoleDefinitionClient struct {
	ctrl     *gomock.Controller
	recorder *MockRoleDefinitionClientMockRecorder
}

// MockRoleDefinitionClientMockRecorder is the mock recorder for MockRoleDefinitionClient.
type MockRoleDefinitionClientMockRecorder struct {
	mock *MockRoleDefinitionClient
}

// NewMockRoleDefinitionClient creates a new mock instance.
func NewMockRoleDefinitionClient(ctrl *gomock.Controller) *MockRoleDefinitionClient {
	mock := &MockRoleDefinitionClient{ctrl: ctrl}
	mock.recorder = &MockRoleDefinitionClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRoleDefinitionClient) EXPECT() *MockRoleDefinitionClientMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockRoleDefinitionClient) List(ctx context.Context, scope, filter string) ([]authorization.RoleDefinition, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, scope, filter)
	ret0, _ := ret[0].([]authorization.RoleDefinition)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockRoleDefinitionClientMockRecorder) List(ctx, scope, filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockRoleDefinitionClient)(nil).List), ctx, scope, filter)
}
