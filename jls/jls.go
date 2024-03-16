package jls

import (
	"fmt"
	"io/fs"
)

func NewFS(pwd func() (string, error), listDir func(string) ([]fs.DirEntry, error)) filesystem {
	return filesystem{pwd, listDir}
}

type filesystem struct {
	Pwd     func() (string, error)
	ListDir func(string) ([]fs.DirEntry, error)
}

type ListingOptions struct {
	IncludeHidden bool
}

type DirectoryListing struct {
	entries []fs.DirEntry
	opts    ListingOptions
}

func (listing *DirectoryListing) Show(includeHidden bool, disp Display) {
	for _, entry := range listing.entries {
		entryLine := entry.Name()
		if listing.opts.IncludeHidden || entryLine[0] != '.' {
			disp.Show(entry.Name(), entry.IsDir())
		}
	}
}

func NewDirectoryListing(fs filesystem, opts ListingOptions) (DirectoryListing, error) {
	dir, err := fs.Pwd()
	if err != nil {
		return DirectoryListing{}, fmt.Errorf("unable to get current directory: %s", err)
	}

	listing, err := fs.ListDir(dir)
	if err != nil {
		return DirectoryListing{}, fmt.Errorf("unable to get contents of current directory: %s", err)
	}
	return DirectoryListing{listing, opts}, nil
}

type Display interface {
	Show(item string, highlight bool)
}
