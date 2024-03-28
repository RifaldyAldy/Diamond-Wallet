package manager

import (
	"github.com/RifaldyAldy/diamond-wallet/repository"
)

type RepoManager interface {
	PingRepo() repository.PingRepository
}

type repoManager struct {
	infra InfraManager
}

func (r *repoManager) PingRepo() repository.PingRepository {
	return repository.NewPingRepository(r.infra.Conn())
}

func NewRepoManager(infra InfraManager) RepoManager {
	return &repoManager{infra: infra}
}
