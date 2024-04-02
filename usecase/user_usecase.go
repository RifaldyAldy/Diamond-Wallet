package usecase

import (
	"errors"
	"fmt"
	"time"

	"github.com/RifaldyAldy/diamond-wallet/model"
	"github.com/RifaldyAldy/diamond-wallet/model/dto"
	"github.com/RifaldyAldy/diamond-wallet/repository"
	"github.com/RifaldyAldy/diamond-wallet/utils/common"
	encryption "github.com/RifaldyAldy/diamond-wallet/utils/encription"
)

type UserUseCase interface {
	CreateUser(payload dto.UserRequestDto) (model.User, error)
	LoginUser(in dto.LoginRequestDto) (dto.LoginResponseDto, error)
	FindById(id string) (model.User, error)
	GetBalanceCase(id string) (model.UserSaldo, error)
	UpdateUser(id string, payload dto.UserRequestDto) (model.User, error)
	VerifyUser(payload dto.VerifyUser) (dto.VerifyUser, error)
	UpdatePinUser(payload dto.UpdatePinUser) (dto.UpdatePinUser, error)
}

type userUseCase struct {
	repo repository.UserRepository
}

func (u *userUseCase) FindById(id string) (model.User, error) {
	user, err := u.repo.Get(id)
	if err != nil {
		return model.User{}, fmt.Errorf("user with ID %s not found", id)
	}

	return user, nil
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

func (u *userUseCase) LoginUser(in dto.LoginRequestDto) (dto.LoginResponseDto, error) {
	userData, err := u.repo.GetByUsername(in.Username)
	if err != nil {
		return dto.LoginResponseDto{}, err
	}
	isValid := encryption.CheckPasswordHash(in.Pass, userData.Password)
	if !isValid {
		return dto.LoginResponseDto{}, errors.New("1")
	}

	loginExpDuration := time.Duration(10) * time.Minute
	expiredAt := time.Now().Add(loginExpDuration).Unix()
	// TODO: tempel generate token jwt
	accessToken, err := common.GenerateTokenJwt(userData.Id, userData.Name, userData.Role, expiredAt)
	if err != nil {
		return dto.LoginResponseDto{}, err
	}
	return dto.LoginResponseDto{
		AccessToken: accessToken,
		UserId:      userData.Id,
	}, nil
}

func (u *userUseCase) GetBalanceCase(id string) (model.UserSaldo, error) {
	response, err := u.repo.GetBalance(id)
	if err != nil {
		return model.UserSaldo{}, err
	}

	return response, nil
}

func (u *userUseCase) UpdateUser(id string, payload dto.UserRequestDto) (model.User, error) {
	updatedUser := model.User{
		Id:          id,
		Name:        payload.Name,
		Username:    payload.Username,
		Role:        payload.Role,
		Email:       payload.Email,
		PhoneNumber: payload.PhoneNumber,
	}
	user, err := u.repo.Update(id, updatedUser)
	if err != nil {
		return model.User{}, fmt.Errorf("failed to update user : %v", err.Error())
	}
	return user, nil
}

func (u *userUseCase) VerifyUser(payload dto.VerifyUser) (dto.VerifyUser, error) {
	response, err := u.repo.Verify(payload)
	if err != nil {
		return dto.VerifyUser{}, err
	}
	return response, nil
}

func (u *userUseCase) UpdatePinUser(payload dto.UpdatePinUser) (dto.UpdatePinUser, error) {
	response, err := u.repo.UpdatePin(payload)
	if err != nil {
		return dto.UpdatePinUser{}, err
	}
	fmt.Println("response usecase", response)
	return response, nil
}

func NewUserUseCase(repo repository.UserRepository) UserUseCase {
	return &userUseCase{repo: repo}
}
