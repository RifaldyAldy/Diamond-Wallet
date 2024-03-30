package manager

import (
	"github.com/RifaldyAldy/diamond-wallet/repository"
)

type RepoManager interface {
	TransferRepo() repository.TransferRepository
	TopupRepo() repository.TopupRepository
	UserRepo() repository.UserRepository
	AdminRepo() repository.AdminRepository
}

type repoManager struct {
	infra InfraManager
}

func (r *repoManager) TransferRepo() repository.TransferRepository {
	return repository.NewTransferRepository(r.infra.Conn())
}

func (r *repoManager) TopupRepo() repository.TopupRepository {
	return repository.NewTopUpRepository(r.infra.Conn())
}

func (r *repoManager) UserRepo() repository.UserRepository {
	return repository.NewUserRepository(r.infra.Conn())
}

func (r *repoManager) AdminRepo() repository.AdminRepository {
	return repository.NewAdminRepository(r.infra.Conn())
}

func NewRepoManager(infra InfraManager) RepoManager {
	return &repoManager{infra: infra}
}
