package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	all := flag.Bool("a", false, "include hidden files/folders")
	flag.Parse()

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Unable to get current directory")
		return
	}

	listing, err := os.ReadDir(dir)
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
