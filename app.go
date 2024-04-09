package main

import (
	"github.com/RifaldyAldy/diamond-wallet/delivery"
	_ "github.com/RifaldyAldy/diamond-wallet/docs"
)

// @title Tag Service API
// @version 1.0
// @description A tag service API in Go using Gin framework

// @host localhost:8080
// @BasePath /api/v1

func main() {
	delivery.NewServer().Run()
}
