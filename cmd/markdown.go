package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"path/filepath"
	"poison-csdn/baned"
	"poison-csdn/utils"
)

const MarkdownFileName = "README.md"

var markdownCmd = &cobra.Command{
	Use:   "markdown",
	Short: "Add baned words in markdown file",
	Long:  `Add baned words in markdown file.`,
	Run: func(cmd *cobra.Command, args []string) {
		repoPath, _ := cmd.Flags().GetString("path")
		//get abs path
		repoPath, err := filepath.Abs(repoPath)
		//println(repoPath)
		if err != nil {
			fmt.Printf("> Get absolute path error: %v\n", err)
			return
		}
		bandwords, _ := cmd.Flags().GetString("word")
		enableGitCommit, _ := cmd.Flags().GetBool("git")
		if bandwords == " " {
			//set default banded words
			bandwords = baned.AntiCSDNBanner[0] + "\n" + baned.BandWords + "\n" + baned.BandWords2 + "\n" + baned.AntiCSDNBanner[1]
		} else {
			bandwords = baned.AntiCSDNBanner[0] + "\n" + bandwords + "\n" + baned.AntiCSDNBanner[1]
		}
		PoisonWithMarkdown(repoPath, bandwords, enableGitCommit)
	},
}

func init() {
	markdownCmd.Flags().StringP("path", "p", ".", "path of repository")
	markdownCmd.Flags().StringP("word", "w", " ", "custom baned words")
	markdownCmd.Flags().BoolP("git", "g", false, "enable git commit or not")
}

func PoisonWithMarkdown(repoPath string, antiWords string, enableGitCommit bool) {
	antiWords = utils.CompactDispalyMarkdown(antiWords)
	if !utils.CheckIfGitRepoExist(repoPath) {
		fmt.Println("> The path need to be a git repository for using this command.")
		return
	}
	if !utils.CheckIfReadmeExist(repoPath) {
		fmt.Println("> The path need contains a README.md file for using this command.")
		return
	}
	//check if has been poisoned
	hasPoison := utils.CheckIfFileHasPoison(MarkdownFileName, repoPath)
	if hasPoison {
		fmt.Println("> Repo has been poisoned for anti CSDN.")
		return
	}
	// append anti data
	err := utils.AppendFileStrData(MarkdownFileName, repoPath, antiWords)
	if err != nil {
		fmt.Println("> Append anti data to file error:", err)
		return
	}
	fmt.Println("> Append anti data to repo successfully.")

	// git add and commit
	if enableGitCommit {
		utils.GitAddAndCommit(MarkdownFileName, repoPath)
	}

}
