package usecase

import "github.com/RifaldyAldy/diamond-wallet/repository"

type TopupUseCase interface {
}

type topupUseCase struct {
	repo repository.TopupRepository
}

func NewTopupUseCase(repo repository.TopupRepository) TopupUseCase {
	return &topupUseCase{repo: repo}
}
