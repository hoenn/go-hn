package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/hoenn/go-hn/pkg/hnapi"
	"github.com/mgutz/ansi"
	"github.com/spf13/cobra"
)

var (
	red = ansi.ColorFunc("red")
)

var (
	client *hnapi.HNClient
)

var rootCmd = &cobra.Command{
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		client = hnapi.NewHNClient()
	},
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Help()
		if err != nil {
			errorMsgWithExit(err)
		}
	},
}

// Execute the command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		errorMsgWithExit(err)
	}
}

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func errorMsgWithExit(err error) {
	if err != nil {
		fmt.Println(red(fmt.Sprintf("An error occured: %s", err)))
		os.Exit(1)
	}
}
