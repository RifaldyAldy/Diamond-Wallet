package usecase

import (
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
	return a.userRepository.GetByID(userID)
}

func NewAdminUseCase(repo repository.AdminRepository) AdminUseCase {
	return &adminUseCase{repo: repo}
}
