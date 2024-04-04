package repomock

import (
	"github.com/RifaldyAldy/diamond-wallet/model"
	"github.com/RifaldyAldy/diamond-wallet/model/dto"
	"github.com/stretchr/testify/mock"
)

type UserRepoMock struct {
	mock.Mock
}

func (u *UserRepoMock) Get(id string) (model.User, error) {
	args := u.Called(id)
	return args.Get(0).(model.User), args.Error(1)
}

func (u *UserRepoMock) Create(payload dto.UserRequestDto) (model.User, error) {
	args := u.Called(payload)
	return args.Get(0).(model.User), args.Error(1)
}

func (u *UserRepoMock) GetBalance(user_id string) (model.UserSaldo, error) {
	args := u.Called(user_id)
	return args.Get(0).(model.UserSaldo), args.Error(1)
}

func (u *UserRepoMock) GetByUsername(username string) (model.User, error) {
	args := u.Called(username)
	return args.Get(0).(model.User), args.Error(1)
}

func (u *UserRepoMock) Update(id string, payload model.User) (model.User, error) {
	args := u.Called(id, payload)
	return args.Get(0).(model.User), args.Error(1)
}

func (u *UserRepoMock) Verify(payload dto.VerifyUser) (dto.VerifyUser, error) {
	args := u.Called(payload)
	return args.Get(0).(dto.VerifyUser), args.Error(1)
}

func (u *UserRepoMock) UpdatePin(payload dto.UpdatePinRequest) (dto.UpdatePinResponse, error) {
	args := u.Called(payload)
	return args.Get(0).(dto.UpdatePinResponse), args.Error(1)
}

func (u *UserRepoMock) GetInfoUser(Info string, limit, offset int) ([]model.User, error) {
	args := u.Called(Info, limit, offset)
	return args.Get(0).([]model.User), args.Error(1)
}

func (u *UserRepoMock) GetRekening(id string) (model.Rekening, error) {
	args := u.Called(id)
	return args.Get(0).(model.Rekening), args.Error(1)
}

func (u *UserRepoMock) CreateRekening(payload model.Rekening) (model.Rekening, error) {
	args := u.Called(payload)
	return args.Get(0).(model.Rekening), args.Error(1)
}
