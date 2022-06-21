package cmd

import (
	"fmt"

	"github.com/hoenn/go-hn/pkg/hnapi"
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
		item, err := client.GetItem(itemID)
		if err != nil {
			errorMsgWithExit(err)
		}
		switch item.Type {
		case hnapi.StoryItem, hnapi.JobItem, hnapi.AskItem:
			s, err := hnapi.ItemToStory(item)
			if err != nil {
				errorMsgWithExit(fmt.Errorf("could not read story item: %w", err))
			}
			fmt.Print(prettyPrint(s))
		case hnapi.CommentItem:
			c, err := hnapi.ItemToComment(item)
			if err != nil {
				errorMsgWithExit(fmt.Errorf("could not read comment item: %w", err))
			}
			fmt.Print(prettyPrint(c))
		case hnapi.PollItem:
			p, err := hnapi.ItemToPoll(item)
			if err != nil {
				errorMsgWithExit(fmt.Errorf("could not read poll item: %w", err))
			}
			fmt.Print(prettyPrint(p))
		case hnapi.PollOptItem:
			p, err := hnapi.ItemToPollOpt(item)
			if err != nil {
				errorMsgWithExit(fmt.Errorf("could not read pollopt item: %w", err))
			}
			fmt.Print(prettyPrint(p))
		default:
			fmt.Println(prettyPrint(item))
		}
	},
}
