package usecase

import (
	"github.com/RifaldyAldy/diamond-wallet/model"
	"github.com/RifaldyAldy/diamond-wallet/repository"
)

type AdminUseCase interface {
	RegisterAdmin(payload model.Admin) (model.Admin, error)
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

func NewAdminUseCase(repo repository.AdminRepository) AdminUseCase {
	return &adminUseCase{repo: repo}
}
