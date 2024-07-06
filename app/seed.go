package app

import (
	"math/rand"
)

func SeedDummyOrders(Number int) {
	if Number == 0 {
		Number = 5 + rand.Intn(5)
	}
	for i := 0; i < Number; i++ {
		AddOrder(Order{ID: "1", Symbol: "SYM", Quantity: 1 + rand.Intn(50), Price: (1 + rand.Float64()) * rand.Float64(), IsBuy: false})
		AddOrder(Order{ID: "1", Symbol: "SYM", Quantity: 1 + rand.Intn(50), Price: (1 + rand.Float64()) * rand.Float64(), IsBuy: true})
	}
}
