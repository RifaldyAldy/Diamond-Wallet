package controller

import (
	"fmt"
	"net/http"

	"github.com/RifaldyAldy/diamond-wallet/model"
	"github.com/RifaldyAldy/diamond-wallet/model/dto"
	"github.com/RifaldyAldy/diamond-wallet/usecase"
	"github.com/RifaldyAldy/diamond-wallet/utils/common"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	uc usecase.UserUseCase
	rg *gin.RouterGroup
}

// GetBalance 	godoc
// @Summary 	admin Get balance user with id param.
// @Description	Return the balance user with jwt id
// @Produce 	application/json
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer"
// @Param		id path string true "User ID"
// @Tags 		Admin
// @Success 	200 {object} model.UserSaldo
// @Router		/users/{id} [get]
func (e *UserController) getHandler(c *gin.Context) {
	id := c.Param("id")

	response, err := e.uc.GetBalanceCase(id)
	if err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, "Error "+err.Error())
		return
	}

	common.SendSingleResponse(c, "OK", response)
}

// CreateUser 	godoc
// @Summary 	User register account.
// @Description	Return the info user
// @Produce 	application/json
// @Param 		User body dto.UserRequestDto true "Create user"
// @Tags 		User
// @Success 	201 {object} model.User
// @Router		/users [post]
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

	common.SendCreateResponse(c, "SUCCESS", payloadResponse)
}

// LoginUser 	godoc
// @Summary 	User login to get jwtAuth
// @Description	Return the access token and id user
// @Produce 	application/json
// @Param		Login body dto.LoginRequestDto true "Login form"
// @Tags 		User
// @Success 	200 {object} dto.LoginResponseDto
// @Router		/users/login [post]
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

// GetBalance 	godoc
// @Summary 	User Get their info balance
// @Description	Return the info user and balance
// @Produce 	application/json
// @Security 	ApiKeyAuth
// @Param 		Authorization header string true "Bearer"
// @Tags 		User
// @Success 	200 {object} model.UserSaldo
// @Router		/users/saldo [get]
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

// UpdateUser 	godoc
// @Summary 	User edit account
// @Description	Return new the info user
// @Produce 	application/json
// @Security 	ApiKeyAuth
// @Param 		Authorization header string true "Bearer"
// @Param		Edit body dto.UserRequestEditDto true "Edit form"
// @Tags 		User
// @Success 	200 {object} model.User
// @Router		/users [PUT]
func (s *UserController) UpdateHandler(c *gin.Context) {
	claims, exists := c.Get("claims")
	if !exists {
		common.SendErrorResponse(c, http.StatusBadRequest, "sepertinya login anda tidak valid")
		return
	}
	id := claims.(*common.JwtClaim).DataClaims.Id
	var payload dto.UserRequestDto
	if err := c.BindJSON(&payload); err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	updatedUser, err := s.uc.UpdateUser(id, payload)
	if err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	common.SendSingleResponse(c, "UPDATE SUCCESS", updatedUser)
}

// VerifyUser 	godoc
// @Summary 	User verify account
// @Description	Return info user data
// @Accept 		multipart/form-data
// @Produce 	json
// @Security 	ApiKeyAuth
// @Param 		Authorization header string true "Bearer"
// @Param 		user formData string true "JSON data of user"
// @Param 		photo formData file true "Photo"
// @Tags 		User
// @Success 	201 {object} dto.VerifyUser
// @Router		/users/verify [POST]
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

// UpdatePinHandler godoc
// @Tags User
// @Summary Update user PIN
// @Description Update user PIN with JSON payload
// @Accept json
// @Produce json
// @Security 	ApiKeyAuth
// @Param 		Authorization header string true "Bearer"
// @Param payload body dto.UpdatePinRequestSwag true "Update PIN Request Payload"
// @Success 200 {object} dto.UpdatePinResponse
// @Router /users/pin [put]
func (p *UserController) UpdatePinHandler(c *gin.Context) {
	var payload dto.UpdatePinRequest
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
	payloadResponse, _ := p.uc.GetBalanceCase(payload.UserId)

	if payload.OldPin != payloadResponse.Pin {
		common.SendErrorResponse(c, http.StatusBadRequest, "Old pin not match")
		return
	}

	payload.OldPin = payloadResponse.Pin

	response, err := p.uc.UpdatePinUser(payload)
	if err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendSingleResponse(c, "success", response)
}

// CreateRekening godoc
// @Tags User
// @Summary Create Rekening User
// @Description Create Rekening user to withdraw
// @Accept json
// @Produce json
// @Security 	ApiKeyAuth
// @Param 		Authorization header string true "Bearer"
// @Param rekening body dto.RekeningDtoSwag true "Rekening data"
// @Success 201 {object} model.Rekening
// @Router /users/rekening [post]
func (p *UserController) CreateRekeningHandler(c *gin.Context) {
	var payload model.Rekening
	err := c.ShouldBind(&payload)
	if err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	claims, exists := c.Get("claims")
	if !exists {
		common.SendErrorResponse(c, http.StatusUnauthorized, "sepertinya login anda tidak valid")
		return
	}
	payload.UserId = claims.(*common.JwtClaim).DataClaims.Id
	res, err := p.uc.CreateRekening(payload)
	if err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendCreateResponse(c, "SUCCESS", res)
}

// GetRekening godoc
// @Tags User
// @Summary Get Rekening User
// @Description Get Rekening user
// @Accept json
// @Produce json
// @Security 	ApiKeyAuth
// @Param 		Authorization header string true "Bearer"
// @Success 200 {object} model.Rekening
// @Router /users/rekening [get]
func (p *UserController) GetRekeningHandler(c *gin.Context) {
	claims, exists := c.Get("claims")
	if !exists {
		common.SendErrorResponse(c, http.StatusUnauthorized, "sepertinya login anda tidak valid")
		return
	}
	id := claims.(*common.JwtClaim).DataClaims.Id
	res, err := p.uc.FindRekening(id)
	if err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendSingleResponse(c, "SUCCESS", res)
}

func (p *UserController) Route() {
	p.rg.POST("/users/login", p.loginHandler)
	p.rg.POST("/users", p.createHandler)
	p.rg.GET("/users/:id", common.JWTAuth("admin"), p.getHandler)
	p.rg.GET("/users/saldo", common.JWTAuth("user"), p.CheckBalance)
	p.rg.PUT("/users", common.JWTAuth("user"), p.UpdateHandler)
	p.rg.POST("/users/verify", common.JWTAuth("user"), p.VerifyHandler)
	p.rg.PUT("/users/pin", common.JWTAuth("user"), p.UpdatePinHandler)
	p.rg.POST("/users/rekening", common.JWTAuth("user"), p.CreateRekeningHandler)
	p.rg.GET("/users/rekening", common.JWTAuth("user"), p.GetRekeningHandler)

}

func NewUserController(uc usecase.UserUseCase, rg *gin.RouterGroup) *UserController {
	return &UserController{
		uc: uc,
		rg: rg,
	}
}
