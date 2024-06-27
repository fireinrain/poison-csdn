package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"poison-csdn/baned"
	"poison-csdn/utils"
)

var tempfileCmd = &cobra.Command{
	Use:   "tempfile",
	Short: "Add baned words in a temp file.",
	Long:  `Add baned words in a temp file.`,
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
		tempfileName, _ := cmd.Flags().GetString("file")
		if bandwords == " " {
			//set default banded words
			bandwords = baned.AntiCSDNBanner[0] + "\n" + baned.BandWords + "\n" + baned.BandWords2 + "\n" + baned.AntiCSDNBanner[1]
		} else {
			bandwords = baned.AntiCSDNBanner[0] + "\n" + bandwords + "\n" + baned.AntiCSDNBanner[1]
		}
		PoisonWithTempfile(repoPath, bandwords, tempfileName, enableGitCommit)
	},
}

func init() {
	tempfileCmd.Flags().StringP("path", "p", ".", "path of repository")
	tempfileCmd.Flags().StringP("word", "w", " ", "custom baned words")
	tempfileCmd.Flags().StringP("file", "f", "tempfile.txt", "custom tempfile for baned words")
	tempfileCmd.Flags().BoolP("git", "g", false, "enable git commit or not")
}

func PoisonWithTempfile(repoPath string, antiWords string, tempfileName string, enableGitCommit bool) {
	if !utils.CheckIfGitRepoExist(repoPath) {
		fmt.Println("> The path need to be a git repository for using this command.")
		return
	}
	if !utils.CheckIfTempFileExist(repoPath, tempfileName) {
		fmt.Println("> No temp file for baned words, create default for current directory.")
		//create a txt file
		filePath := repoPath + string(os.PathSeparator) + tempfileName
		file, err := os.Create(filePath)
		if err != nil {
			fmt.Println("> Error creating file:", err)
			return
		}
		defer file.Close()
	}
	//check if has been poisoned
	hasPoison := utils.CheckIfFileHasPoison(tempfileName, repoPath)
	if hasPoison {
		fmt.Println("> Repo has been poisoned for anti CSDN.")
		return
	}
	// append anti data
	err := utils.AppendFileStrData(tempfileName, repoPath, antiWords)
	if err != nil {
		fmt.Println("> Append anti data to file error:", err)
		return
	}
	fmt.Println("> Append anti data to repo successfully.")

	// git add and commit
	if enableGitCommit {
		utils.GitAddAndCommit(tempfileName, repoPath)
	}

}
