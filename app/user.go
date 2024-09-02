package app

import (
	"fmt"
	"math/rand"

	"github.com/google/uuid"
)

type User struct {
	ID string `json:"id"`
}

type Portfolio struct {
	UserID         string     `json:"userId"`
	CurrentHolding []TradeLog `json:"currentHolding"`
}

type CompletePortfolio struct {
	UserID            string     `json:"userId"`
	CurrentHolding    []TradeLog `json:"currentHolding"`
	PendingBuyOrders  []Order    `json:"pendingBuyOrders"`
	PendingSellOrders []Order    `json:"pendingSellOrders"`
}

var userList []string
var userPortfolios map[string]Portfolio = make(map[string]Portfolio)

func ChooseRandomUser() string {
	if len(userList) < 2 {
		for i := 0; i < rand.Intn(5); i++ {
			RegisterUser()
		}
	}
	return userList[rand.Intn(len(userList))]
}

func RegisterUser() User {
	userId := uuid.NewString()
	userList = append(userList, userId)
	userPortfolios[userId] = Portfolio{UserID: userId}
	return User{ID: userId}
}

func AddHolding(userId string, tradelog TradeLog) {
	holdingPortfolio := userPortfolios[userId]
	holdingPortfolio.CurrentHolding = append(holdingPortfolio.CurrentHolding, tradelog)
	userPortfolios[userId] = holdingPortfolio
}

func PrintPortfolio(userId string) {
	curPortfolio, ok := userPortfolios[userId]
	if !ok {
		fmt.Println("User doesn't exists, Please try Again!")
		return
	}
	fmt.Printf("Portfolio for User %v", curPortfolio.UserID)
	fmt.Println("\nCurrent Holdings: ")

	fmt.Printf("|%10v|%10v|%10v|%10v|%10v|\n", "Sr No.", "UserID", "Symbol", "Quantity", "Price")
	for i, tradelog := range curPortfolio.CurrentHolding {
		PrintTradeLog(tradelog, i+1)
	}

	fmt.Println("\nPending Buy Orders: ")
	for _, buyOrder := range getPendingOrdersByUserId(userId, true) {
		PrintOrder(buyOrder)
	}

	fmt.Println("\nPending Sell Orders: ")
	for _, sellOrder := range getPendingOrdersByUserId(userId, false) {
		PrintOrder(sellOrder)
	}

}

func ReturnPortfolio(userId string) CompletePortfolio {
	holdingPortfolio := userPortfolios[userId]

	return CompletePortfolio{UserID: holdingPortfolio.UserID, CurrentHolding: holdingPortfolio.CurrentHolding, PendingBuyOrders: getPendingOrdersByUserId(userId, true), PendingSellOrders: getPendingOrdersByUserId(userId, false)}
}
