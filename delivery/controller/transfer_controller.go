package controller

import (
	"github.com/RifaldyAldy/diamond-wallet/usecase"
	"github.com/gin-gonic/gin"
)

type TransferController struct {
	uc usecase.TransferUseCase
	rg *gin.RouterGroup
}

// tulis handler code kalian disini

func (t *TransferController) Route() {
	// rg := t.rg.Group("/transfer")
	{
		//tulis route disini
	}
}

func NewTransferController(uc usecase.TransferUseCase, rg *gin.RouterGroup) *TransferController {
	return &TransferController{uc: uc, rg: rg}
}
