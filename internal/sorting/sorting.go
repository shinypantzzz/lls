package sorting

import (
	"sort"

	"github.com/shinypantzzz/lls/types"
)

type FSItems []types.FSItem

func (items FSItems) Len() int {
	return len([]types.FSItem(items))
}

func (items FSItems) Less(i, j int) bool {
	return items[i].IsDir && !items[j].IsDir
}

func (items FSItems) Swap(i, j int) {
	items[i], items[j] = items[j], items[i]
}

func SortItems(items []types.FSItem) {
	sort.Sort(FSItems(items))
}
