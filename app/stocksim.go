package app

import "sort"

type Order struct {
	ID       string  `json:"id"`
	Symbol   string  `json:"symbol"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

var buyerList []Order
var sellerList []Order

func SortAndAppend(List []Order, CurOrder Order, IsBuy bool) []Order {
	List = append(List, CurOrder)
	sort.SliceStable(List, func(i, j int) bool {
		if IsBuy {
			return List[i].Price > List[j].Price
		} else {
			return List[i].Price < List[j].Price
		}
	})
	return List
}

func AddOrder(order Order, IsBuy bool) {
	if IsBuy {
		buyerList = SortAndAppend(buyerList, order, IsBuy)
	} else {
		sellerList = SortAndAppend(sellerList, order, IsBuy)
	}
}

func GetAllOrders(IsBuy bool) []Order {

	if IsBuy {
		return buyerList
	} else {
		return sellerList
	}
}
