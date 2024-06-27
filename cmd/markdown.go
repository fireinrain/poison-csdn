package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var markdownCmd = &cobra.Command{
	Use:   "markdown",
	Short: "Add baned words in markdown file",
	Long:  `Add baned words in markdown file.`,
	Run: func(cmd *cobra.Command, args []string) {
		repoPath, _ := cmd.Flags().GetString("path")
		fmt.Printf("Hello, %s!\n", repoPath)
	},
}

func init() {
	markdownCmd.Flags().StringP("path", "p", ".", "path of repository")
}
