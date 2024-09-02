package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/ganimtron-10/EquityFusion/app"
	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

type ReturnOrders struct {
	BuyOrders  []app.Order `json:"buy_orders"`
	SellOrders []app.Order `json:"sell_orders"`
}

func checkHealth(context *gin.Context) {
	context.JSON(http.StatusOK, "App is running fine!")
}

func RegisterUser(context *gin.Context) {
	context.JSON(http.StatusOK, app.RegisterUser())
}

func GetSellOrders(context *gin.Context) {
	context.JSON(http.StatusOK, app.GetAllOrders(false))
}

func GetBuyOrders(context *gin.Context) {
	context.JSON(http.StatusOK, app.GetAllOrders(true))
}

func AddSellOrder(context *gin.Context) {
	var newOrder app.Order

	if err := context.BindJSON(&newOrder); err != nil {
		context.JSON(http.StatusBadRequest, ErrorResponse{StatusCode: http.StatusBadRequest, Message: "Unable to parse the request body."})
		log.Print(err.Error())
		return
	}
	newOrder.IsBuy = false
	app.AddOrder(newOrder)

	context.JSON(http.StatusCreated, newOrder)
}

func AddBuyOrder(context *gin.Context) {
	var newOrder app.Order

	if err := context.BindJSON(&newOrder); err != nil {
		context.JSON(http.StatusBadRequest, ErrorResponse{StatusCode: http.StatusBadRequest, Message: "Unable to parse the request body."})
		log.Print(err.Error())
		return
	}

	newOrder.IsBuy = true
	app.AddOrder(newOrder)

	context.JSON(http.StatusCreated, newOrder)
}

func SeedDummyOrders(context *gin.Context) {
	var Number int

	if context.Query("number") != "" {
		parsedNumber, err := strconv.Atoi(context.Query("number"))
		if err == nil {
			Number = parsedNumber
		}
	}

	app.SeedDummyOrders(Number)
	context.JSON(http.StatusOK, ReturnOrders{BuyOrders: app.GetAllOrders(true), SellOrders: app.GetAllOrders(false)})
}

func GetTradebook(context *gin.Context) {
	context.JSON(http.StatusOK, app.ReturnTradeBook())
}

func GetPortfolio(context *gin.Context) {
	var userId string

	if context.Query("userId") != "" {
		userId = context.Query("userId")
	}

	context.JSON(http.StatusOK, app.ReturnPortfolio(userId))
}
