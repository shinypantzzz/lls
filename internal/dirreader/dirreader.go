package dirreader

import (
	"io/fs"
	"os"
	"path/filepath"

	"github.com/shinypantzzz/lls/types"
)

func ReadDir(dir string) ([]types.FSItem, int64, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, 0, err
	}

	var data []types.FSItem
	var totalSize int64 = 0
	for _, entry := range entries {
		item := types.FSItem{Name: entry.Name(), IsDir: entry.IsDir(), Size: -1}

		fileInfo, err := entry.Info()
		if err != nil {
			data = append(data, item)
			continue
		}

		if item.IsDir {
			item.Size = dirSize(filepath.Join(dir, entry.Name()))
		} else {
			item.Size = fileInfo.Size()
		}
		totalSize += item.Size
		data = append(data, item)
	}

	return data, totalSize, nil
}

func dirSize(dir string) int64 {
	var length int64 = 0
	filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() {
			fileInfo, err := d.Info()
			if err != nil {
				return err
			}
			length += fileInfo.Size()

		}
		return nil
	})
	return length
}
