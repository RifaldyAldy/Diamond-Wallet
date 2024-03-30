package middleware

import (
	"log"
	"os"
	"time"

	"github.com/RifaldyAldy/diamond-wallet/model"
	"github.com/gin-gonic/gin"
)

func LogMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		file, err := os.OpenFile("Logger log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal("err", err.Error())
		}
		defer file.Close()
		log.SetOutput(file)

		t := time.Now()

		ctx.Next()

		logString := model.SendLogRequest(model.LogModel{
			AccesTime: t,
			Latency:   time.Since(t),
			ClientIP:  ctx.ClientIP(),
			Method:    ctx.Request.Method,
			Code:      ctx.Writer.Status(),
			Path:      ctx.Request.URL.Path,
			UserAgent: ctx.Request.UserAgent(),
		})

		_, err = file.WriteString(logString)
		if err != nil {
			log.Fatal("Failed to writer", err.Error())
		}
	}
}
