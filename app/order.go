package app

import "sort"

type Order struct {
	ID       string  `json:"id"`
	Symbol   string  `json:"symbol"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

func SortList(List []Order, IsBuy bool) {
	sort.SliceStable(List, func(i, j int) bool {
		if IsBuy {
			return List[i].Price > List[j].Price
		} else {
			return List[i].Price < List[j].Price
		}
	})
}

func AddOrder(order Order, IsBuy bool) {
	if IsBuy {
		buyerList = append(buyerList, order)
		go SortList(buyerList, IsBuy)
	} else {
		sellerList = append(sellerList, order)
		go SortList(sellerList, IsBuy)
	}
}

func GetAllOrders(IsBuy bool) []Order {

	if IsBuy {
		return buyerList
	} else {
		return sellerList
	}
}
