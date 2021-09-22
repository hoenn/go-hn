package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(maxItemCmd)
}

var maxItemCmd = &cobra.Command{
	Use:   "MaxItem",
	Short: "Gets the current max item id",
	Long:  "Gets the current max item id",
	Run: func(cmd *cobra.Command, args []string) {
		item, err := client.MaxItemID()
		if err != nil {
			errorMsgWithExit(err)
		}
		fmt.Printf("%v\n", prettyPrint(item))
	},
}
