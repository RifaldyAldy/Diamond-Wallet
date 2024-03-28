package manager

import (
	"github.com/RifaldyAldy/diamond-wallet/usecase"
)

type UseCaseManager interface {
	PingUseCase() usecase.PingUseCase
}

type useCaseManager struct {
	repo RepoManager
}

func (u *useCaseManager) PingUseCase() usecase.PingUseCase {
	return usecase.NewPingUseCase(u.repo.PingRepo())
}

func NewUseCaseManager(repo RepoManager) UseCaseManager {
	return &useCaseManager{repo: repo}
}
