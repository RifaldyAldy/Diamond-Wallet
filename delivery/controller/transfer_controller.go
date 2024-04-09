package controller

import (
	"net/http"
	"strconv"

	"github.com/RifaldyAldy/diamond-wallet/model"
	"github.com/RifaldyAldy/diamond-wallet/model/dto"
	"github.com/RifaldyAldy/diamond-wallet/usecase"
	"github.com/RifaldyAldy/diamond-wallet/utils/common"
	"github.com/gin-gonic/gin"
)

type TransferController struct {
	ut usecase.TransferUseCase
	uc usecase.UserUseCase
	rg *gin.RouterGroup
}

// CreateTransfer godoc
// @Tags Transfer
// @Summary Create request Transfer
// @Description User may transfer to other user
// @Accept json
// @Produce json
// @Security 	ApiKeyAuth
// @Param 		Authorization header string true "Bearer"
// @Param Data body dto.TransferRequestSwag true "Payload data"
// @Success 201 {object} model.Transfer
// @Router /transfer [post]
func (t *TransferController) TransferHandler(c *gin.Context) {
	payload := dto.TransferRequest{}
	c.ShouldBind(&payload)
	claims, exists := c.Get("claims")
	if !exists {
		common.SendErrorResponse(c, http.StatusBadRequest, "Sepertinya login anda tidak valid")
		return
	}
	claimsJwt := claims.(*common.JwtClaim)
	payload.UserId = claimsJwt.DataClaims.Id
	send, err := t.uc.FindById(payload.UserId)
	if err != nil {
		if err.Error() == "1" {
			common.SendErrorResponse(c, http.StatusBadRequest, "Anda harus memverifikasi akun terlebih dahulu")
			return
		}
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	receive, err := t.uc.FindById(payload.TujuanTransfer)
	if err != nil {
		if err.Error() == "1" {
			common.SendErrorResponse(c, http.StatusBadRequest, "Penerima harus memverifikasi akun terlebih dahulu")
			return
		}
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	sendBalance, _ := t.uc.GetBalanceCase(send.Id)
	if sendBalance.Pin != payload.Pin { // cek apakah pin di input benar
		common.SendErrorResponse(c, http.StatusBadRequest, "pin salah!")
		return
	}
	receiveBalance, err := t.uc.GetBalanceCase(receive.Id)
	if err != nil {
		if err.Error() == "1" {
			common.SendErrorResponse(c, http.StatusBadRequest, "Penerima harus memverifikasi akun terlebih dahulu")
			return
		}
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	send.Saldo = sendBalance.Saldo
	receive.Saldo = receiveBalance.Saldo
	response, err := t.ut.TransferRequest(payload, send, receive)
	if err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	common.SendSingleResponse(c, "SUCCESS", response)
}

// GetSendTransfer godoc
// @Tags Transfer
// @Summary User can get their histories send transfer @3datas/page
// @Description User can get their histories send transfer @3datas/page
// @Accept json
// @Produce json
// @Security 	ApiKeyAuth
// @Param 		Authorization header string true "Bearer"
// @Success 200 {object} []model.Transfer
// @Router /transfer/history/send [get]
func (t *TransferController) GetSendTransferHandler(c *gin.Context) {
	var id string
	var page int
	page, _ = strconv.Atoi(c.Query("page"))

	if page == 0 {
		page = 1
	}
	claims, exists := c.Get("claims")
	if !exists {
		common.SendErrorResponse(c, http.StatusBadRequest, "Sepertinya login anda tidak valid")
		return
	}
	id = claims.(*common.JwtClaim).DataClaims.Id

	datas, err := t.ut.GetSend(id, page)
	if err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	common.SendSingleResponse(c, "SUCCESS", datas)
}

// GetReceiveTransfer godoc
// @Tags Transfer
// @Summary User can get their histories receive transfer @3datas/page
// @Description User can get their histories receive transfer @3datas/page
// @Accept json
// @Produce json
// @Security 	ApiKeyAuth
// @Param 		Authorization header string true "Bearer"
// @Success 200 {object} []model.Transfer
// @Router /transfer/history/receive [get]
func (t *TransferController) GetReceiveTransferHandler(c *gin.Context) {
	var id string
	var page int
	page, _ = strconv.Atoi(c.Query("page"))

	if page == 0 {
		page = 1
	}
	claims, exists := c.Get("claims")
	if !exists {
		common.SendErrorResponse(c, http.StatusBadRequest, "Sepertinya login anda tidak valid")
		return
	}
	id = claims.(*common.JwtClaim).DataClaims.Id

	datas, err := t.ut.GetReceive(id, page)
	if err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	common.SendSingleResponse(c, "SUCCESS", datas)
}

// GetSendTransferAdmin godoc
// @Tags Admin
// @Summary Admin can get their histories send transfer @3datas/page
// @Description Admin can get their histories send transfer @3datas/page
// @Accept json
// @Produce json
// @Security 	ApiKeyAuth
// @Param 		Authorization header string true "Bearer"
// @Param		id path string true "User ID"
// @Success 200 {object} []model.Transfer
// @Router /transfer/history/admin/send/{id} [get]
func (t *TransferController) AdminGetSendTransferHandler(c *gin.Context) {
	var id string
	var page int
	page, _ = strconv.Atoi(c.Query("page"))

	if page == 0 {
		page = 1
	}

	id = c.Param("id")

	datas, err := t.ut.GetSend(id, page)
	if err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	common.SendSingleResponse(c, "SUCCESS", datas)
}

// GetReceiveTransferAdmin godoc
// @Tags Admin
// @Summary Admin can get user histories receive transfer @3datas/page
// @Description Admin can get user histories receive transfer @3datas/page
// @Accept json
// @Produce json
// @Security 	ApiKeyAuth
// @Param 		Authorization header string true "Bearer"
// @Param		id path string true "User ID"
// @Success 200 {object} []model.Transfer
// @Router /transfer/history/admin/receive/{id} [get]
func (t *TransferController) AdminGetReceiveTransferHandler(c *gin.Context) {
	var id string
	var page int
	page, _ = strconv.Atoi(c.Query("page"))

	if page == 0 {
		page = 1
	}
	id = c.Param("id")

	datas, err := t.ut.GetReceive(id, page)
	if err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	common.SendSingleResponse(c, "SUCCESS", datas)
}

// CreateRequestWithdraw godoc
// @Tags Transfer
// @Summary User can withdraw their balance to rekening bank
// @Description User can withdraw their balance to rekening bank
// @Accept json
// @Produce json
// @Security 	ApiKeyAuth
// @Param 		Authorization header string true "Bearer"
// @Param		data body dto.WithdrawDto true "Withdraw ammount"
// @Success 201 {object} model.Withdraw
// @Router /transfer/withdraw [post]
func (t *TransferController) WithdrawHander(c *gin.Context) {
	payload := model.Withdraw{}
	c.ShouldBind(&payload)
	claims, exists := c.Get("claims")
	if !exists {
		common.SendErrorResponse(c, http.StatusBadRequest, "Sepertinya login anda tidak valid")
		return
	}
	payload.UserId = claims.(*common.JwtClaim).DataClaims.Id
	_, err := t.uc.FindRekening(payload.UserId)
	if err != nil {
		if err.Error() == "1" {
			common.SendErrorResponse(c, http.StatusBadRequest, "rekening tidak ada, silahkan atur rekening anda")
			return
		}
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	saldo, _ := t.uc.GetBalanceCase(payload.UserId)
	saldo.Saldo -= payload.Withdraw
	if saldo.Saldo < 0 {
		common.SendErrorResponse(c, http.StatusBadRequest, "saldo anda tidak mencukup untuk transfer")
		return
	}
	response, err := t.ut.Withdraw(payload, saldo)
	if err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	common.SendCreateResponse(c, "SUCCESS", response)
}

// GetHistorytWithdraw godoc
// @Tags Transfer
// @Summary User can get histories withdraw
// @Description User can get histories withdraw
// @Accept json
// @Produce json
// @Security 	ApiKeyAuth
// @Param 		Authorization header string true "Bearer"
// @Success 200 {object} []model.Withdraw
// @Router /transfer/withdraw [get]
func (t *TransferController) GetWithdrawsHandler(c *gin.Context) {
	var page int
	page, _ = strconv.Atoi(c.Query("page"))
	if page == 0 {
		page = 1
	}
	claims, exists := c.Get("claims")
	if !exists {
		common.SendErrorResponse(c, http.StatusBadRequest, "Sepertinya login anda tidak valid")
		return
	}
	id := claims.(*common.JwtClaim).DataClaims.Id

	datas, err := t.ut.GetAllWithDraw(id, page)
	if err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendSingleResponse(c, "SUCCESS", datas)
}

func (t *TransferController) Route() {
	rg := t.rg.Group("/transfer")
	{
		// tulis route disini
		rg.POST("/", common.JWTAuth("user"), t.TransferHandler)
		rg.POST("/withdraw", common.JWTAuth("user"), t.WithdrawHander)
		rg.GET("/withdraw", common.JWTAuth("user"), t.GetWithdrawsHandler)
		rh := rg.Group("/history")
		{
			rh.GET("/send", common.JWTAuth("user"), t.GetSendTransferHandler)
			rh.GET("/receive", common.JWTAuth("user"), t.GetReceiveTransferHandler)
			rh.GET("/admin/send/:id", common.JWTAuth("admin"), t.AdminGetSendTransferHandler)
			rh.GET("/admin/receive/:id", common.JWTAuth("admin"), t.AdminGetReceiveTransferHandler)
		}
	}
}

func NewTransferController(ut usecase.TransferUseCase, uc usecase.UserUseCase, rg *gin.RouterGroup) *TransferController {
	return &TransferController{uc: uc, ut: ut, rg: rg}
}
