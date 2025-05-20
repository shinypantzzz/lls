package table

import (
	"io"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/shinypantzzz/lls/internal/types"
	"github.com/shinypantzzz/lls/internal/util"
)

type Color int

const (
	WHITE Color = iota
	BLUE
	RED
)

var COLOR_MAP = map[Color]color.Attribute{
	WHITE: color.FgWhite,
	BLUE:  color.FgHiBlue,
	RED:   color.FgRed,
}

var dirStyle = color.New(COLOR_MAP[BLUE])

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

func SetDirStyle(c Color) {
	dirStyle = color.New(COLOR_MAP[c])
}

func SetDirStyleRGB(r int, g int, b int) {
	dirStyle = color.RGB(r, g, b)
}
