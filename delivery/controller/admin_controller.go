package controller

import (
	"net/http"

	"github.com/RifaldyAldy/diamond-wallet/model"
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

func (a *AdminController) GetUserInfo(c *gin.Context) {
	userID := c.Param("id")

	user, err := a.ua.GetUserInfo(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (a *AdminController) Route() {
	rg := a.rg.Group("/admin")
	{
		rg.POST("/", a.RegisterHandler)
		rg.GET("/user/:id", a.GetUserInfo)
	}
}

func NewAdminController(ua usecase.AdminUseCase, uc usecase.UserUseCase, rg *gin.RouterGroup) *AdminController {
	return &AdminController{ua: ua, uc: uc, rg: rg}
}
