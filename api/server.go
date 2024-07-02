package api

import (
	"log"

	"github.com/gin-gonic/gin"
)

func CreateAndInitServer(port string) {
	router := gin.Default()

	router.GET("/register", RegisterUser)

	router.GET("/sell", AddSellOrder)
	router.GET("/buy", AddBuyOrder)

	log.Print("Starting Server on Port ", port)
	router.Run("localhost:" + port)
}
