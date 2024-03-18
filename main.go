package main

import (
	"flag"

	"github.com/Jimmeh/jim-ls/cmd/ls"
)

func main() {
	showHidden := flag.Bool("a", false, "show files beginning with '.'")
	flag.Parse()

	lister := ls.FileLister{
		ShowHidden: *showHidden,
	}
	console := ls.ConsoleOutput{
		Separator: "  ",
	}

	ls.Ls(lister, console)
}
