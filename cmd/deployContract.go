/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"

	"openmyth/blockchain/cmd/srv/deploy_contract"
)

// deployContractCmd represents the deployContract command
var deployContractCmd = &cobra.Command{
	Use:   "deployContract",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		srv := deploy_contract.NewServer()
		srv.Run(cmd.Context())
	},
}

func init() {
	rootCmd.AddCommand(deployContractCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deployContractCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deployContractCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
