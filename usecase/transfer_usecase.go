package usecase

import (
	"github.com/RifaldyAldy/diamond-wallet/model"
	"github.com/RifaldyAldy/diamond-wallet/model/dto"
	"github.com/RifaldyAldy/diamond-wallet/repository"
)

type TransferUseCase interface {
	TransferRequest(payload dto.TransferRequest, send, receive model.User) (model.Transfer, error)
	GetSend(id string, page int) ([]model.Transfer, error)
	GetReceive(id string, page int) ([]model.Transfer, error)
	Withdraw(payload model.Withdraw, saldo model.UserSaldo) (model.Withdraw, error)
	GetAllWithDraw(id string, page int) ([]model.Withdraw, error)
}

type transferUseCase struct {
	repo repository.TransferRepository
}

// tulis code kalian disini

func (t *transferUseCase) TransferRequest(payload dto.TransferRequest, send, receive model.User) (model.Transfer, error) {
	response, err := t.repo.Create(payload, send, receive)
	if err != nil {
		return model.Transfer{}, err
	}

	return response, nil
}

func (t *transferUseCase) GetSend(id string, page int) ([]model.Transfer, error) {
	datas, err := t.repo.GetSend(id, page)
	if err != nil {
		return []model.Transfer{}, err
	}

	return datas, nil
}

func (t *transferUseCase) GetReceive(id string, page int) ([]model.Transfer, error) {
	datas, err := t.repo.GetReceive(id, page)
	if err != nil {
		return []model.Transfer{}, err
	}

	return datas, nil
}

func (t *transferUseCase) Withdraw(payload model.Withdraw, saldo model.UserSaldo) (model.Withdraw, error) {
	res, err := t.repo.CreateWithdraw(payload, saldo)
	if err != nil {
		return model.Withdraw{}, err
	}

	return res, nil
}

func (t *transferUseCase) GetAllWithDraw(id string, page int) ([]model.Withdraw, error) {

	datas, err := t.repo.GetWithdraw(id, page)
	if err != nil {
		return []model.Withdraw{}, err
	}

	return datas, nil
}

func NewTransferUseCase(repo repository.TransferRepository) TransferUseCase {
	return &transferUseCase{repo: repo}
}
