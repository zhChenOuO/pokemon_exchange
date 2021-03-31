package main

import (
	"fmt"
	"os"
	cmd "pokemon/cmd"

	cobra "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{Use: "server migration"}

func main() {
	rootCmd.AddCommand(cmd.ServerCmd, cmd.MigrationCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
