package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// todo: ファイルの中身まで探索
var RootCmd = &cobra.Command{
	Use:   "fingo",
	Short: "File system crawler",
	Run: func(cmd *cobra.Command, args []string) {
		file_ext, _ := cmd.Flags().GetString("extention")
		targetFileName, _ := cmd.Flags().GetString("name")
		// Todo: argument varidation
		crawl(args, file_ext, targetFileName)
	},
}

// Todo: split func
func crawl(args []string, fileExt string, targetFileName string) {
	if err := filepath.Walk(args[0], func(path string, file_info os.FileInfo, err error) error {
		if fileExt != "" {
			ext := filepath.Ext(path)
			if ext == "."+fileExt {
				fmt.Printf("file : %s\n", path)
				return nil
			} else {
				return nil
			}
			// may
		} else if targetFileName != "" {
			if strings.Contains(path, targetFileName) {
				fmt.Printf("file : %s\n", path)
				return nil
			} else {
				return nil
			}
		} else if file_info.IsDir() {
			fmt.Printf(" dir : %s\n", path)
			return nil
		} else {
			fmt.Printf("file : %s\n", path)
			return nil
		}
	}); err != nil {
		fmt.Println(err)
	}
}

// func matchExtentionName(fileExtention string) {

// }

// func matchFileName(targetFileName string) {

// }
