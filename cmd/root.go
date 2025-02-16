package cmd

import (
	"fmt"
	"os"

	"github.com/shinypantzzz/lls/internal/dirreader"
	"github.com/shinypantzzz/lls/internal/sorting"
	"github.com/shinypantzzz/lls/internal/table"
	"github.com/spf13/cobra"
)

var sortBy string

var sortByMap = map[string]sorting.SortFunc{
	"size": sorting.BySize,
	"name": sorting.ByName,
}

var reverse bool

var rootCmd = &cobra.Command{
	Use:   "lls [path]",
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

		userSortFunc := sortByMap[sortBy]
		if reverse {
			userSortFunc = sorting.Reverse(userSortFunc)
		}

		sorting.SortItems(items, sorting.DirFirst, userSortFunc)

		table := table.BuildTable(items, totalSize, os.Stdout)

		table.Render()
	},
}

func init() {
	rootCmd.Flags().StringVarP(&sortBy, "sort", "s", "size", "column output should be sorted by")
	rootCmd.Flags().BoolVarP(&reverse, "reverse", "r", false, "sort in reverse order")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
