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

// CreateAdmin godoc
// @Tags Admin
// @Summary Admin can create account
// @Description Admin can create account
// @Accept json
// @Produce json
// @Param payload body dto.AdminRequestDto true "payload register"
// @Success 201 {object} model.Admin
// @Router /admin [post]
func (a *AdminController) RegisterHandler(c *gin.Context) {
	payload := model.Admin{}
	c.ShouldBind(&payload)
	payload.Password, _ = encryption.HashPassword(payload.Password)
	res, err := a.ua.RegisterAdmin(payload)
	if err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	common.SendCreateResponse(c, "SUCCESS", res)
}

// AdminLogin godoc
// @Tags Admin
// @Summary Admin can login
// @Description Admin can login
// @Accept json
// @Produce json
// @Param payload body dto.LoginRequestDto true "payload login"
// @Success 201 {object} dto.LoginResponseDto
// @Router /admin/login [post]
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

// AdminGetInfoUser godoc
// @Tags Admin
// @Summary Admin can get info and balance user
// @Description Admin can get info and balance user
// @Accept json
// @Produce json
// @Security 	ApiKeyAuth
// @Param 		Authorization header string true "Bearer"
// @Param id path string true "User ID"
// @Success 201 {object} model.User
// @Router /admin/user/{id} [get]
func (a *AdminController) GetUserInfo(c *gin.Context) {
	userID := c.Param("id")

	user, err := a.uc.FindById(userID)
	if err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	common.SendSingleResponse(c, "SUCCESS", user)
}

func (a *AdminController) Route() {
	rg := a.rg.Group("/admin")
	{
		rg.POST("/", a.RegisterHandler)
		rg.POST("/login", a.LoginHandler)
		rg.GET("/user/:id", common.JWTAuth("admin"), a.GetUserInfo)
	}
}

func NewAdminController(ua usecase.AdminUseCase, uc usecase.UserUseCase, rg *gin.RouterGroup) *AdminController {
	return &AdminController{ua: ua, uc: uc, rg: rg}
}
