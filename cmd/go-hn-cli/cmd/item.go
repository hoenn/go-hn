package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var itemID string

func init() {
	rootCmd.AddCommand(getItemCmd)
	getItemCmd.Flags().StringVarP(&itemID, "id", "i", "", "id of an item")
	err := getItemCmd.MarkFlagRequired("id")
	if err != nil {
		errorMsgWithExit(err)
	}
}

var getItemCmd = &cobra.Command{
	Use:   "GetItem",
	Short: "Get an item by ID from the hackernews api",
	Long:  "Gets an item (Story, Comment, Poll, PollOpt) from the hacker news api by its item ID",
	Run: func(cmd *cobra.Command, args []string) {
		item, err := client.Item(itemID)
		if err != nil {
			errorMsgWithExit(err)
		}
		fmt.Print(prettyPrint(item))
	},
}
