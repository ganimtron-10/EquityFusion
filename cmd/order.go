package cmd

import (
	"github.com/ganimtron-10/EquityFusion/app"
	"github.com/spf13/cobra"
)

var newOrder app.Order

var orderCmd = &cobra.Command{
	Use:   "order",
	Short: "Create a Buy/Sell Order",
	// 	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		app.AddOrder(newOrder)
	},
}

func init() {
	rootCmd.AddCommand(orderCmd)

	orderCmd.Flags().StringVarP(&newOrder.ID, "id", "i", "", "user id")
	orderCmd.Flags().StringVarP(&newOrder.Symbol, "symbol", "s", "", "order symbol")
	orderCmd.Flags().IntVarP(&newOrder.Quantity, "quantity", "q", 0, "order quantity")
	orderCmd.Flags().Float64VarP(&newOrder.Price, "price", "p", 0, "order price")
	orderCmd.Flags().BoolVarP(&newOrder.IsBuy, "buy", "b", false, "sepcifies whether its a buy order")

	orderCmd.MarkFlagRequired("id")
	orderCmd.MarkFlagRequired("symbol")
	orderCmd.MarkFlagRequired("quantity")
	orderCmd.MarkFlagRequired("price")
}
