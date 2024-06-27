package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var tempfileCmd = &cobra.Command{
	Use:   "tempfile",
	Short: "Add baned words in a temp file.",
	Long:  `Add baned words in a temp file.`,
	Run: func(cmd *cobra.Command, args []string) {
		repoPath, _ := cmd.Flags().GetString("path")
		fmt.Printf("Hello, %s!\n", repoPath)
	},
}
