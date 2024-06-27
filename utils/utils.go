package utils

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"os"
	"strings"
	"time"
)

const AntiFlag = "Anti CSDN Stolen Repo With Baned Words"

// CheckIfFileExist 判断文件是否存在
func CheckIfFileExist(fileName string, fileDir string) bool {
	absPath := fileDir + string(os.PathSeparator) + fileName
	_, err := os.Stat(absPath)
	return err == nil || os.IsExist(err)

}

// CheckIfDirExist 判断文件夹是否存在
func CheckIfDirExist(parentDir string, fileDir string) bool {
	absPath := parentDir + string(os.PathSeparator) + fileDir
	fileInfo, err := os.Stat(absPath)
	return err == nil || os.IsExist(err) && fileInfo.IsDir()
}

// CheckIfLicenseExist 判断是否为LICENSE
func CheckIfLicenseExist(fileDir string) bool {
	return CheckIfFileExist("LICENSE", fileDir)
}

// CheckIfReadmeExist 判断是否为README
func CheckIfReadmeExist(fileDir string) bool {
	return CheckIfFileExist("README.md", fileDir)
}

// CheckIfTempFileExist 判断是否为临时文件
func CheckIfTempFileExist(fileDir string) bool {
	return CheckIfFileExist("tempfile.txt", fileDir)
}

// CheckIfGitRepoExist 判断是否为git仓库
func CheckIfGitRepoExist(fileDir string) bool {
	return CheckIfDirExist(fileDir, ".git")
}

// AppendFileStrData 在文件末尾追加数据
func AppendFileStrData(fileName string, fileDir string, data string) error {
	absPath := fileDir + string(os.PathSeparator) + fileName
	f, err := os.OpenFile(absPath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err = f.WriteString(data); err != nil {
		return err
	}
	return nil
}

// CheckIfFileHasPoison 判断文件是否存在anti words
func CheckIfFileHasPoison(fileName string, fileDir string) bool {
	//read file and find if a string in data or not
	absPath := fileDir + string(os.PathSeparator) + fileName
	data, err := os.ReadFile(absPath)
	if err != nil {
		return false
	}
	return strings.Contains(string(data), AntiFlag)
}

func GitAddAndCommit(fileName string, fileDir string) {
	// Open the existing repository
	repo, err := git.PlainOpen(fileDir)
	if err != nil {
		fmt.Println("> Error opening repository:", err)
		return
	}

	// Get the worktree for the repository
	worktree, err := repo.Worktree()
	if err != nil {
		fmt.Println("> Error getting worktree:", err)
		return
	}

	// Create a new file and write to it
	//filePath := "/path/to/your/repo/example.txt"
	//fileContent := []byte("This is an example file")
	//err = os.WriteFile(filePath, fileContent, 0644)
	//if err != nil {
	//	fmt.Println("Error writing file:", err)
	//	return
	//}

	// Add the file to the staging area
	absFilePath := fileDir + string(os.PathSeparator) + fileName
	_, err = worktree.Add(absFilePath)
	if err != nil {
		fmt.Println("> Error adding file:", err)
		return
	}

	// Create a new commit
	commitMsg := "Add data for anti CSDN"
	commit, err := worktree.Commit(commitMsg, &git.CommitOptions{
		Author: &object.Signature{
			Name:  "V for anti CSDN",
			Email: "v.for.anti.csdn@anti.com",
			When:  time.Now(),
		},
	})
	if err != nil {
		fmt.Println("> Error creating commit:", err)
		return
	}

	// Print the commit hash
	fmt.Println("> Git Commit hash:", commit)

	// Optional: print the commit details
	_, err = repo.CommitObject(commit)
	if err != nil {
		fmt.Println("> Error getting commit object:", err)
		return
	}

	fmt.Println("> Anti data has been commit with git.")
}
