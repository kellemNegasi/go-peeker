package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const version = "1.0.0"

var rootCmd = &cobra.Command{
	Use:   "go-peeker",
	Short: "go-peeker peeks a large csv file.",
	Long:  "go-peeker peeks a large csv file.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("it works!")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
