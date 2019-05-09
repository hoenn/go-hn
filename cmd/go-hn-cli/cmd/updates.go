package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(getUpdatesCmd)
}

var getUpdatesCmd = &cobra.Command{
	Use:   "GetUpdates",
	Short: "Gets the latest updates from hackernews api",
	Long:  "Gets the latest profiles, posts, and comments from the hackernews api",
	Run: func(cmd *cobra.Command, args []string) {
		updates, err := client.Updates()
		if err != nil {
			fmt.Println(red(fmt.Sprintf("An error occured: %v", err)))
		}
		fmt.Print(prettyPrint(updates))
	},
}
