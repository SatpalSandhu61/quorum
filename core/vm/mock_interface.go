// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package vm is a generated GoMock package.
package vm

import (
	big "math/big"
	reflect "reflect"

	common "github.com/ethereum/go-ethereum/common"
	state "github.com/ethereum/go-ethereum/core/state"
	types "github.com/ethereum/go-ethereum/core/types"
	gomock "github.com/golang/mock/gomock"
)

// MockAccountExtraDataStateReader is a mock of AccountExtraDataStateReader interface.
type MockAccountExtraDataStateReader struct {
	ctrl     *gomock.Controller
	recorder *MockAccountExtraDataStateReaderMockRecorder
}

// MockAccountExtraDataStateReaderMockRecorder is the mock recorder for MockAccountExtraDataStateReader.
type MockAccountExtraDataStateReaderMockRecorder struct {
	mock *MockAccountExtraDataStateReader
}

// NewMockAccountExtraDataStateReader creates a new mock instance.
func NewMockAccountExtraDataStateReader(ctrl *gomock.Controller) *MockAccountExtraDataStateReader {
	mock := &MockAccountExtraDataStateReader{ctrl: ctrl}
	mock.recorder = &MockAccountExtraDataStateReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountExtraDataStateReader) EXPECT() *MockAccountExtraDataStateReaderMockRecorder {
	return m.recorder
}

// ReadPrivacyMetadata mocks base method.
func (m *MockAccountExtraDataStateReader) ReadPrivacyMetadata(addr common.Address) (*state.PrivacyMetadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadPrivacyMetadata", addr)
	ret0, _ := ret[0].(*state.PrivacyMetadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadPrivacyMetadata indicates an expected call of ReadPrivacyMetadata.
func (mr *MockAccountExtraDataStateReaderMockRecorder) ReadPrivacyMetadata(addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadPrivacyMetadata", reflect.TypeOf((*MockAccountExtraDataStateReader)(nil).ReadPrivacyMetadata), addr)
}

// ReadManagedParties mocks base method.
func (m *MockAccountExtraDataStateReader) ReadManagedParties(addr common.Address) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadManagedParties", addr)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadManagedParties indicates an expected call of ReadManagedParties.
func (mr *MockAccountExtraDataStateReaderMockRecorder) ReadManagedParties(addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadManagedParties", reflect.TypeOf((*MockAccountExtraDataStateReader)(nil).ReadManagedParties), addr)
}

// MockAccountExtraDataStateWriter is a mock of AccountExtraDataStateWriter interface.
type MockAccountExtraDataStateWriter struct {
	ctrl     *gomock.Controller
	recorder *MockAccountExtraDataStateWriterMockRecorder
}

// MockAccountExtraDataStateWriterMockRecorder is the mock recorder for MockAccountExtraDataStateWriter.
type MockAccountExtraDataStateWriterMockRecorder struct {
	mock *MockAccountExtraDataStateWriter
}

// NewMockAccountExtraDataStateWriter creates a new mock instance.
func NewMockAccountExtraDataStateWriter(ctrl *gomock.Controller) *MockAccountExtraDataStateWriter {
	mock := &MockAccountExtraDataStateWriter{ctrl: ctrl}
	mock.recorder = &MockAccountExtraDataStateWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountExtraDataStateWriter) EXPECT() *MockAccountExtraDataStateWriterMockRecorder {
	return m.recorder
}

// WritePrivacyMetadata mocks base method.
func (m *MockAccountExtraDataStateWriter) WritePrivacyMetadata(addr common.Address, pm *state.PrivacyMetadata) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "WritePrivacyMetadata", addr, pm)
}

// WritePrivacyMetadata indicates an expected call of WritePrivacyMetadata.
func (mr *MockAccountExtraDataStateWriterMockRecorder) WritePrivacyMetadata(addr, pm interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WritePrivacyMetadata", reflect.TypeOf((*MockAccountExtraDataStateWriter)(nil).WritePrivacyMetadata), addr, pm)
}

// WriteManagedParties mocks base method.
func (m *MockAccountExtraDataStateWriter) WriteManagedParties(addr common.Address, managedParties []string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "WriteManagedParties", addr, managedParties)
}

// WriteManagedParties indicates an expected call of WriteManagedParties.
func (mr *MockAccountExtraDataStateWriterMockRecorder) WriteManagedParties(addr, managedParties interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteManagedParties", reflect.TypeOf((*MockAccountExtraDataStateWriter)(nil).WriteManagedParties), addr, managedParties)
}

// MockMinimalApiState is a mock of MinimalApiState interface.
type MockMinimalApiState struct {
	ctrl     *gomock.Controller
	recorder *MockMinimalApiStateMockRecorder
}

// MockMinimalApiStateMockRecorder is the mock recorder for MockMinimalApiState.
type MockMinimalApiStateMockRecorder struct {
	mock *MockMinimalApiState
}

// NewMockMinimalApiState creates a new mock instance.
func NewMockMinimalApiState(ctrl *gomock.Controller) *MockMinimalApiState {
	mock := &MockMinimalApiState{ctrl: ctrl}
	mock.recorder = &MockMinimalApiStateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMinimalApiState) EXPECT() *MockMinimalApiStateMockRecorder {
	return m.recorder
}

// ReadPrivacyMetadata mocks base method.
func (m *MockMinimalApiState) ReadPrivacyMetadata(addr common.Address) (*state.PrivacyMetadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadPrivacyMetadata", addr)
	ret0, _ := ret[0].(*state.PrivacyMetadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadPrivacyMetadata indicates an expected call of ReadPrivacyMetadata.
func (mr *MockMinimalApiStateMockRecorder) ReadPrivacyMetadata(addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadPrivacyMetadata", reflect.TypeOf((*MockMinimalApiState)(nil).ReadPrivacyMetadata), addr)
}

// ReadManagedParties mocks base method.
func (m *MockMinimalApiState) ReadManagedParties(addr common.Address) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadManagedParties", addr)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadManagedParties indicates an expected call of ReadManagedParties.
func (mr *MockMinimalApiStateMockRecorder) ReadManagedParties(addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadManagedParties", reflect.TypeOf((*MockMinimalApiState)(nil).ReadManagedParties), addr)
}

// GetBalance mocks base method.
func (m *MockMinimalApiState) GetBalance(addr common.Address) *big.Int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBalance", addr)
	ret0, _ := ret[0].(*big.Int)
	return ret0
}

// GetBalance indicates an expected call of GetBalance.
func (mr *MockMinimalApiStateMockRecorder) GetBalance(addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBalance", reflect.TypeOf((*MockMinimalApiState)(nil).GetBalance), addr)
}

// SetBalance mocks base method.
func (m *MockMinimalApiState) SetBalance(addr common.Address, balance *big.Int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetBalance", addr, balance)
}

// SetBalance indicates an expected call of SetBalance.
func (mr *MockMinimalApiStateMockRecorder) SetBalance(addr, balance interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetBalance", reflect.TypeOf((*MockMinimalApiState)(nil).SetBalance), addr, balance)
}

// GetCode mocks base method.
func (m *MockMinimalApiState) GetCode(addr common.Address) []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCode", addr)
	ret0, _ := ret[0].([]byte)
	return ret0
}

// GetCode indicates an expected call of GetCode.
func (mr *MockMinimalApiStateMockRecorder) GetCode(addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCode", reflect.TypeOf((*MockMinimalApiState)(nil).GetCode), addr)
}

// GetState mocks base method.
func (m *MockMinimalApiState) GetState(a common.Address, b common.Hash) common.Hash {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetState", a, b)
	ret0, _ := ret[0].(common.Hash)
	return ret0
}

// GetState indicates an expected call of GetState.
func (mr *MockMinimalApiStateMockRecorder) GetState(a, b interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetState", reflect.TypeOf((*MockMinimalApiState)(nil).GetState), a, b)
}

// GetNonce mocks base method.
func (m *MockMinimalApiState) GetNonce(addr common.Address) uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNonce", addr)
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GetNonce indicates an expected call of GetNonce.
func (mr *MockMinimalApiStateMockRecorder) GetNonce(addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNonce", reflect.TypeOf((*MockMinimalApiState)(nil).GetNonce), addr)
}

// SetNonce mocks base method.
func (m *MockMinimalApiState) SetNonce(addr common.Address, nonce uint64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetNonce", addr, nonce)
}

// SetNonce indicates an expected call of SetNonce.
func (mr *MockMinimalApiStateMockRecorder) SetNonce(addr, nonce interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetNonce", reflect.TypeOf((*MockMinimalApiState)(nil).SetNonce), addr, nonce)
}

// SetCode mocks base method.
func (m *MockMinimalApiState) SetCode(arg0 common.Address, arg1 []byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetCode", arg0, arg1)
}

// SetCode indicates an expected call of SetCode.
func (mr *MockMinimalApiStateMockRecorder) SetCode(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCode", reflect.TypeOf((*MockMinimalApiState)(nil).SetCode), arg0, arg1)
}

// GetRLPEncodedStateObject mocks base method.
func (m *MockMinimalApiState) GetRLPEncodedStateObject(addr common.Address) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRLPEncodedStateObject", addr)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRLPEncodedStateObject indicates an expected call of GetRLPEncodedStateObject.
func (mr *MockMinimalApiStateMockRecorder) GetRLPEncodedStateObject(addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRLPEncodedStateObject", reflect.TypeOf((*MockMinimalApiState)(nil).GetRLPEncodedStateObject), addr)
}

// GetProof mocks base method.
func (m *MockMinimalApiState) GetProof(arg0 common.Address) ([][]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProof", arg0)
	ret0, _ := ret[0].([][]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProof indicates an expected call of GetProof.
func (mr *MockMinimalApiStateMockRecorder) GetProof(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProof", reflect.TypeOf((*MockMinimalApiState)(nil).GetProof), arg0)
}

// GetStorageProof mocks base method.
func (m *MockMinimalApiState) GetStorageProof(arg0 common.Address, arg1 common.Hash) ([][]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStorageProof", arg0, arg1)
	ret0, _ := ret[0].([][]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStorageProof indicates an expected call of GetStorageProof.
func (mr *MockMinimalApiStateMockRecorder) GetStorageProof(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStorageProof", reflect.TypeOf((*MockMinimalApiState)(nil).GetStorageProof), arg0, arg1)
}

// StorageTrie mocks base method.
func (m *MockMinimalApiState) StorageTrie(addr common.Address) state.Trie {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StorageTrie", addr)
	ret0, _ := ret[0].(state.Trie)
	return ret0
}

// StorageTrie indicates an expected call of StorageTrie.
func (mr *MockMinimalApiStateMockRecorder) StorageTrie(addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StorageTrie", reflect.TypeOf((*MockMinimalApiState)(nil).StorageTrie), addr)
}

// Error mocks base method.
func (m *MockMinimalApiState) Error() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Error")
	ret0, _ := ret[0].(error)
	return ret0
}

// Error indicates an expected call of Error.
func (mr *MockMinimalApiStateMockRecorder) Error() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockMinimalApiState)(nil).Error))
}

// GetCodeHash mocks base method.
func (m *MockMinimalApiState) GetCodeHash(arg0 common.Address) common.Hash {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCodeHash", arg0)
	ret0, _ := ret[0].(common.Hash)
	return ret0
}

// GetCodeHash indicates an expected call of GetCodeHash.
func (mr *MockMinimalApiStateMockRecorder) GetCodeHash(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCodeHash", reflect.TypeOf((*MockMinimalApiState)(nil).GetCodeHash), arg0)
}

// SetState mocks base method.
func (m *MockMinimalApiState) SetState(arg0 common.Address, arg1, arg2 common.Hash) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetState", arg0, arg1, arg2)
}

// SetState indicates an expected call of SetState.
func (mr *MockMinimalApiStateMockRecorder) SetState(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetState", reflect.TypeOf((*MockMinimalApiState)(nil).SetState), arg0, arg1, arg2)
}

// SetStorage mocks base method.
func (m *MockMinimalApiState) SetStorage(addr common.Address, storage map[common.Hash]common.Hash) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetStorage", addr, storage)
}

// SetStorage indicates an expected call of SetStorage.
func (mr *MockMinimalApiStateMockRecorder) SetStorage(addr, storage interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetStorage", reflect.TypeOf((*MockMinimalApiState)(nil).SetStorage), addr, storage)
}

// MockStateDB is a mock of StateDB interface.
type MockStateDB struct {
	ctrl     *gomock.Controller
	recorder *MockStateDBMockRecorder
}

// MockStateDBMockRecorder is the mock recorder for MockStateDB.
type MockStateDBMockRecorder struct {
	mock *MockStateDB
}

// NewMockStateDB creates a new mock instance.
func NewMockStateDB(ctrl *gomock.Controller) *MockStateDB {
	mock := &MockStateDB{ctrl: ctrl}
	mock.recorder = &MockStateDBMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStateDB) EXPECT() *MockStateDBMockRecorder {
	return m.recorder
}

// ReadPrivacyMetadata mocks base method.
func (m *MockStateDB) ReadPrivacyMetadata(addr common.Address) (*state.PrivacyMetadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadPrivacyMetadata", addr)
	ret0, _ := ret[0].(*state.PrivacyMetadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadPrivacyMetadata indicates an expected call of ReadPrivacyMetadata.
func (mr *MockStateDBMockRecorder) ReadPrivacyMetadata(addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadPrivacyMetadata", reflect.TypeOf((*MockStateDB)(nil).ReadPrivacyMetadata), addr)
}

// ReadManagedParties mocks base method.
func (m *MockStateDB) ReadManagedParties(addr common.Address) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadManagedParties", addr)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadManagedParties indicates an expected call of ReadManagedParties.
func (mr *MockStateDBMockRecorder) ReadManagedParties(addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadManagedParties", reflect.TypeOf((*MockStateDB)(nil).ReadManagedParties), addr)
}

// GetBalance mocks base method.
func (m *MockStateDB) GetBalance(addr common.Address) *big.Int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBalance", addr)
	ret0, _ := ret[0].(*big.Int)
	return ret0
}

// GetBalance indicates an expected call of GetBalance.
func (mr *MockStateDBMockRecorder) GetBalance(addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBalance", reflect.TypeOf((*MockStateDB)(nil).GetBalance), addr)
}

// SetBalance mocks base method.
func (m *MockStateDB) SetBalance(addr common.Address, balance *big.Int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetBalance", addr, balance)
}

// SetBalance indicates an expected call of SetBalance.
func (mr *MockStateDBMockRecorder) SetBalance(addr, balance interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetBalance", reflect.TypeOf((*MockStateDB)(nil).SetBalance), addr, balance)
}

// GetCode mocks base method.
func (m *MockStateDB) GetCode(addr common.Address) []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCode", addr)
	ret0, _ := ret[0].([]byte)
	return ret0
}

// GetCode indicates an expected call of GetCode.
func (mr *MockStateDBMockRecorder) GetCode(addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCode", reflect.TypeOf((*MockStateDB)(nil).GetCode), addr)
}

// GetState mocks base method.
func (m *MockStateDB) GetState(a common.Address, b common.Hash) common.Hash {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetState", a, b)
	ret0, _ := ret[0].(common.Hash)
	return ret0
}

// GetState indicates an expected call of GetState.
func (mr *MockStateDBMockRecorder) GetState(a, b interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetState", reflect.TypeOf((*MockStateDB)(nil).GetState), a, b)
}

// GetNonce mocks base method.
func (m *MockStateDB) GetNonce(addr common.Address) uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNonce", addr)
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GetNonce indicates an expected call of GetNonce.
func (mr *MockStateDBMockRecorder) GetNonce(addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNonce", reflect.TypeOf((*MockStateDB)(nil).GetNonce), addr)
}

// SetNonce mocks base method.
func (m *MockStateDB) SetNonce(addr common.Address, nonce uint64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetNonce", addr, nonce)
}

// SetNonce indicates an expected call of SetNonce.
func (mr *MockStateDBMockRecorder) SetNonce(addr, nonce interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetNonce", reflect.TypeOf((*MockStateDB)(nil).SetNonce), addr, nonce)
}

// SetCode mocks base method.
func (m *MockStateDB) SetCode(arg0 common.Address, arg1 []byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetCode", arg0, arg1)
}

// SetCode indicates an expected call of SetCode.
func (mr *MockStateDBMockRecorder) SetCode(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCode", reflect.TypeOf((*MockStateDB)(nil).SetCode), arg0, arg1)
}

// GetRLPEncodedStateObject mocks base method.
func (m *MockStateDB) GetRLPEncodedStateObject(addr common.Address) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRLPEncodedStateObject", addr)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRLPEncodedStateObject indicates an expected call of GetRLPEncodedStateObject.
func (mr *MockStateDBMockRecorder) GetRLPEncodedStateObject(addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRLPEncodedStateObject", reflect.TypeOf((*MockStateDB)(nil).GetRLPEncodedStateObject), addr)
}

// GetProof mocks base method.
func (m *MockStateDB) GetProof(arg0 common.Address) ([][]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProof", arg0)
	ret0, _ := ret[0].([][]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProof indicates an expected call of GetProof.
func (mr *MockStateDBMockRecorder) GetProof(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProof", reflect.TypeOf((*MockStateDB)(nil).GetProof), arg0)
}

// GetStorageProof mocks base method.
func (m *MockStateDB) GetStorageProof(arg0 common.Address, arg1 common.Hash) ([][]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStorageProof", arg0, arg1)
	ret0, _ := ret[0].([][]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStorageProof indicates an expected call of GetStorageProof.
func (mr *MockStateDBMockRecorder) GetStorageProof(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStorageProof", reflect.TypeOf((*MockStateDB)(nil).GetStorageProof), arg0, arg1)
}

// StorageTrie mocks base method.
func (m *MockStateDB) StorageTrie(addr common.Address) state.Trie {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StorageTrie", addr)
	ret0, _ := ret[0].(state.Trie)
	return ret0
}

// StorageTrie indicates an expected call of StorageTrie.
func (mr *MockStateDBMockRecorder) StorageTrie(addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StorageTrie", reflect.TypeOf((*MockStateDB)(nil).StorageTrie), addr)
}

// Error mocks base method.
func (m *MockStateDB) Error() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Error")
	ret0, _ := ret[0].(error)
	return ret0
}

// Error indicates an expected call of Error.
func (mr *MockStateDBMockRecorder) Error() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockStateDB)(nil).Error))
}

// GetCodeHash mocks base method.
func (m *MockStateDB) GetCodeHash(arg0 common.Address) common.Hash {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCodeHash", arg0)
	ret0, _ := ret[0].(common.Hash)
	return ret0
}

// GetCodeHash indicates an expected call of GetCodeHash.
func (mr *MockStateDBMockRecorder) GetCodeHash(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCodeHash", reflect.TypeOf((*MockStateDB)(nil).GetCodeHash), arg0)
}

// SetState mocks base method.
func (m *MockStateDB) SetState(arg0 common.Address, arg1, arg2 common.Hash) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetState", arg0, arg1, arg2)
}

// SetState indicates an expected call of SetState.
func (mr *MockStateDBMockRecorder) SetState(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetState", reflect.TypeOf((*MockStateDB)(nil).SetState), arg0, arg1, arg2)
}

// SetStorage mocks base method.
func (m *MockStateDB) SetStorage(addr common.Address, storage map[common.Hash]common.Hash) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetStorage", addr, storage)
}

// SetStorage indicates an expected call of SetStorage.
func (mr *MockStateDBMockRecorder) SetStorage(addr, storage interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetStorage", reflect.TypeOf((*MockStateDB)(nil).SetStorage), addr, storage)
}

// WritePrivacyMetadata mocks base method.
func (m *MockStateDB) WritePrivacyMetadata(addr common.Address, pm *state.PrivacyMetadata) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "WritePrivacyMetadata", addr, pm)
}

// WritePrivacyMetadata indicates an expected call of WritePrivacyMetadata.
func (mr *MockStateDBMockRecorder) WritePrivacyMetadata(addr, pm interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WritePrivacyMetadata", reflect.TypeOf((*MockStateDB)(nil).WritePrivacyMetadata), addr, pm)
}

// WriteManagedParties mocks base method.
func (m *MockStateDB) WriteManagedParties(addr common.Address, managedParties []string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "WriteManagedParties", addr, managedParties)
}

// WriteManagedParties indicates an expected call of WriteManagedParties.
func (mr *MockStateDBMockRecorder) WriteManagedParties(addr, managedParties interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteManagedParties", reflect.TypeOf((*MockStateDB)(nil).WriteManagedParties), addr, managedParties)
}

// CreateAccount mocks base method.
func (m *MockStateDB) CreateAccount(arg0 common.Address) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CreateAccount", arg0)
}

// CreateAccount indicates an expected call of CreateAccount.
func (mr *MockStateDBMockRecorder) CreateAccount(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccount", reflect.TypeOf((*MockStateDB)(nil).CreateAccount), arg0)
}

// SubBalance mocks base method.
func (m *MockStateDB) SubBalance(arg0 common.Address, arg1 *big.Int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SubBalance", arg0, arg1)
}

// SubBalance indicates an expected call of SubBalance.
func (mr *MockStateDBMockRecorder) SubBalance(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubBalance", reflect.TypeOf((*MockStateDB)(nil).SubBalance), arg0, arg1)
}

// AddBalance mocks base method.
func (m *MockStateDB) AddBalance(arg0 common.Address, arg1 *big.Int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddBalance", arg0, arg1)
}

// AddBalance indicates an expected call of AddBalance.
func (mr *MockStateDBMockRecorder) AddBalance(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddBalance", reflect.TypeOf((*MockStateDB)(nil).AddBalance), arg0, arg1)
}

// GetCodeSize mocks base method.
func (m *MockStateDB) GetCodeSize(arg0 common.Address) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCodeSize", arg0)
	ret0, _ := ret[0].(int)
	return ret0
}

// GetCodeSize indicates an expected call of GetCodeSize.
func (mr *MockStateDBMockRecorder) GetCodeSize(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCodeSize", reflect.TypeOf((*MockStateDB)(nil).GetCodeSize), arg0)
}

// AddRefund mocks base method.
func (m *MockStateDB) AddRefund(arg0 uint64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddRefund", arg0)
}

// AddRefund indicates an expected call of AddRefund.
func (mr *MockStateDBMockRecorder) AddRefund(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRefund", reflect.TypeOf((*MockStateDB)(nil).AddRefund), arg0)
}

// SubRefund mocks base method.
func (m *MockStateDB) SubRefund(arg0 uint64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SubRefund", arg0)
}

// SubRefund indicates an expected call of SubRefund.
func (mr *MockStateDBMockRecorder) SubRefund(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubRefund", reflect.TypeOf((*MockStateDB)(nil).SubRefund), arg0)
}

// GetRefund mocks base method.
func (m *MockStateDB) GetRefund() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRefund")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GetRefund indicates an expected call of GetRefund.
func (mr *MockStateDBMockRecorder) GetRefund() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRefund", reflect.TypeOf((*MockStateDB)(nil).GetRefund))
}

// GetCommittedState mocks base method.
func (m *MockStateDB) GetCommittedState(arg0 common.Address, arg1 common.Hash) common.Hash {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCommittedState", arg0, arg1)
	ret0, _ := ret[0].(common.Hash)
	return ret0
}

// GetCommittedState indicates an expected call of GetCommittedState.
func (mr *MockStateDBMockRecorder) GetCommittedState(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCommittedState", reflect.TypeOf((*MockStateDB)(nil).GetCommittedState), arg0, arg1)
}

// Suicide mocks base method.
func (m *MockStateDB) Suicide(arg0 common.Address) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Suicide", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Suicide indicates an expected call of Suicide.
func (mr *MockStateDBMockRecorder) Suicide(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Suicide", reflect.TypeOf((*MockStateDB)(nil).Suicide), arg0)
}

// HasSuicided mocks base method.
func (m *MockStateDB) HasSuicided(arg0 common.Address) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasSuicided", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasSuicided indicates an expected call of HasSuicided.
func (mr *MockStateDBMockRecorder) HasSuicided(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasSuicided", reflect.TypeOf((*MockStateDB)(nil).HasSuicided), arg0)
}

// Exist mocks base method.
func (m *MockStateDB) Exist(arg0 common.Address) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exist", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Exist indicates an expected call of Exist.
func (mr *MockStateDBMockRecorder) Exist(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exist", reflect.TypeOf((*MockStateDB)(nil).Exist), arg0)
}

// Empty mocks base method.
func (m *MockStateDB) Empty(arg0 common.Address) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Empty", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Empty indicates an expected call of Empty.
func (mr *MockStateDBMockRecorder) Empty(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Empty", reflect.TypeOf((*MockStateDB)(nil).Empty), arg0)
}

// RevertToSnapshot mocks base method.
func (m *MockStateDB) RevertToSnapshot(arg0 int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RevertToSnapshot", arg0)
}

// RevertToSnapshot indicates an expected call of RevertToSnapshot.
func (mr *MockStateDBMockRecorder) RevertToSnapshot(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevertToSnapshot", reflect.TypeOf((*MockStateDB)(nil).RevertToSnapshot), arg0)
}

// Snapshot mocks base method.
func (m *MockStateDB) Snapshot() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Snapshot")
	ret0, _ := ret[0].(int)
	return ret0
}

// Snapshot indicates an expected call of Snapshot.
func (mr *MockStateDBMockRecorder) Snapshot() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Snapshot", reflect.TypeOf((*MockStateDB)(nil).Snapshot))
}

// AddLog mocks base method.
func (m *MockStateDB) AddLog(arg0 *types.Log) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddLog", arg0)
}

// AddLog indicates an expected call of AddLog.
func (mr *MockStateDBMockRecorder) AddLog(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddLog", reflect.TypeOf((*MockStateDB)(nil).AddLog), arg0)
}

// AddPreimage mocks base method.
func (m *MockStateDB) AddPreimage(arg0 common.Hash, arg1 []byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddPreimage", arg0, arg1)
}

// AddPreimage indicates an expected call of AddPreimage.
func (mr *MockStateDBMockRecorder) AddPreimage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPreimage", reflect.TypeOf((*MockStateDB)(nil).AddPreimage), arg0, arg1)
}

// ForEachStorage mocks base method.
func (m *MockStateDB) ForEachStorage(arg0 common.Address, arg1 func(common.Hash, common.Hash) bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForEachStorage", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ForEachStorage indicates an expected call of ForEachStorage.
func (mr *MockStateDBMockRecorder) ForEachStorage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForEachStorage", reflect.TypeOf((*MockStateDB)(nil).ForEachStorage), arg0, arg1)
}

// MockCallContext is a mock of CallContext interface.
type MockCallContext struct {
	ctrl     *gomock.Controller
	recorder *MockCallContextMockRecorder
}

// MockCallContextMockRecorder is the mock recorder for MockCallContext.
type MockCallContextMockRecorder struct {
	mock *MockCallContext
}

// NewMockCallContext creates a new mock instance.
func NewMockCallContext(ctrl *gomock.Controller) *MockCallContext {
	mock := &MockCallContext{ctrl: ctrl}
	mock.recorder = &MockCallContextMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCallContext) EXPECT() *MockCallContextMockRecorder {
	return m.recorder
}

// Call mocks base method.
func (m *MockCallContext) Call(env *EVM, me ContractRef, addr common.Address, data []byte, gas, value *big.Int) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Call", env, me, addr, data, gas, value)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Call indicates an expected call of Call.
func (mr *MockCallContextMockRecorder) Call(env, me, addr, data, gas, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Call", reflect.TypeOf((*MockCallContext)(nil).Call), env, me, addr, data, gas, value)
}

// CallCode mocks base method.
func (m *MockCallContext) CallCode(env *EVM, me ContractRef, addr common.Address, data []byte, gas, value *big.Int) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CallCode", env, me, addr, data, gas, value)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CallCode indicates an expected call of CallCode.
func (mr *MockCallContextMockRecorder) CallCode(env, me, addr, data, gas, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CallCode", reflect.TypeOf((*MockCallContext)(nil).CallCode), env, me, addr, data, gas, value)
}

// DelegateCall mocks base method.
func (m *MockCallContext) DelegateCall(env *EVM, me ContractRef, addr common.Address, data []byte, gas *big.Int) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DelegateCall", env, me, addr, data, gas)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DelegateCall indicates an expected call of DelegateCall.
func (mr *MockCallContextMockRecorder) DelegateCall(env, me, addr, data, gas interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DelegateCall", reflect.TypeOf((*MockCallContext)(nil).DelegateCall), env, me, addr, data, gas)
}

// Create mocks base method.
func (m *MockCallContext) Create(env *EVM, me ContractRef, data []byte, gas, value *big.Int) ([]byte, common.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", env, me, data, gas, value)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(common.Address)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Create indicates an expected call of Create.
func (mr *MockCallContextMockRecorder) Create(env, me, data, gas, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCallContext)(nil).Create), env, me, data, gas, value)
}
