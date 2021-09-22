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
	yellow = ansi.ColorFunc("yellow")
	red    = ansi.ColorFunc("red")
	green  = ansi.ColorFunc("green")
)

var (
	client *hnapi.HNClient
)

var rootCmd = &cobra.Command{
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		client = hnapi.NewHNClient()
	},
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// Execute the command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
