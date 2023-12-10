package main

import (
	"autocommitai/internal/Command"
	"fmt"
	"os"
)

func main() {

	var rootCmd = Command.RootCommand()
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Exit(0)
}
