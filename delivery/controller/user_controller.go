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

	response, err := e.uc.FindById(id)
	if err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, "Error "+err.Error())
		return
	}

	common.SendSingleResponse(c, "OK", response)
}

func (e *UserController) createHandler(c *gin.Context) {
	var payload dto.UserRequestDto
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	payloadResponse, err := e.uc.CreateUser(payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
		"data":    payloadResponse,
	})
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

func (u *UserController) CheckBalance(c *gin.Context) {
	claims, exists := c.Get("claims")
	if !exists {
		common.SendErrorResponse(c, http.StatusInternalServerError, "Claims jwt tidak ada!")
		return
	}
	id := claims.(*common.JwtClaim).DataClaims.Id
	response, err := u.uc.GetBalanceCase(id)
	if err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	common.SendSingleResponse(c, "SUCCESS", response)
}

func (s *UserController) UpdateHandler(c *gin.Context) {
	id := c.Param("id")
	var payload dto.UserRequestDto
	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Invalid Request Payload",
		})
		return
	}

	updatedUser, err := s.uc.UpdateUser(id, payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Update User Succesful",
		"data":    updatedUser,
	})
}

func (u *UserController) VerifyHandler(c *gin.Context) {
	var payload dto.VerifyUser
	claims, exists := c.Get("claims")
	if !exists {
		common.SendErrorResponse(c, http.StatusInternalServerError, "Claims jwt tidak ada!")
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
	}
	fmt.Println("ini response", response)

	common.SendSingleResponse(c, "success", response)
}

func (p *UserController) Route() {
	p.rg.POST("/users/login", p.loginHandler)
	p.rg.POST("/users", p.createHandler)
	p.rg.GET("/users/:id", p.getHandler)
	p.rg.GET("/users/saldo", common.JWTAuth("user"), p.CheckBalance)
	p.rg.PUT("/users/:id", p.UpdateHandler)
	p.rg.POST("/users/verify", common.JWTAuth("user"), p.VerifyHandler)
	p.rg.PUT("/users/pin", common.JWTAuth("user"), p.UpdatePinHandler)
}

func NewUserController(uc usecase.UserUseCase, rg *gin.RouterGroup) *UserController {
	return &UserController{
		uc: uc,
		rg: rg,
	}
}
