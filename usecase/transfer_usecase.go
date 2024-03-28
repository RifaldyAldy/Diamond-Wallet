package usecase

import "github.com/RifaldyAldy/diamond-wallet/repository"

type TransferUseCase interface {
}

type transferUseCase struct {
	repo repository.TransferRepository
}

// tulis code kalian disini

func NewTransferUseCase(repo repository.TransferRepository) TransferUseCase {
	return &transferUseCase{repo: repo}
}
