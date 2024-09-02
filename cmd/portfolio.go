package cmd

import (
	"github.com/ganimtron-10/EquityFusion/app"
	"github.com/spf13/cobra"
)

var userId string

var portfolioCmd = &cobra.Command{
	Use:   "portfolio",
	Short: "Show user Portfolio",
	// Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		app.PrintPortfolio(userId)
	},
}

func init() {
	rootCmd.AddCommand(portfolioCmd)
	exposeCmd.Flags().StringVarP(&userId, "id", "i", "", "user id for the portfolio")
	orderCmd.MarkFlagRequired("id")
}
