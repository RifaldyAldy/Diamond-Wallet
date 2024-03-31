package usecase

import (
	"errors"
	"time"

	"github.com/RifaldyAldy/diamond-wallet/model"
	"github.com/RifaldyAldy/diamond-wallet/model/dto"
	"github.com/RifaldyAldy/diamond-wallet/repository"
	"github.com/RifaldyAldy/diamond-wallet/utils/common"
	encryption "github.com/RifaldyAldy/diamond-wallet/utils/encription"
)

type AdminUseCase interface {
	RegisterAdmin(payload model.Admin) (model.Admin, error)
	LoginAdmin(payload dto.LoginRequestDto) (dto.LoginResponseDto, error)
}

type adminUseCase struct {
	repo repository.AdminRepository
}

func (a *adminUseCase) RegisterAdmin(payload model.Admin) (model.Admin, error) {
	response, err := a.repo.Register(payload)
	if err != nil {
		return model.Admin{}, err
	}
	return response, nil
}

func (a *adminUseCase) LoginAdmin(payload dto.LoginRequestDto) (dto.LoginResponseDto, error) {
	var claims dto.LoginResponseDto
	response, err := a.repo.Get(payload)
	if err != nil {
		return dto.LoginResponseDto{}, err
	}
	isValid := encryption.CheckPasswordHash(payload.Pass, response.Password)
	if !isValid {
		return dto.LoginResponseDto{}, errors.New("password salah")
	}

	loginExpDuration := time.Duration(10) * time.Minute
	expiredAt := time.Now().Add(loginExpDuration).Unix()
	// TODO: tempel generate token jwt
	accessToken, err := common.GenerateTokenJwt(response.Id, response.Name, response.Role, expiredAt)
	if err != nil {
		return dto.LoginResponseDto{}, err
	}

	claims.AccessToken = accessToken
	claims.UserId = response.Id
	return claims, nil
}

func NewAdminUseCase(repo repository.AdminRepository) AdminUseCase {
	return &adminUseCase{repo: repo}
}
