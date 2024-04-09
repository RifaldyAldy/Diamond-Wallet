package usecase

import (
	"testing"

	repomock "github.com/RifaldyAldy/diamond-wallet/mock/repo_mock"
	"github.com/RifaldyAldy/diamond-wallet/model"
	"github.com/RifaldyAldy/diamond-wallet/model/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type UserUseCaseTestSuite struct {
	suite.Suite
	urm *repomock.UserRepoMock
	uu  UserUseCase
}

func (suite *UserUseCaseTestSuite) SetupTest() {
	suite.urm = new(repomock.UserRepoMock)
	suite.uu = NewUserUseCase(suite.urm)
}

func TestUserUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(UserUseCaseTestSuite))
}

func (suite *UserUseCaseTestSuite) TestGetUser_Success() {
	user := model.User{}
	suite.urm.On("Get", user.Id).Return(user, nil)

	actual, err := suite.uu.FindById(user.Id)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), user.Id, actual.Id)
}

func (suite *UserUseCaseTestSuite) TestCreateUser_Success() {
	payloadMock := dto.UserRequestDto{
		Id:          "1",
		Name:        "user 1",
		Username:    "user",
		Password:    "password",
		Role:        "user",
		Email:       "user@gmail.com",
		PhoneNumber: "021341241",
	}
	expected := model.User{Id: "1"}
	hashPasswordMock := payloadMock
	suite.urm.On("hashPassword", hashPasswordMock).Return(hashPasswordMock, nil)
	suite.urm.On("Create", mock.Anything).Return(expected, nil)

	actual, err := suite.uu.CreateUser(payloadMock)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expected.Id, actual.Id)
}

func (suite *UserUseCaseTestSuite) TestUpdateUser_Success() {
	id := "1"
	updateUserMock := dto.UserRequestDto{Name: "user"}
	updatedUserMock := model.User{Id: id, Name: "user"}

	suite.urm.On("Get", id).Return(updatedUserMock, nil)

	suite.urm.On("Update", id, mock.Anything).Return(updatedUserMock, nil)

	actual, err := suite.uu.UpdateUser(id, updateUserMock)

	assert.Nil(suite.T(), err)

	assert.Equal(suite.T(), updatedUserMock.Name, actual.Name)
	assert.Equal(suite.T(), updatedUserMock.Id, actual.Id)
}

func (suite *UserUseCaseTestSuite) TestGetBalance_Success() {
	var id string
	expected := model.UserSaldo{}
	suite.urm.On("GetBalance", id).Return(expected, nil)

	actual, err := suite.uu.GetBalanceCase(id)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expected.User, actual.User)
}

//* untuk Login user harus buat mock dari generate token jwt
// func (suite *UserUseCaseTestSuite) TestLoginUser_Success() {
// inMock :=dto.LoginRequestDto{}
// suite.urm.On("GetByUserUsername",inMock).Return(inMock,nil)

// suite.
// }

func (suite *UserUseCaseTestSuite) TestVerifyUser_Success() {
	payloadMock := dto.VerifyUser{}
	suite.urm.On("Verify", payloadMock).Return(payloadMock, nil)
	actual, err := suite.uu.VerifyUser(payloadMock)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), payloadMock, actual)
}

func (suite *UserUseCaseTestSuite) TestUpdatePinUser_Success() {
	payloadMock := dto.UpdatePinRequest{NewPin: "123456"}
	responseMock := dto.UpdatePinResponse{Pin: "123456"}

	suite.urm.On("UpdatePin", payloadMock.OldPin).Return(responseMock.Pin, nil)

	actual, err := suite.uu.UpdatePinUser(payloadMock)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), payloadMock.NewPin, actual.Pin)
}
func (suite *UserUseCaseTestSuite) TestFindRekening_Success() {
	var id string
	rekeningMock := model.Rekening{}
	suite.urm.On("GetRekening", id).Return(rekeningMock, nil)

	actual, err := suite.uu.FindRekening(id)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), rekeningMock, actual)
}

func (suite *UserUseCaseTestSuite) TestCreateRekening_Success() {
	payloadMock := model.Rekening{}
	suite.urm.On("CreateRekening", payloadMock).Return(payloadMock, nil)

	actual, err := suite.uu.CreateRekening(payloadMock)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), payloadMock, actual)
}
