package main

import (
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

func main() {
	all := flag.Bool("a", false, "include hidden files/folders")
	flag.Parse()

	dir, err := filesystem.pwd()
	if err != nil {
		fmt.Println("Unable to get current directory")
		return
	}

	listing, err := listItems(dir)
	if err != nil {
		fmt.Println("Unable to get contents of current directory")
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
