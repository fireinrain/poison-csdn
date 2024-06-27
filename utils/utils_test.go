package utils

import (
	"fmt"
	"testing"
)

func TestCheckIfFileExist(t *testing.T) {
	type args struct {
		fileName string
		fileDir  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "File exists",
			args: args{
				fileName: "utils.go",
				fileDir:  ".",
			},
			want: true,
		},
		{
			name: "File does not exist",
			args: args{
				fileName: "nonexistent_file.txt",
				fileDir:  "/path/to/directory/",
			},
			want: false,
		},
		{
			name: "Directory does not exist",
			args: args{
				fileName: "existing_file.txt",
				fileDir:  "/path/to/nonexistent_directory/",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckIfFileExist(tt.args.fileName, tt.args.fileDir); got != tt.want {
				t.Errorf("CheckIfFileExist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckIfGitRepoExist(t *testing.T) {
	type args struct {
		fileDir string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Git repo exists",
			args: args{
				fileDir: "../",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckIfGitRepoExist(tt.args.fileDir); got != tt.want {
				t.Errorf("CheckIfGitRepoExist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompactDispalyMarkdown(t *testing.T) {
	markdown := CompactDispalyMarkdown("fuck stolen")
	fmt.Println(markdown)
}
