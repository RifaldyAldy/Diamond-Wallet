package controller

import (
	"fmt"
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

func (e *UserController) getHandler(c *gin.Context) {
	id := c.Param("id")

	response, err := e.uc.GetBalanceCase(id)
	if err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, "Error "+err.Error())
		return
	}

	common.SendSingleResponse(c, "OK", response)
}

func (e *UserController) createHandler(c *gin.Context) {
	var payload dto.UserRequestDto
	if err := c.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	payloadResponse, err := e.uc.CreateUser(payload)
	if err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendSingleResponse(c, "success", payloadResponse)
}

func (u *UserController) loginHandler(c *gin.Context) {
	var payload dto.LoginRequestDto
	if err := c.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	loginData, err := u.uc.LoginUser(payload)
	if err != nil {
		if err.Error() == "1" {
			common.SendErrorResponse(c, http.StatusForbidden, "Password salah")
			return
		}
		common.SendErrorResponse(c, http.StatusForbidden, err.Error())
		return
	}
	common.SendSingleResponse(c, "success", loginData)
}

func (u *UserController) CheckBalance(c *gin.Context) {
	claims, exists := c.Get("claims")
	if !exists {
		common.SendErrorResponse(c, http.StatusInternalServerError, "Claims jwt tidak ada!")
		return
	}
	id := claims.(*common.JwtClaim).DataClaims.Id
	response, err := u.uc.GetBalanceCase(id)
	if err != nil {
		if err.Error() == "1" {
			common.SendErrorResponse(c, http.StatusBadRequest, "Verifikasi akun anda terlebih dahulu untuk akses cek saldo")
			return
		}
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	common.SendSingleResponse(c, "SUCCESS", response)
}

func (s *UserController) UpdateHandler(c *gin.Context) {
	claims, exists := c.Get("claims")
	if !exists {
		common.SendErrorResponse(c, http.StatusBadRequest, "Invalid Request JWT")
		return
	}
	id := claims.(*common.JwtClaim).DataClaims.Id
	var payload dto.UserRequestDto
	if err := c.BindJSON(&payload); err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, "Invalid Request Payload")
		return
	}

	updatedUser, err := s.uc.UpdateUser(id, payload)
	if err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	common.SendSingleResponse(c, "success", updatedUser)
}

func (u *UserController) VerifyHandler(c *gin.Context) {
	var payload dto.VerifyUser
	claims, exists := c.Get("claims")
	if !exists {
		common.SendErrorResponse(c, http.StatusInternalServerError, "Claims jwt tidak ada!")
		fmt.Println(2)
		return
	}
	payload, err := common.FileVerifyHandler(c)
	if err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, "failed upload photo"+err.Error())
		return
	}
	payload.UserId = claims.(*common.JwtClaim).DataClaims.Id

	response, err := u.uc.VerifyUser(payload)
	if err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	common.SendCreateResponse(c, "success", response)
}

func (p *UserController) UpdatePinHandler(c *gin.Context) {
	var payload dto.UpdatePinUser
	if err := c.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	claims, exists := c.Get("claims")
	if !exists {
		common.SendErrorResponse(c, http.StatusInternalServerError, "Claims jwt tidak ada!")
		return
	}
	payload.UserId = claims.(*common.JwtClaim).DataClaims.Id
	response, err := p.uc.UpdatePinUser(payload)
	fmt.Println("ini payload", payload)
	if err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Println("ini response", response)

	common.SendSingleResponse(c, "success", response)
}

func (p *UserController) Route() {
	rg := p.rg.Group("users")
	{
		rg.POST("/login", p.loginHandler)
		rg.POST("/", p.createHandler)
		rg.GET("/:id", common.JWTAuth("admin"), p.getHandler)
		rg.GET("/saldo", common.JWTAuth("user"), p.CheckBalance)
		rg.PUT("/", common.JWTAuth("user"), p.UpdateHandler)
		rg.POST("/verify", common.JWTAuth("user"), p.VerifyHandler)
		rg.PUT("/pin", common.JWTAuth("user"), p.UpdatePinHandler)
	}
}

func NewUserController(uc usecase.UserUseCase, rg *gin.RouterGroup) *UserController {
	return &UserController{
		uc: uc,
		rg: rg,
	}
}
