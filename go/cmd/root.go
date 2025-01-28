package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ytcli",
	Short: "ytcli provides commands on your YT channels",
	Long:  "ytcli provides ReadOnly commands to view stats and other info on you authorized YouTube channels",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("This is the ytcli tool! Use --help for usage")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oops. An error while executing ytcli '%s'\n", err)
		os.Exit(1)
	}
}
