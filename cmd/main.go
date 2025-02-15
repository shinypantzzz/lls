package main

import (
	"fmt"
	"os"

	"github.com/shinypantzzz/lls/internal/dirreader"
	"github.com/shinypantzzz/lls/internal/sorting"
	"github.com/shinypantzzz/lls/internal/table"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "lss [path]",
	Short: "List files and subdirectories with sizes in a directory",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var dir string
		if len(args) == 0 {
			dir = "."
		} else {
			dir = args[0]
		}

		items, totalSize, err := dirreader.ReadDir(dir)
		if err != nil {
			fmt.Printf("Error reading directory: %v\n", err)
			os.Exit(1)
		}

		sorting.SortItems(items)

		table := table.BuildTable(items, totalSize, os.Stdout)

		table.Render()
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
