package usecasemock

import (
	"github.com/RifaldyAldy/diamond-wallet/model"
	"github.com/RifaldyAldy/diamond-wallet/model/dto"
	"github.com/stretchr/testify/mock"
)

type UserUseCaseMock struct {
	mock.Mock
}

func (u *UserUseCaseMock) CreateUser(payload dto.UserRequestDto) (model.User, error) {
	args := u.Called(payload)
	return args.Get(0).(model.User), args.Error(1)
}

func (u *UserUseCaseMock) LoginUser(in dto.LoginRequestDto) (dto.LoginResponseDto, error) {
	args := u.Called(in)
	return args.Get(0).(dto.LoginResponseDto), args.Error(1)
}

func (u *UserUseCaseMock) FindById(id string) (model.User, error) {
	args := u.Called(id)
	return args.Get(0).(model.User), args.Error(1)
}

func (u *UserUseCaseMock) GetBalanceCase(id string) (model.UserSaldo, error) {
	args := u.Called(id)
	return args.Get(0).(model.UserSaldo), args.Error(1)
}

func (u *UserUseCaseMock) UpdateUser(id string, payload dto.UserRequestDto) (model.User, error) {
	args := u.Called(id, payload)
	return args.Get(0).(model.User), args.Error(1)
}

func (u *UserUseCaseMock) VerifyUser(payload dto.VerifyUser) (dto.VerifyUser, error) {
	args := u.Called(payload)
	return args.Get(0).(dto.VerifyUser), args.Error(1)
}

func (u *UserUseCaseMock) UpdatePinUser(payload dto.UpdatePinRequest) (dto.UpdatePinResponse, error) {
	args := u.Called(payload)
	return args.Get(0).(dto.UpdatePinResponse), args.Error(1)
}

func (u *UserUseCaseMock) FindRekening(id string) (model.Rekening, error) {
	args := u.Called(id)
	return args.Get(0).(model.Rekening), args.Error(1)
}

func (u *UserUseCaseMock) CreateRekening(payload model.Rekening) (model.Rekening, error) {
	args := u.Called(payload)
	return args.Get(0).(model.Rekening), args.Error(1)
}
