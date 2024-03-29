package usecase

import (
	"github.com/RifaldyAldy/diamond-wallet/model"
	"github.com/RifaldyAldy/diamond-wallet/repository"
)

type TopupUseCase interface {
	CreateTopup(payload model.TopupModel) (any, error)
}

type topupUseCase struct {
	repo repository.TopupRepository
}

func (t *topupUseCase) CreateTopup(payload model.TopupModel) (any, error) {

	midtransResponse, err := t.repo.Create(payload)
	if err != nil {
		return "", err
	}

	return midtransResponse, nil
}

func NewTopupUseCase(repo repository.TopupRepository) TopupUseCase {
	return &topupUseCase{repo: repo}
}
