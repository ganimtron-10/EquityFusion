package cmd

import (
	"fmt"

	"github.com/ganimtron-10/EquityFusion/app"
	"github.com/spf13/cobra"
)

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register User",
	// 	Long: `A longer description that spans multiple lines and likely contains examples
	// and usage of using your command. For example:

	// Cobra is a CLI library for Go that empowers applications.
	// This application is a tool to generate the needed files
	// to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Create a seperate command user and then mange multiple users using Persistent flag for passing the id
		fmt.Printf("User Created with ID: %s\n", app.RegisterUser().ID)
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)
}
