package controller

import (
	"net/http"
	"strconv"

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

// tulis handler code kalian disini

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

func (t *TransferController) Route() {
	rg := t.rg.Group("/transfer")
	{
		// tulis route disini
		rg.POST("/", common.JWTAuth("user"), t.TransferHandler)
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
