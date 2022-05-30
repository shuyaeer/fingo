package main

import (
	"fmt"
	"os"

	"example.com/fingo/cmd"
)

func main() {
	cmd.RootCmd.Flags().StringP("extention", "e", "", "Target File extention")
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
		os.Exit(-1)
	}
}
