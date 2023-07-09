// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/repository/interface/user.go

// Package repository is a generated GoMock package.
package repository

import (
	reflect "reflect"

	models "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"
	gomock "github.com/golang/mock/gomock"
)

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// AddAddress mocks base method.
func (m *MockUserRepository) AddAddress(address models.AddressInfo, userID int) ([]models.AddressInfoResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAddress", address, userID)
	ret0, _ := ret[0].([]models.AddressInfoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddAddress indicates an expected call of AddAddress.
func (mr *MockUserRepositoryMockRecorder) AddAddress(address, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAddress", reflect.TypeOf((*MockUserRepository)(nil).AddAddress), address, userID)
}

// AddToWishList mocks base method.
func (m *MockUserRepository) AddToWishList(userID, productID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddToWishList", userID, productID)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddToWishList indicates an expected call of AddToWishList.
func (mr *MockUserRepositoryMockRecorder) AddToWishList(userID, productID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddToWishList", reflect.TypeOf((*MockUserRepository)(nil).AddToWishList), userID, productID)
}

// ApplyReferral mocks base method.
func (m *MockUserRepository) ApplyReferral(userID int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplyReferral", userID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ApplyReferral indicates an expected call of ApplyReferral.
func (mr *MockUserRepositoryMockRecorder) ApplyReferral(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyReferral", reflect.TypeOf((*MockUserRepository)(nil).ApplyReferral), userID)
}

// CheckUserAvailability mocks base method.
func (m *MockUserRepository) CheckUserAvailability(email string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckUserAvailability", email)
	ret0, _ := ret[0].(bool)
	return ret0
}

// CheckUserAvailability indicates an expected call of CheckUserAvailability.
func (mr *MockUserRepositoryMockRecorder) CheckUserAvailability(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckUserAvailability", reflect.TypeOf((*MockUserRepository)(nil).CheckUserAvailability), email)
}

// CreateReferralEntry mocks base method.
func (m *MockUserRepository) CreateReferralEntry(users models.UserDetailsResponse, userReferral, referralCode string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateReferralEntry", users, userReferral, referralCode)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateReferralEntry indicates an expected call of CreateReferralEntry.
func (mr *MockUserRepositoryMockRecorder) CreateReferralEntry(users, userReferral, referralCode interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateReferralEntry", reflect.TypeOf((*MockUserRepository)(nil).CreateReferralEntry), users, userReferral, referralCode)
}

// FindUserAddressByOrderID mocks base method.
func (m *MockUserRepository) FindUserAddressByOrderID(orderID string) (models.AddressInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserAddressByOrderID", orderID)
	ret0, _ := ret[0].(models.AddressInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserAddressByOrderID indicates an expected call of FindUserAddressByOrderID.
func (mr *MockUserRepositoryMockRecorder) FindUserAddressByOrderID(orderID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserAddressByOrderID", reflect.TypeOf((*MockUserRepository)(nil).FindUserAddressByOrderID), orderID)
}

// FindUserByEmail mocks base method.
func (m *MockUserRepository) FindUserByEmail(user models.UserLogin) (models.UserSignInResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByEmail", user)
	ret0, _ := ret[0].(models.UserSignInResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByEmail indicates an expected call of FindUserByEmail.
func (mr *MockUserRepositoryMockRecorder) FindUserByEmail(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByEmail", reflect.TypeOf((*MockUserRepository)(nil).FindUserByEmail), user)
}

// FindUserByOrderID mocks base method.
func (m *MockUserRepository) FindUserByOrderID(orderID string) (models.UsersProfileDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByOrderID", orderID)
	ret0, _ := ret[0].(models.UsersProfileDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByOrderID indicates an expected call of FindUserByOrderID.
func (mr *MockUserRepositoryMockRecorder) FindUserByOrderID(orderID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByOrderID", reflect.TypeOf((*MockUserRepository)(nil).FindUserByOrderID), orderID)
}

// GetAllAddresses mocks base method.
func (m *MockUserRepository) GetAllAddresses(userID int) ([]models.AddressInfoResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllAddresses", userID)
	ret0, _ := ret[0].([]models.AddressInfoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllAddresses indicates an expected call of GetAllAddresses.
func (mr *MockUserRepositoryMockRecorder) GetAllAddresses(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllAddresses", reflect.TypeOf((*MockUserRepository)(nil).GetAllAddresses), userID)
}

// GetAllPaymentOption mocks base method.
func (m *MockUserRepository) GetAllPaymentOption() ([]models.PaymentDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllPaymentOption")
	ret0, _ := ret[0].([]models.PaymentDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllPaymentOption indicates an expected call of GetAllPaymentOption.
func (mr *MockUserRepositoryMockRecorder) GetAllPaymentOption() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllPaymentOption", reflect.TypeOf((*MockUserRepository)(nil).GetAllPaymentOption))
}

// GetWalletDetails mocks base method.
func (m *MockUserRepository) GetWalletDetails(userID int) (models.Wallet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWalletDetails", userID)
	ret0, _ := ret[0].(models.Wallet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWalletDetails indicates an expected call of GetWalletDetails.
func (mr *MockUserRepositoryMockRecorder) GetWalletDetails(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWalletDetails", reflect.TypeOf((*MockUserRepository)(nil).GetWalletDetails), userID)
}

// GetWishList mocks base method.
func (m *MockUserRepository) GetWishList(userID int) ([]models.WishListResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWishList", userID)
	ret0, _ := ret[0].([]models.WishListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWishList indicates an expected call of GetWishList.
func (mr *MockUserRepositoryMockRecorder) GetWishList(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWishList", reflect.TypeOf((*MockUserRepository)(nil).GetWishList), userID)
}

// LoginHandler mocks base method.
func (m *MockUserRepository) LoginHandler(user models.UserDetails) (models.UserDetailsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoginHandler", user)
	ret0, _ := ret[0].(models.UserDetailsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoginHandler indicates an expected call of LoginHandler.
func (mr *MockUserRepositoryMockRecorder) LoginHandler(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoginHandler", reflect.TypeOf((*MockUserRepository)(nil).LoginHandler), user)
}

// ProductExistInWishList mocks base method.
func (m *MockUserRepository) ProductExistInWishList(productID, userId int) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProductExistInWishList", productID, userId)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProductExistInWishList indicates an expected call of ProductExistInWishList.
func (mr *MockUserRepositoryMockRecorder) ProductExistInWishList(productID, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProductExistInWishList", reflect.TypeOf((*MockUserRepository)(nil).ProductExistInWishList), productID, userId)
}

// RemoveFromWishList mocks base method.
func (m *MockUserRepository) RemoveFromWishList(userID, productID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveFromWishList", userID, productID)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveFromWishList indicates an expected call of RemoveFromWishList.
func (mr *MockUserRepositoryMockRecorder) RemoveFromWishList(userID, productID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveFromWishList", reflect.TypeOf((*MockUserRepository)(nil).RemoveFromWishList), userID, productID)
}

// ResetPassword mocks base method.
func (m *MockUserRepository) ResetPassword(userID int, password string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResetPassword", userID, password)
	ret0, _ := ret[0].(error)
	return ret0
}

// ResetPassword indicates an expected call of ResetPassword.
func (mr *MockUserRepositoryMockRecorder) ResetPassword(userID, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetPassword", reflect.TypeOf((*MockUserRepository)(nil).ResetPassword), userID, password)
}

// UpdateAddress mocks base method.
func (m *MockUserRepository) UpdateAddress(address models.AddressInfo, addressID, userID int) (models.AddressInfoResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAddress", address, addressID, userID)
	ret0, _ := ret[0].(models.AddressInfoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAddress indicates an expected call of UpdateAddress.
func (mr *MockUserRepositoryMockRecorder) UpdateAddress(address, addressID, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAddress", reflect.TypeOf((*MockUserRepository)(nil).UpdateAddress), address, addressID, userID)
}

// UpdateUserEmail mocks base method.
func (m *MockUserRepository) UpdateUserEmail(email string, userID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserEmail", email, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserEmail indicates an expected call of UpdateUserEmail.
func (mr *MockUserRepositoryMockRecorder) UpdateUserEmail(email, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserEmail", reflect.TypeOf((*MockUserRepository)(nil).UpdateUserEmail), email, userID)
}

// UpdateUserName mocks base method.
func (m *MockUserRepository) UpdateUserName(name string, userID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserName", name, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserName indicates an expected call of UpdateUserName.
func (mr *MockUserRepositoryMockRecorder) UpdateUserName(name, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserName", reflect.TypeOf((*MockUserRepository)(nil).UpdateUserName), name, userID)
}

// UpdateUserPassword mocks base method.
func (m *MockUserRepository) UpdateUserPassword(password string, userID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserPassword", password, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserPassword indicates an expected call of UpdateUserPassword.
func (mr *MockUserRepositoryMockRecorder) UpdateUserPassword(password, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserPassword", reflect.TypeOf((*MockUserRepository)(nil).UpdateUserPassword), password, userID)
}

// UpdateUserPhone mocks base method.
func (m *MockUserRepository) UpdateUserPhone(phone string, userID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserPhone", phone, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserPhone indicates an expected call of UpdateUserPhone.
func (mr *MockUserRepositoryMockRecorder) UpdateUserPhone(phone, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserPhone", reflect.TypeOf((*MockUserRepository)(nil).UpdateUserPhone), phone, userID)
}

// UserBlockStatus mocks base method.
func (m *MockUserRepository) UserBlockStatus(email string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserBlockStatus", email)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserBlockStatus indicates an expected call of UserBlockStatus.
func (mr *MockUserRepositoryMockRecorder) UserBlockStatus(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserBlockStatus", reflect.TypeOf((*MockUserRepository)(nil).UserBlockStatus), email)
}

// UserDetails mocks base method.
func (m *MockUserRepository) UserDetails(userID int) (models.UsersProfileDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserDetails", userID)
	ret0, _ := ret[0].(models.UsersProfileDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserDetails indicates an expected call of UserDetails.
func (mr *MockUserRepositoryMockRecorder) UserDetails(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserDetails", reflect.TypeOf((*MockUserRepository)(nil).UserDetails), userID)
}

// UserPassword mocks base method.
func (m *MockUserRepository) UserPassword(userID int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserPassword", userID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserPassword indicates an expected call of UserPassword.
func (mr *MockUserRepositoryMockRecorder) UserPassword(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserPassword", reflect.TypeOf((*MockUserRepository)(nil).UserPassword), userID)
}

// UserSignUp mocks base method.
func (m *MockUserRepository) UserSignUp(user models.UserDetails) (models.UserDetailsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserSignUp", user)
	ret0, _ := ret[0].(models.UserDetailsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserSignUp indicates an expected call of UserSignUp.
func (mr *MockUserRepositoryMockRecorder) UserSignUp(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserSignUp", reflect.TypeOf((*MockUserRepository)(nil).UserSignUp), user)
}