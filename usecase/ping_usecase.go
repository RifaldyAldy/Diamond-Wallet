package usecase

import (
	"github.com/RifaldyAldy/diamond-wallet/repository"
)

type PingUseCase interface {
	Ping() error
}

type pingUseCase struct {
	repo repository.PingRepository
}

func (p *pingUseCase) Ping() error {
	err := p.repo.Ping()
	if err != nil {
		return err
	}
	return nil
}

func NewPingUseCase(repo repository.PingRepository) PingUseCase {
	return &pingUseCase{repo: repo}
}
