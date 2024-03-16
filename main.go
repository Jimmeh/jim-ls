package main

import (
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"os"
)

type myFS struct {
	pwd     func() (string, error)
	listDir func(string) ([]fs.DirEntry, error)
}

var filesystem = myFS{
	pwd:     os.Getwd,
	listDir: os.ReadDir,
}

func NewDirectoryListing() ([]fs.DirEntry, error) {
	dir, err := filesystem.pwd()
	if err != nil {
		return nil, errors.New("unable to get current directory")
	}

	listing, err := listItems(dir)
	if err != nil {
		return nil, errors.New("unable to get contents of current directory")
	}
	return listing, nil
}

func main() {
	all := flag.Bool("a", false, "include hidden files/folders")
	flag.Parse()

	listing, err := NewDirectoryListing()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, entry := range listing {
		entryLine := entry.Name()
		if *all || entryLine[0] != '.' {
			if entry.IsDir() {
				fmt.Printf("\u001b[33m%s\u001b[0m  ", entryLine)
			} else {
				fmt.Printf("%s  ", entryLine)
			}
		}
	}
	fmt.Println()
}

func listItems(dir string) ([]fs.DirEntry, error) {
	listing, err := filesystem.listDir(dir)
	if err != nil {
		return nil, err
	}
	return listing, nil
}
