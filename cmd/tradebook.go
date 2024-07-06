package cmd

import (
	"github.com/ganimtron-10/EquityFusion/app"
	"github.com/spf13/cobra"
)

var tradebookCmd = &cobra.Command{
	Use:   "tradebook",
	Short: "Print TradeBook",
	// Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		app.PrintTradeBook()
	},
}

func init() {
	rootCmd.AddCommand(tradebookCmd)
}
