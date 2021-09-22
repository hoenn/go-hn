package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var userID string

func init() {
	rootCmd.AddCommand(getUserCmd)
	getUserCmd.Flags().StringVarP(&userID, "userid", "u", "", "id of a user")
	err := getUserCmd.MarkFlagRequired("userid")
	if err != nil {
		errorMsgWithExit(err)
	}
}

var getUserCmd = &cobra.Command{
	Use:   "GetUser",
	Short: "Gets a user's profile information by ID from the hackernews api",
	Long:  "Gets a user's profile information by ID from the hackernews api",
	Run: func(cmd *cobra.Command, args []string) {
		user, err := client.User(userID)
		if err != nil {
			errorMsgWithExit(err)
		}
		fmt.Print(prettyPrint(user))
	},
}
