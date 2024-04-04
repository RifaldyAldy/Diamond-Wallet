package repository

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/RifaldyAldy/diamond-wallet/model/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UserRepositoryTestSuite struct {
	suite.Suite
	mockDB  *sql.DB
	mockSql sqlmock.Sqlmock
	repo    UserRepository
}

func (suite *UserRepositoryTestSuite) SetupTest() {
	db, mock, err := sqlmock.New()
	assert.NoError(suite.T(), err)
	suite.mockDB = db
	suite.mockSql = mock
	suite.repo = NewUserRepository(suite.mockDB)
}

func TestUserRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}

func (suite *UserRepositoryTestSuite) TestCreateUser_Success() {
	mockUser := dto.UserRequestDto{
		Id:          "1",
		Name:        "Chigan",
		Username:    "chigan",
		Password:    "user",
		Role:        "user",
		Email:       "chigan@example.com",
		PhoneNumber: "08123456789",
	}

	suite.mockSql.ExpectBegin()
	rows := sqlmock.NewRows([]string{"id", "name", "username", "password", "role", "email", "phone_number"}).
		AddRow(mockUser.Id, mockUser.Name, mockUser.Username, mockUser.Password, mockUser.Role, mockUser.Email, mockUser.PhoneNumber)
	suite.mockSql.ExpectQuery("INSERT INTO mst_user").WillReturnRows(rows)

	suite.mockSql.ExpectCommit()

	actual, err := suite.repo.Create(mockUser)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), mockUser.Id, actual.Id)
}
