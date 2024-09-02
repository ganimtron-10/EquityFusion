package cmd

import (
	"github.com/ganimtron-10/EquityFusion/api"
	"github.com/spf13/cobra"
)

var Port string

var exposeCmd = &cobra.Command{
	Use:   "expose",
	Short: "Exposes RESTful API for EquityFusion",
	// 	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		api.CreateAndInitServer(Port)
	},
}

func init() {
	rootCmd.AddCommand(exposeCmd)
	exposeCmd.Flags().StringVarP(&Port, "port", "p", "8080", "port exposing the RESTful API")
}
