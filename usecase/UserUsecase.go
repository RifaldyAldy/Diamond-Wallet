package usecase

import (
	"fmt"

	"github.com/RifaldyAldy/diamond-wallet/model"
	"github.com/RifaldyAldy/diamond-wallet/model/dto"
	"github.com/RifaldyAldy/diamond-wallet/repository"
	"github.com/RifaldyAldy/diamond-wallet/utils/encryption"
)

type UserUseCase interface {
	CreateUser(payload dto.UserRequestDto) (model.User, error)
}

type userUseCase struct {
	repo repository.UserRepository
}

func (u *userUseCase) CreateUser(payload dto.UserRequestDto) (model.User, error) {
	hashPassword, err := encryption.HashPassword(payload.Password)
	if err != nil {
		return model.User{}, err
	}
	newUser := model.User{
		Id:          payload.Id,
		Name:        payload.Name,
		Username:    payload.Username,
		Password:    hashPassword,
		Role:        payload.Role,
		Email:       payload.Email,
		PhoneNumber: payload.PhoneNumber,
	}
	user, err := u.repo.Create(newUser)
	if err != nil {
		return model.User{}, fmt.Errorf("failed to create user: %v", err.Error())
	}
	return user, nil
}

func NewUserUseCase(repo repository.UserRepository) UserUseCase {
	return &userUseCase{repo: repo}
}
