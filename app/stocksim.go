package app

type Order struct {
	ID       string  `json:"id"`
	Symbol   string  `json:"symbol"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

var buyerList []Order
var sellerList []Order

func AddOrder(order Order, IsBuy bool) {
	if IsBuy {
		buyerList = append(buyerList, order)
	} else {
		sellerList = append(sellerList, order)
	}
}

func GetAllOrders(IsBuy bool) []Order {

	if IsBuy {
		return buyerList
	} else {
		return sellerList
	}
}
