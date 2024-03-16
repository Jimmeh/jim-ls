package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Jimmeh/jim-ls/jls"
)

var filesystem = jls.NewFS(os.Getwd, os.ReadDir)

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

	opts := jls.ListingOptions{
		IncludeHidden: *all,
	}
	listing, err := jls.NewDirectoryListing(filesystem, opts)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	display := ConsoleDisplay{}

	listing.Show(*all, display)
	fmt.Println()
}
