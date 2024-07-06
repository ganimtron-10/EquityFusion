package api

import (
	"log"

	"github.com/gin-gonic/gin"
)

func CreateAndInitServer(port string) {
	router := gin.Default()

	router.GET("/health", checkHealth)

	router.GET("/register", RegisterUser)
	router.GET("/sell", GetSellOrders)
	router.GET("/buy", GetBuyOrders)
	router.GET("/seed", SeedDummyOrders)
	router.GET("/tradebook", GetTradebook)

	router.POST("/sell", AddSellOrder)
	router.POST("/buy", AddBuyOrder)

	log.Print("Starting Server on Port ", port)
	router.Run("localhost:" + port)
}
