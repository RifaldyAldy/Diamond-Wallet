package controller

import (
	"github.com/RifaldyAldy/diamond-wallet/usecase"
	"github.com/gin-gonic/gin"
)

type PingController struct {
	uc usecase.PingUseCase
	rg *gin.RouterGroup
}

func (p *PingController) GetHandler(c *gin.Context) {
	err := p.uc.Ping()
	if err != nil {
		c.String(404, err.Error())
		return
	}
	c.JSON(200, "PONG! - Database connected")
}

func (p *PingController) Route() {
	p.rg.GET("/ping", p.GetHandler)
}

func NewPingController(uc usecase.PingUseCase, rg *gin.RouterGroup) *PingController {
	return &PingController{
		uc: uc,
		rg: rg,
	}
}
