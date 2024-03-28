package controller

import (
	"net/http"

	"github.com/RifaldyAldy/diamond-wallet/model/dto"
	"github.com/RifaldyAldy/diamond-wallet/usecase"
	"github.com/RifaldyAldy/diamond-wallet/utils/common"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	uc usecase.UserUseCase
	rg *gin.RouterGroup
}

func (u *UserController) loginHandler(c *gin.Context) {
	var payload dto.LoginRequestDto
	if err := c.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	loginData, err := u.uc.LoginUser(payload)
	if err != nil {
		if err.Error() == "1" {
			common.SendErrorResponse(c, http.StatusForbidden, err.Error())
			return
		}
	}
	if err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	common.SendSingleResponse(c, "success", loginData)
}

func (p *UserController) Route() {
	p.rg.POST("/users/login", p.loginHandler)
}

func NewUserController(uc usecase.UserUseCase, rg *gin.RouterGroup) *UserController {
	return &UserController{
		uc: uc,
		rg: rg,
	}
}
