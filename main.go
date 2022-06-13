package main

import (
	"fmt"
	"os"

	"github.com/shuyaeer/fingo/cmd"
)

func main() {
	cmd.RootCmd.Flags().StringP("extention", "e", "", "Target File extention")
	cmd.RootCmd.Flags().StringP("name", "n", "", "Target File name")
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
		os.Exit(-1)
	}
}
