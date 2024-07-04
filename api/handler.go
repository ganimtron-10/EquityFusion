package api

import (
	"log"
	"net/http"

	"github.com/ganimtron-10/EquityFusion/app"
	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

func checkHealth(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, "App is running fine!")
}

func RegisterUser(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, app.RegisterUser())
}

func GetSellOrders(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, app.GetAllOrders(false))
}

func GetBuyOrders(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, app.GetAllOrders(true))
}

func AddSellOrder(context *gin.Context) {
	var newOrder app.Order

	if err := context.BindJSON(&newOrder); err != nil {
		context.IndentedJSON(http.StatusBadRequest, ErrorResponse{StatusCode: http.StatusBadRequest, Message: "Unable to parse the request body."})
		log.Print(err.Error())
		return
	}

	app.AddOrder(newOrder, false)

	context.IndentedJSON(http.StatusCreated, newOrder)
}

func AddBuyOrder(context *gin.Context) {
	var newOrder app.Order

	if err := context.BindJSON(&newOrder); err != nil {
		context.IndentedJSON(http.StatusBadRequest, ErrorResponse{StatusCode: http.StatusBadRequest, Message: "Unable to parse the request body."})
		log.Print(err.Error())
		return
	}

	app.AddOrder(newOrder, true)

	context.IndentedJSON(http.StatusCreated, newOrder)
}
