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

type DirectoryListing struct {
	entries []fs.DirEntry
}

func (listing *DirectoryListing) Print(includeHidden bool, disp Display) {
	for _, entry := range listing.entries {
		entryLine := entry.Name()
		if includeHidden || entryLine[0] != '.' {
			disp.Show(entry.Name(), entry.IsDir())
		}
	}
}

func NewDirectoryListing() (DirectoryListing, error) {
	dir, err := filesystem.pwd()
	if err != nil {
		return DirectoryListing{}, fmt.Errorf("unable to get current directory: %s", err)
	}

	listing, err := filesystem.listDir(dir)
	if err != nil {
		return DirectoryListing{}, fmt.Errorf("unable to get contents of current directory: %s", err)
	}
	return DirectoryListing{listing}, nil
}

type Display interface {
	Show(item string, highlight bool)
}

type ConsoleDisplay struct {
}

func (c ConsoleDisplay) Show(item string, highlight bool) {
	if highlight {
		fmt.Printf("\u001b[33m%s\u001b[0m  ", item)
	} else {
		fmt.Printf("%s  ", item)
	}
}

func main() {
	all := flag.Bool("a", false, "include hidden files/folders")
	flag.Parse()

	listing, err := NewDirectoryListing()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	display := ConsoleDisplay{}

	listing.Print(*all, display)
	fmt.Println()
}
