// Code generated by MockGen. DO NOT EDIT.
// Source: service.go
//
// Generated by this command:
//
//	mockgen -source=service.go -destination=mock/service_mock.go
//

// Package mock_service is a generated GoMock package.
package mock_service

import (
	reflect "reflect"

	ethrpc "github.com/g3co/twtest/pkg/ethrpc"
	gomock "go.uber.org/mock/gomock"
)

// MockstorageProvider is a mock of storageProvider interface.
type MockstorageProvider struct {
	ctrl     *gomock.Controller
	recorder *MockstorageProviderMockRecorder
}

// MockstorageProviderMockRecorder is the mock recorder for MockstorageProvider.
type MockstorageProviderMockRecorder struct {
	mock *MockstorageProvider
}

// NewMockstorageProvider creates a new mock instance.
func NewMockstorageProvider(ctrl *gomock.Controller) *MockstorageProvider {
	mock := &MockstorageProvider{ctrl: ctrl}
	mock.recorder = &MockstorageProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockstorageProvider) EXPECT() *MockstorageProviderMockRecorder {
	return m.recorder
}

// AddAddress mocks base method.
func (m *MockstorageProvider) AddAddress(address string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAddress", address)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddAddress indicates an expected call of AddAddress.
func (mr *MockstorageProviderMockRecorder) AddAddress(address any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAddress", reflect.TypeOf((*MockstorageProvider)(nil).AddAddress), address)
}

// GetCurrentBlock mocks base method.
func (m *MockstorageProvider) GetCurrentBlock() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentBlock")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrentBlock indicates an expected call of GetCurrentBlock.
func (mr *MockstorageProviderMockRecorder) GetCurrentBlock() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentBlock", reflect.TypeOf((*MockstorageProvider)(nil).GetCurrentBlock))
}

// GetTXByAddress mocks base method.
func (m *MockstorageProvider) GetTXByAddress(address string) ([]ethrpc.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTXByAddress", address)
	ret0, _ := ret[0].([]ethrpc.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTXByAddress indicates an expected call of GetTXByAddress.
func (mr *MockstorageProviderMockRecorder) GetTXByAddress(address any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTXByAddress", reflect.TypeOf((*MockstorageProvider)(nil).GetTXByAddress), address)
}

// SaveTX mocks base method.
func (m *MockstorageProvider) SaveTX(tx ethrpc.Transaction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveTX", tx)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveTX indicates an expected call of SaveTX.
func (mr *MockstorageProviderMockRecorder) SaveTX(tx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveTX", reflect.TypeOf((*MockstorageProvider)(nil).SaveTX), tx)
}

// MockchainViewerProvider is a mock of chainViewerProvider interface.
type MockchainViewerProvider struct {
	ctrl     *gomock.Controller
	recorder *MockchainViewerProviderMockRecorder
}

// MockchainViewerProviderMockRecorder is the mock recorder for MockchainViewerProvider.
type MockchainViewerProviderMockRecorder struct {
	mock *MockchainViewerProvider
}

// NewMockchainViewerProvider creates a new mock instance.
func NewMockchainViewerProvider(ctrl *gomock.Controller) *MockchainViewerProvider {
	mock := &MockchainViewerProvider{ctrl: ctrl}
	mock.recorder = &MockchainViewerProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockchainViewerProvider) EXPECT() *MockchainViewerProviderMockRecorder {
	return m.recorder
}

// GetBlockInfo mocks base method.
func (m *MockchainViewerProvider) GetBlockInfo(blockNumber string) (*ethrpc.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlockInfo", blockNumber)
	ret0, _ := ret[0].(*ethrpc.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlockInfo indicates an expected call of GetBlockInfo.
func (mr *MockchainViewerProviderMockRecorder) GetBlockInfo(blockNumber any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockInfo", reflect.TypeOf((*MockchainViewerProvider)(nil).GetBlockInfo), blockNumber)
}

// GetCurrentBlock mocks base method.
func (m *MockchainViewerProvider) GetCurrentBlock() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentBlock")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrentBlock indicates an expected call of GetCurrentBlock.
func (mr *MockchainViewerProviderMockRecorder) GetCurrentBlock() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentBlock", reflect.TypeOf((*MockchainViewerProvider)(nil).GetCurrentBlock))
}
