package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/RifaldyAldy/diamond-wallet/model"
	"github.com/RifaldyAldy/diamond-wallet/model/dto"
	"github.com/RifaldyAldy/diamond-wallet/usecase"
	"github.com/RifaldyAldy/diamond-wallet/utils/common"
	"github.com/gin-gonic/gin"
)

type TopupController struct {
	ut usecase.TopupUseCase
	uc usecase.UserUseCase
	rg *gin.RouterGroup
}

// tulis handler code kalian disini
func (t *TopupController) CreateTopupHandler(c *gin.Context) {
	var payload model.TopupModel
	var ammount dto.TopupRequest
	c.ShouldBind(&ammount)
	payload.TransactionDetails.GrossAmt = int64(ammount.Ammount)
	claims, exists := c.Get("claims")
	if !exists {
		common.SendErrorResponse(c, http.StatusBadRequest, "Sepertinya login anda tidak valid")
		return
	}
	payload.User.Id = claims.(*common.JwtClaim).DataClaims.Id
	fmt.Println(payload.User.Id)
	payload.User, _ = t.uc.FindById(payload.User.Id)
	res, err := t.ut.CreateTopup(payload)
	if err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendCreateResponse(c, "SUCCESS", res)
}

func (t *TopupController) ResponseTopupHandler(c *gin.Context) {
	var payload dto.ResponsePayment
	payload.OrderId = c.Query("order_id")
	payload.StatusCode, _ = strconv.Atoi(c.Query("status_code"))
	payload.TransactionStatus = c.Query("transaction_status")

	res, err := t.ut.PaymentUpdate(payload)
	if err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	common.SendSingleResponse(c, "SUCCESS", res)
}

func (t *TopupController) Route() {
	rg := t.rg.Group("/topup")
	{
		//tulis route disini
		rg.POST("/", common.JWTAuth("user"), t.CreateTopupHandler)
		rg.GET("/response", t.ResponseTopupHandler)
	}
}

func NewTopupController(ut usecase.TopupUseCase, uc usecase.UserUseCase, rg *gin.RouterGroup) *TopupController {
	return &TopupController{ut: ut, uc: uc, rg: rg}
}
