package controller

import (
	"net/http"

	"github.com/RifaldyAldy/diamond-wallet/model"
	"github.com/RifaldyAldy/diamond-wallet/model/dto"
	"github.com/RifaldyAldy/diamond-wallet/usecase"
	"github.com/RifaldyAldy/diamond-wallet/utils/common"
	"github.com/RifaldyAldy/diamond-wallet/utils/encryption"
	"github.com/gin-gonic/gin"
)

type AdminController struct {
	ua usecase.AdminUseCase
	uc usecase.UserUseCase
	rg *gin.RouterGroup
}

func (a *AdminController) RegisterHandler(c *gin.Context) {
	payload := model.Admin{}
	c.ShouldBind(&payload)
	payload.Password, _ = encryption.HashPassword(payload.Password)
	res, err := a.ua.RegisterAdmin(payload)
	if err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	common.SendSingleResponse(c, "SUCCESS", res)
}

func (a *AdminController) LoginHandler(c *gin.Context) {
	payload := dto.LoginRequestDto{}
	c.ShouldBind(&payload)

	response, err := a.ua.LoginAdmin(payload)
	if err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	common.SendSingleResponse(c, "SUCCESS", response)
}

func (a *AdminController) Route() {
	rg := a.rg.Group("/admin")
	{
		rg.POST("/", a.RegisterHandler)
		rg.POST("/login", a.LoginHandler)
	}
}

func NewAdminController(ua usecase.AdminUseCase, uc usecase.UserUseCase, rg *gin.RouterGroup) *AdminController {
	return &AdminController{ua: ua, uc: uc, rg: rg}
}
