package cmd

import (
	"github.com/ganimtron-10/EquityFusion/app"
	"github.com/spf13/cobra"
)

var Number int

var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Seed Dummy Orders",
	// 	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		app.SeedDummyOrders(Number)
	},
}

func init() {
	rootCmd.AddCommand(seedCmd)

	seedCmd.Flags().IntVarP(&Number, "number", "n", 0, "order price")
}
