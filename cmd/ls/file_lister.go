package ls

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

type FileLister struct {
	ShowHidden bool
}

func (f FileLister) GetEntries() ([]fs.DirEntry, error) {
	entries, err := getAllEntries()
	if err != nil {
		return entries, err
	}

	result := []fs.DirEntry{}
	for _, entry := range entries {
		if f.ShowHidden || shouldShow(entry) {
			result = append(result, entry)
		}
	}

	return result, nil
}

func getAllEntries() ([]fs.DirEntry, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return nil, errors.New("error finding working directory")
	}
	entries, err := os.ReadDir(pwd)
	if err != nil {
		return nil, fmt.Errorf("error listing directory: %s", pwd)
	}
	return entries, nil
}

func shouldShow(entry fs.DirEntry) bool {
	return entry.Name()[0] != '.'
}
