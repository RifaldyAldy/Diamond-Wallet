package usecase

import (
	"errors"

	"github.com/RifaldyAldy/diamond-wallet/model"
	"github.com/RifaldyAldy/diamond-wallet/repository"
)

type AdminUseCase interface {
	RegisterAdmin(payload model.Admin) (model.Admin, error)
	GetUserInfo(userID string) (model.User, error)
}

type adminUseCase struct {
	repo           repository.AdminRepository
	userRepository repository.UserRepository
}

func (a *adminUseCase) RegisterAdmin(payload model.Admin) (model.Admin, error) {
	response, err := a.repo.Register(payload)
	if err != nil {
		return model.Admin{}, err
	}
	return response, nil
}

func (a *adminUseCase) GetUserInfo(userID string) (model.User, error) {

	info := "u.id= '" + userID + "'"
	limit := 1
	offset := 0

	users, err := a.userRepository.GetInfoUser(info, limit, offset)
	if err != nil {
		return model.User{}, err
	}
	if len(users) == 0 {
		return model.User{}, errors.New("user tidak ada")
	}
	return users[0], nil

}

func NewAdminUseCase(repo repository.AdminRepository) AdminUseCase {
	return &adminUseCase{repo: repo}
}
