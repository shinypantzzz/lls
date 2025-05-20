package cmd

import (
	"fmt"
	"os"

	"github.com/shinypantzzz/lls/internal/dirreader"
	"github.com/shinypantzzz/lls/internal/sorting"
	"github.com/shinypantzzz/lls/internal/table"
	"github.com/shinypantzzz/lls/internal/util"
	"github.com/spf13/cobra"
)

var REVERSE bool
var DIR_COLOR string
var SORT_BY string

var SORT_BY_MAP = map[string]sorting.SortFunc{
	"size": sorting.BySize,
	"name": sorting.ByName,
}

var COLOR_MAP = map[string]table.Color{
	"white": table.WHITE,
	"red":   table.RED,
	"blue":  table.BLUE,
}

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

		userSortFunc := SORT_BY_MAP[SORT_BY]
		if REVERSE {
			userSortFunc = sorting.Reverse(userSortFunc)
		}

		sorting.SortItems(items, sorting.DirFirst, userSortFunc)

		dirColor, ok := COLOR_MAP[DIR_COLOR]
		if !ok {
			r, g, b, err := util.ParseRGB(DIR_COLOR)
			if err == nil {
				table.SetDirStyleRGB(r, g, b)
			}
		} else {
			table.SetDirStyle(dirColor)
		}
		table := table.BuildTable(items, totalSize, os.Stdout)

		table.Render()
	},
}

func init() {
	defaultDirStyle := os.Getenv("LLS_DIR_COLOR")
	rootCmd.Flags().StringVarP(&SORT_BY, "sort", "s", "size", "column output should be sorted by")
	rootCmd.Flags().BoolVarP(&REVERSE, "reverse", "r", false, "sort in reverse order")
	rootCmd.Flags().StringVar(&DIR_COLOR, "dirColor", defaultDirStyle, "color to paint directories, can be either one of preset values or string in 'r,g,b' format")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
