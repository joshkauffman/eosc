package cmd

import (
	"github.com/spf13/cobra"
)

var voteCmd = &cobra.Command{
	Use:   "vote",
	Short: "Command to vote for Block Producers or delegate your vote to a proxy",
}

func init() {
	RootCmd.AddCommand(voteCmd)
}
