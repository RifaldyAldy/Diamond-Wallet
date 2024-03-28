package controller

import (
	"github.com/RifaldyAldy/diamond-wallet/usecase"
	"github.com/gin-gonic/gin"
)

type TopupController struct {
	uc usecase.TopupUseCase
	rg *gin.RouterGroup
}

// tulis handler code kalian disini

func (t *TopupController) Route() {
	rg := t.rg.Group("/topup")
	{
		//tulis route disini
	}
}

func NewTopupController(uc usecase.TopupUseCase, rg *gin.RouterGroup) *TopupController {
	return &TopupController{uc: uc, rg: rg}
}
