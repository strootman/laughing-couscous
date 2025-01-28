package cmd

import (
	"encoding/json"
	"fmt"
	"ytcli/yt"

	"github.com/spf13/cobra"
)

/* Status: Currently struggling with YT API auth.  */
var listActivitiesCmd = &cobra.Command{
	Use:   "listactivities",
	Short: "Lists all activities",
	Long:  "List all activities associated with auth'd user",
	Run: func(cmd *cobra.Command, args []string) {
		vids := yt.Activities()
		json, _ := json.Marshal(vids)
		fmt.Println(string(json))
	},
}

func init() {
	rootCmd.AddCommand(listActivitiesCmd)
}
