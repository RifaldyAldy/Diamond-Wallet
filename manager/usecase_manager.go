package manager

import (
	"github.com/RifaldyAldy/diamond-wallet/usecase"
)

type UseCaseManager interface {
	TransferUseCase() usecase.TransferUseCase
	TopupUseCase() usecase.TopupUseCase
	UserUseCase() usecase.UserUseCase
}

type useCaseManager struct {
	repo RepoManager
}

func (u *useCaseManager) TransferUseCase() usecase.TransferUseCase {
	return usecase.NewTransferUseCase(u.repo.TransferRepo())
}

func (u *useCaseManager) TopupUseCase() usecase.TopupUseCase {
	return usecase.NewTopupUseCase(u.repo.TopupRepo())
}
func (u *useCaseManager) UserUseCase() usecase.UserUseCase {
	return usecase.NewUserUseCase(u.repo.UserRepo())
}

func NewUseCaseManager(repo RepoManager) UseCaseManager {
	return &useCaseManager{repo: repo}
}
