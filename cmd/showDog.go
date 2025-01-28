package cmd

import (
	"encoding/json"
	"fmt"
	"ytcli/other"

	"github.com/spf13/cobra"
)

var download bool
var showDogCmd = &cobra.Command{
	Use:   "showDog",
	Short: "Fetches a random dog breed pic",
	Long:  "Fetches a random dog breed pic",
	Run: func(cmd *cobra.Command, args []string) {
		dog := other.GetDogBreedPic(download)
		json, _ := json.Marshal(dog)
		fmt.Println(string(json))
	},
}

func init() {
	showDogCmd.Flags().BoolVarP(&download, "download", "d", false, "Downloads the pic to the current directory")
	rootCmd.AddCommand(showDogCmd)
}
