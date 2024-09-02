package app

import "fmt"

var buyerList []Order
var sellerList []Order

var tradebook []TradeLog

type TradeLog struct {
	BuyUserID  string  `json:"buyUserId"`
	SellUserID string  `json:"sellUserId"`
	Symbol     string  `json:"symbol"`
	Quantity   int     `json:"quantity"`
	BuyPrice   float64 `json:"buyPrice"`
	SellPrice  float64 `json:"sellPrice"`
}

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"

// TODO: Add more explicit logging thoughout just to make sure transactions are correct and eaiser for us to navigate the process
// Also add necessary locks to make sure async functions are working fine

func LogTrade(BuyOrder, SellOrder Order) {

	tradeQuantity := min(BuyOrder.Quantity, SellOrder.Quantity)
	tradelog := TradeLog{BuyUserID: BuyOrder.ID, SellUserID: SellOrder.ID, Symbol: BuyOrder.Symbol, Quantity: tradeQuantity, BuyPrice: BuyOrder.Price, SellPrice: SellOrder.Price}

	AddHolding(BuyOrder.ID, tradelog)
	tradebook = append(tradebook, tradelog)
}

func PrintTradeLog(log TradeLog, srnum int) {

	fmt.Printf("|%10v|%10v|%10v|%10v|%10.5f|\n", srnum, log.BuyUserID, log.Symbol, log.Quantity, log.BuyPrice)
	fmt.Printf("|%10v|%10v|%10v|%10v|%10.5f|\n", "", log.SellUserID, log.Symbol, log.Quantity, log.SellPrice)

}

func PrintTradeBook() {

	// if len(tradebook) == 0 {
	// 	tradeString := fmt.Sprintf("|%10v|%10v|%10v|%10v|%10v|", "Sr No.", "UserID", "Symbol", "Quantity", "Price")
	// 	tradebook = append(tradebook, tradeString)
	// }
	// tradeQuantity := min(BuyOrder.Quantity, SellOrder.Quantity)

	// tradeString := fmt.Sprintf("|%10v|%10v|%10v|%10v|%10.5f|", len(tradebook)/2+1, BuyOrder.ID, BuyOrder.Symbol, tradeQuantity, BuyOrder.Price)
	// tradebook = append(tradebook, Green+tradeString+Reset)

	// tradeString = fmt.Sprintf("|%10v|%10v|%10v|%10v|%10.5f|", "", SellOrder.ID, SellOrder.Symbol, tradeQuantity, SellOrder.Price)
	// tradebook = append(tradebook, Red+tradeString+Reset)

	// for _, log := range tradebook {
	// 	// TODO: Instead of printing return a slice of TradeLog
	// 	fmt.Println(log)
	// }

	if len(tradebook) == 0 {
		fmt.Println("Empty Tradebook! No trade executed yet.")
		return
	}

	fmt.Printf("|%10v|%10v|%10v|%10v|%10v|", "Sr No.", "UserID", "Symbol", "Quantity", "Price")
	for _, log := range tradebook {
		PrintTradeLog(log, len(tradebook)/2+1)
	}
}

func ReturnTradeBook() []TradeLog {
	return tradebook
}

func matchOrder(BuyOrder, SellOrder Order) {
	buyerList = buyerList[1:]
	sellerList = sellerList[1:]

	if BuyOrder.Quantity > SellOrder.Quantity {
		tempOrder := BuyOrder
		tempOrder.Quantity = BuyOrder.Quantity - SellOrder.Quantity
		AddOrder(tempOrder)
	} else if BuyOrder.Quantity < SellOrder.Quantity {
		tempOrder := SellOrder
		tempOrder.Quantity = SellOrder.Quantity - BuyOrder.Quantity
		AddOrder(tempOrder)
	}
	LogTrade(BuyOrder, SellOrder)
}

func CheckAndMatchOrders() {
	for (len(buyerList) > 0 && len(sellerList) > 0) && (buyerList[0].Symbol == sellerList[0].Symbol) && (buyerList[0].Price >= sellerList[0].Price) {
		matchOrder(buyerList[0], sellerList[0])
	}
}
