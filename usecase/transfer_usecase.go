package usecase

import (
	"github.com/RifaldyAldy/diamond-wallet/model"
	"github.com/RifaldyAldy/diamond-wallet/model/dto"
	"github.com/RifaldyAldy/diamond-wallet/repository"
)

type TransferUseCase interface {
	TransferRequest(payload dto.TransferRequest, send, receive model.User) (model.Transfer, error)
}

type transferUseCase struct {
	repo   repository.TransferRepository
	userUC userUseCase
}

// tulis code kalian disini

func (t *transferUseCase) TransferRequest(payload dto.TransferRequest, send, receive model.User) (model.Transfer, error) {
	response, err := t.repo.Create(payload, send, receive)
	if err != nil {
		return model.Transfer{}, err
	}

	return response, nil
}

func NewTransferUseCase(repo repository.TransferRepository) TransferUseCase {
	return &transferUseCase{repo: repo}
}
