package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type User struct {
	ID string `json:"session_id"`
}

func RegisterUser(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, User{ID: uuid.NewString()})
}

func AddSellOrder(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, "Beep Boop Beep... Under Construction")
}

func AddBuyOrder(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, "Beep Boop Beep... Under Construction")
}
