package usecase

import (
	"errors"
	"time"

	"github.com/RifaldyAldy/diamond-wallet/model/dto"
	"github.com/RifaldyAldy/diamond-wallet/repository"
	"github.com/RifaldyAldy/diamond-wallet/utils/common"
	encryption "github.com/RifaldyAldy/diamond-wallet/utils/encription"
)

type UserUseCase interface {
	LoginUser(in dto.LoginRequestDto) (dto.LoginResponseDto, error)
}

type userUseCase struct {
	repo repository.UserRepository
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
	accessToken, err := common.GenerateTokenJwt(userData, expiredAt)
	if err != nil {
		return dto.LoginResponseDto{}, err
	}
	return dto.LoginResponseDto{
		AccessToken: accessToken,
		UserId:      userData.Id,
	}, nil
}

func NewUserUseCase(repo repository.UserRepository) UserUseCase {
	return &userUseCase{repo: repo}
}
