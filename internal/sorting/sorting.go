package sorting

import (
	"sort"
	"strings"

	"github.com/shinypantzzz/lls/internal/types"
)

type SortFunc func(types.FSItem, types.FSItem) int

type sorter struct {
	data         []types.FSItem
	sortFuctions []SortFunc
}

func (items *sorter) Len() int {
	return len(items.data)
}

func (items *sorter) Swap(i, j int) {
	items.data[i], items.data[j] = items.data[j], items.data[i]
}

func (items *sorter) Less(i, j int) bool {
	for _, f := range items.sortFuctions {
		if result := f(items.data[i], items.data[j]); result != 0 {
			return result < 0
		}
	}
	return false
}

func SortItems(items []types.FSItem, sortBy ...SortFunc) {
	s := &sorter{
		data:         items,
		sortFuctions: sortBy,
	}
	sort.Sort(s)
}

func DirFirst(first types.FSItem, second types.FSItem) int {
	if first.IsDir == second.IsDir {
		return 0
	} else if first.IsDir {
		return -1
	} else {
		return 1
	}
}

func ByName(first types.FSItem, second types.FSItem) int {
	return strings.Compare(first.Name, second.Name)
}

func BySize(first types.FSItem, second types.FSItem) int {
	return int(first.Size - second.Size)
}

func Reverse(f SortFunc) SortFunc {
	return func(first, second types.FSItem) int {
		return -f(first, second)
	}
}
