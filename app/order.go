package app

import (
	"fmt"
	"sort"
)

type Order struct {
	ID       string  `json:"id"`
	Symbol   string  `json:"symbol"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
	IsBuy    bool    `json:"isBuy,omitempty"`
}

func PrintOrder(Order Order) {
	var OrderType string
	if Order.IsBuy {
		OrderType = "Buy"
	} else {
		OrderType = "Sell"
	}
	fmt.Printf("%s Order Created with Details:\n\tUser ID: %s\n\tSymbol: %s\n\tQuantity: %v\n\tPrice: %.2f\n",
		OrderType, Order.ID, Order.Symbol, Order.Quantity, Order.Price)
}

func SortList(List []Order) {
	sort.SliceStable(List, func(i, j int) bool {
		if List[i].IsBuy {
			return List[i].Price > List[j].Price
		} else {
			return List[i].Price < List[j].Price
		}
	})
}

func AddOrder(NewOrder Order) {
	// TODO: Check this execution once doubtful about the sorting and matching are not in order. Might need defer for CheckAndMatchOrders.
	PrintOrder(NewOrder)
	if NewOrder.IsBuy {
		buyerList = append(buyerList, NewOrder)
		go SortList(buyerList)
	} else {
		sellerList = append(sellerList, NewOrder)
		go SortList(sellerList)
	}
	go CheckAndMatchOrders()
}

func GetAllOrders(IsBuy bool) []Order {

	if IsBuy {
		return buyerList
	} else {
		return sellerList
	}
}
