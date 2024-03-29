package usecase

import (
	"github.com/RifaldyAldy/diamond-wallet/model"
	"github.com/RifaldyAldy/diamond-wallet/model/dto"
	"github.com/RifaldyAldy/diamond-wallet/repository"
	"github.com/RifaldyAldy/diamond-wallet/utils/common"
)

type TopupUseCase interface {
	CreateTopup(payload model.TopupModel) (common.ResponseMidtrans, error)
	FindById(orderId string) (model.TableTopupPayment, error)
	PaymentUpdate(payload dto.ResponsePayment) (dto.ResponsePayment, error)
}

type topupUseCase struct {
	repo repository.TopupRepository
}

func (t *topupUseCase) CreateTopup(payload model.TopupModel) (common.ResponseMidtrans, error) {

	midtransResponse, err := t.repo.Create(payload)
	if err != nil {
		return common.ResponseMidtrans{}, err
	}

	return midtransResponse, nil
}

func (t *topupUseCase) FindById(orderId string) (model.TableTopupPayment, error) {
	tabel, err := t.repo.Getbyid(orderId)
	if err != nil {
		return model.TableTopupPayment{}, err
	}

	return tabel, err
}

func (t *topupUseCase) PaymentUpdate(payload dto.ResponsePayment) (dto.ResponsePayment, error) {
	payload, err := t.repo.Payment(payload)
	if err != nil {
		return dto.ResponsePayment{}, err
	}

	return payload, nil
}

func NewTopupUseCase(repo repository.TopupRepository) TopupUseCase {
	return &topupUseCase{repo: repo}
}
