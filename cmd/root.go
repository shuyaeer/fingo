package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// todo: 特定の拡張子のみ検出、ファイルの中身まで探索
var RootCmd = &cobra.Command{
	Use:   "fingo",
	Short: "File system crawler",
	Run: func(cmd *cobra.Command, args []string) {
		exec(args)
	},
}

func exec(args []string) {
	if err := filepath.Walk(args[0], crawl); err != nil {
		fmt.Println(err)
	}
}

func crawl(path string, file_info os.FileInfo, err error) error {
	if file_info.IsDir() {
		fmt.Printf(" dir : %s\n", path)
		return nil
	} else {
		fmt.Printf("file : %s\n", path)
	}
	return nil
}
