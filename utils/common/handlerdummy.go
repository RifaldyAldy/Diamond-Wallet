package common

import "github.com/gin-gonic/gin"

func HandlerDummy(c *gin.Context) {
	c.JSON(200, "OK")
}
