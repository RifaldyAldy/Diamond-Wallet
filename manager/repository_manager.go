package manager

import (
	"github.com/RifaldyAldy/diamond-wallet/repository"
)

type RepoManager interface {
	UserRepo() repository.UserRepository
}

type repoManager struct {
	infra InfraManager
}

func (r *repoManager) UserRepo() repository.UserRepository {
	return repository.NewUserRepository(r.infra.Conn())
}

func NewRepoManager(infra InfraManager) RepoManager {
	return &repoManager{infra: infra}
}
