package table

import (
	"io"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/shinypantzzz/lls/internal/types"
	"github.com/shinypantzzz/lls/internal/util"
)

var dirStyle = color.New(color.FgBlue)

func BuildTable(items []types.FSItem, totalSize int64, writer io.Writer) *tablewriter.Table {
	table := tablewriter.NewWriter(writer)
	table.SetHeader([]string{"Name", "Size"})

	for _, item := range items {
		name := item.Name
		if item.IsDir {
			name = dirStyle.Sprint(name)
		}
		table.Append([]string{name, util.HumanSize(item.Size)})
	}

	table.SetFooter([]string{"TOTAL:", util.HumanSize(totalSize)})

	return table
}
