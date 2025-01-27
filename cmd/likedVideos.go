package cmd

import (
	"encoding/json"
	"fmt"
	"ytcli/yt"

	"github.com/spf13/cobra"
)

var likedVideosCmd = &cobra.Command{
	Use:   "likedvideos",
	Short: "Lists all liked videos",
	Long:  "List all liked videos associated with auth'd user",
	Run: func(cmd *cobra.Command, args []string) {
		vids := yt.LikedVideos()
		json, _ := json.Marshal(vids)
		fmt.Println(string(json))
	},
}

func init() {
	rootCmd.AddCommand(likedVideosCmd)
}
